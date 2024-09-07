package repo

import (
	"context"
	"encoding/json"
	"invoice/internal/inventory"
	"invoice/internal/invoice/domain"
	"invoice/internal/note"
	note_domain "invoice/internal/note/domain"
	"invoice/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

type BoqRepo struct {
	db        *sqlx.DB
	logger    *zap.Logger
	validator *validator.Validate
	inventory inventory.Repo
	note      note.Repo
}

func BoqNewRepo(logger *zap.Logger, db *sqlx.DB, inventory inventory.Repo, note note.Repo) *BoqRepo {
	return &BoqRepo{
		db:        db,
		logger:    logger,
		validator: validator.New(),
		inventory: inventory,
		note:      note,
	}
}

func (repo *BoqRepo) Create(ctx context.Context, invoice_id string, invoiceboq []*domain.UpdateInvoiceBoqReq) (id utils.DefaultCreateRes, err error) {

	var boqm struct {
		Boq []*domain.UpdateInvoiceBoqReq `json:"boq" validate:"required"`
	}

	boqm.Boq = invoiceboq

	if err := repo.validator.Struct(boqm); err != nil {
		repo.logger.Error("Error validating invoice", zap.Error(err))
		return id, err
	}

	tx, err := repo.db.BeginTxx(ctx, nil)

	if err != nil {
		return id, err
	}

	invoice := &domain.Invoice{}
	if err := tx.GetContext(ctx, invoice, retrieveInvoiceByID, invoice_id); err != nil {
		tx.Rollback()
		return id, errors.Wrap(err, "Error fetching invoice")
	}

	for _, boq := range invoiceboq {
		if boq.ID != "" {
			if _, err := tx.ExecContext(
				ctx,
				updateInvoiceBoq,
				boq.Quantity,
				boq.UnitPrice,
				boq.Total,
				boq.ID,
			); err != nil {
				tx.Rollback()
				return id, errors.Wrap(err, "Error updating invoice boq")
			}

			marshal, _ := json.Marshal(boq)
			repo.note.Create(ctx, &note_domain.NoteReq{
				Entity:    "invoice",
				EntityID:  invoice_id,
				Note:      "Invoice Boq updated",
				Command:   marshal,
				CreatedBy: invoice.CreatedBy, // Wrong - use logged in user
			})
			continue
		}
		if _, err := tx.ExecContext(
			ctx,
			createInvoiceBoq,
			ksuid.New().String(),
			boq.VariantID,
			invoice.ID,
			boq.MeasureID,
			boq.Quantity,
			boq.UnitPrice,
			boq.Total,
		); err != nil {
			tx.Rollback()
			return id, errors.Wrap(err, "Error creating currency variant measure")
		}
	}

	if err := tx.Commit(); err != nil {
		return id, errors.Wrap(err, "Error committing transaction")
	}

	return id, nil
}
