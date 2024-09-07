package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"invoice/internal/inventory"
	"invoice/internal/invoice/domain"
	"invoice/internal/note"
	"invoice/pkg/utils"
	"time"

	inventory_domain "invoice/internal/inventory/domain"
	note_domain "invoice/internal/note/domain"

	"github.com/cockroachdb/apd/v3"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

type Repo struct {
	db        *sqlx.DB
	logger    *zap.Logger
	validator *validator.Validate
	inventory inventory.Repo
	note      note.Repo
}

func NewRepo(logger *zap.Logger, db *sqlx.DB, inventory inventory.Repo, note note.Repo) *Repo {
	return &Repo{
		db:        db,
		logger:    logger,
		validator: validator.New(),
		inventory: inventory,
		note:      note,
	}
}

func (repo *Repo) Create(ctx context.Context, invoice *domain.InvoiceReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(invoice); err != nil {
		repo.logger.Error("Error validating invoice", zap.Error(err))
		return id, err
	}

	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return id, errors.Wrap(err, "Error beginning transaction")
	}

	ctxapd := apd.BaseContext
	total := new(apd.Decimal)

	for _, boq := range invoice.Boq {

		boq_quantity, _, verr := apd.NewFromString(fmt.Sprintf("%f", boq.Quantity))
		if verr != nil {
			repo.logger.Error("Error creating decimal from string:", zap.Error(verr))
			return id, errors.Wrap(verr, "Error creating decimal from string")
		}

		boq_unit_price, _, uperr := apd.NewFromString(fmt.Sprintf("%f", boq.UnitPrice))
		if uperr != nil {
			repo.logger.Error("Error creating decimal from string:", zap.Error(uperr))
			return id, errors.Wrap(uperr, "Error creating decimal from string")
		}

		boq_total := new(apd.Decimal)
		_, terr := ctxapd.Mul(boq_total, boq_quantity, boq_unit_price)
		if terr != nil {
			repo.logger.Error("Error multiplying numbers:", zap.Error(terr))
			return id, errors.Wrap(terr, "Error multiplying numbers")
		}

		_, rerr := ctxapd.Add(total, total, boq_total)
		if rerr != nil {
			repo.logger.Error("Error adding numbers:", zap.Error(rerr))
			return id, errors.Wrap(rerr, "Error adding numbers")
		}
	}

	fmt.Printf("Total: %s\n", total.String())

	invoice_id := ksuid.New().String()
	if err := tx.QueryRowxContext(
		ctx,
		createInvoice,
		invoice_id,
		invoice.Name,
		invoice.Ref,
		invoice.DueDate,
		total.String(),
		invoice.Status,
		invoice.Reminder,
		invoice.CreatedBy,
		invoice.BankAccountID,
		invoice.CurrencyID,
		invoice.CustomerID,
		invoice.SubTotal,
		invoice.Discount,
		invoice.OutletID,
	).StructScan(&id); err != nil {
		tx.Rollback()
		return id, errors.Wrap(err, "Error creating invoice")
	}

	for _, boq := range invoice.Boq {

		if _, err := tx.ExecContext(
			ctx,
			createInvoiceBoq,
			ksuid.New().String(),
			boq.VariantID,
			invoice_id,
			boq.MeasureID,
			boq.Quantity,
			boq.UnitPrice,
			boq.Total,
		); err != nil {
			tx.Rollback()
			return id, errors.Wrap(err, "Error creating currency variant measure")
		}
	}

	due, _ := time.Parse("2006-01-02 15:04:05", invoice.DueDate)
	due_date := utils.ConvertToCron(due)
	fmt.Print(due_date)
	if _, err := repo.db.ExecContext(ctx, triggerOverDueInvoice, due_date); err != nil {
		repo.logger.Error("Error creating cron job", zap.Error(err))
	}

	marshal, _ := json.Marshal(invoice)
	repo.note.Create(ctx, &note_domain.NoteReq{
		Entity:    "invoice",
		EntityID:  invoice_id,
		Note:      "Invoice created",
		Command:   marshal,
		CreatedBy: invoice.CreatedBy,
	})

	if err := tx.Commit(); err != nil {
		return id, errors.Wrap(err, "Error committing transaction")
	}

	return id, nil
}

func (repo *Repo) UpdateStatus(ctx context.Context, status *domain.InvoiceStatusReq, uid *string) (id *string, err error) {
	if err := repo.validator.Var(status, "required"); err != nil {
		repo.logger.Error("Error validating status", zap.Error(err))
		return id, err
	}
	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error beginning transaction")
	}

	if _, err := tx.ExecContext(ctx, updateInvoiceStatus, status.Status, uid); err != nil {
		return nil, errors.Wrap(err, "Error updating invoice status")
	}

	if status.Status == "pending" {
		boqs := []*domain.InvoiceBoq{}
		if err := tx.SelectContext(ctx, &boqs, retrieveInvoiceBoqByInvoiceID, uid); err != nil {
			return nil, errors.Wrap(err, "Error retrieving invoice Boqs")
		}
		for _, boq := range boqs {
			repo.inventory.Process(ctx, &inventory_domain.InventoryProcess{
				VariantID: boq.VariantID,
				MeasureID: boq.MeasureID,
				Quantity:  boq.Quantity,
				State:     "hold",
			})
		}
	}

	marshal_note, _ := json.Marshal(map[string]any{
		"payload": status,
		"id":      *uid,
	})
	repo.note.Create(ctx, &note_domain.NoteReq{
		Entity:    "invoice",
		EntityID:  *uid,
		Note:      fmt.Sprintf("Invoice status updated to %s", status.Status),
		CreatedBy: status.UpdatedBy,
		Command:   marshal_note,
	})

	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "Error committing transaction")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.InvoiceRes, error) {
	invoice := &domain.InvoiceRes{}
	if err := repo.db.GetContext(ctx, invoice, retrieveInvoiceByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving invoice")
	}
	return invoice, nil
}

func (repo *Repo) FetchBoqs(ctx context.Context, id string) ([]*domain.InvoiceBoq, error) {
	boqs := []*domain.InvoiceBoq{}
	if err := repo.db.SelectContext(ctx, &boqs, retrieveInvoiceBoqByInvoiceID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving invoice Boqs")
	}
	return boqs, nil
}

func (repo *Repo) RetrieveSummary(ctx context.Context, id string) (*domain.InvoiceSummary, error) {
	summary := &domain.InvoiceSummary{}
	if err := repo.db.GetContext(ctx, summary, retrieveInvoiceSummary, id); err != nil {
		return nil, errors.Wrap(err, "Error retrieving invoice summary")
	}
	return summary, nil
}
