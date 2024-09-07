package repo

import (
	"context"
	"invoice/internal/inventory/domain"
	variant_repo "invoice/internal/variant"

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
	variant   variant_repo.Repo
}

func NewRepo(logger *zap.Logger, db *sqlx.DB, variant variant_repo.Repo) *Repo {
	return &Repo{
		db:        db,
		logger:    logger,
		validator: validator.New(),
		variant:   variant,
	}
}

func (repo *Repo) Process(ctx context.Context, req *domain.InventoryProcess) (id *string, err error) {
	if err := repo.validator.Struct(req); err != nil {
		repo.logger.Error("Error validating customer", zap.Error(err))
		return id, err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return id, errors.Wrap(err, "Error starting transaction")
	}

	var inventory_id string
	if err := tx.QueryRowContext(
		ctx,
		fetchInventory,
		req.VariantID,
		req.MeasureID,
	).Scan(&inventory_id); err != nil {
		return id, errors.Wrap(err, "Error creating inventory")
	}

	repo.logger.Info("Result", zap.Any("result", inventory_id))

	measure, err := repo.variant.RetrieveMeasureByID(ctx, req.MeasureID)
	if err != nil {
		return id, errors.Wrap(err, "Error retrieving measure")
	}
	measure.ID = req.MeasureID

	repo.logger.Info("Measure", zap.Any("measure", measure))

	switch req.State {
	case "open":
		{
			req.EntityID = inventory_id
			req.Entity = "inventory"
			if _, err := tx.ExecContext(
				ctx,
				createInventory,
				req.VariantID,
				req.MeasureID,
			); err != nil {
				tx.Rollback()
				return id, errors.Wrap(err, "Error creating inventory")
			}
		}
	case "none":
		{
			if _, err := tx.ExecContext(
				ctx,
				createInventory,
				req.VariantID,
				req.MeasureID,
			); err != nil {
				tx.Rollback()
				return id, errors.Wrap(err, "Error creating inventory")
			}
		}
	case "added":
		fallthrough
	case "returned":
		{
			req.EntityID = inventory_id
			req.Entity = "inventory"
			if _, err := tx.ExecContext(
				ctx,
				addInventory,
				req.Quantity,
				req.VariantID,
				req.MeasureID,
			); err != nil {
				tx.Rollback()
				return id, errors.Wrap(err, "Error processing inventory")
			}
		}
	case "issued":
		{
			if _, err := tx.ExecContext(
				ctx,
				issueInventory,
				req.Quantity,
				req.VariantID,
				req.MeasureID,
			); err != nil {
				tx.Rollback()
				return id, errors.Wrap(err, "Error processing inventory")
			}
		}
	case "hold":
		{
			if _, err := tx.ExecContext(
				ctx,
				issueInventory,
				req.Quantity,
				req.VariantID,
				req.MeasureID,
			); err != nil {
				tx.Rollback()
				return id, errors.Wrap(err, "Error processing inventory")
			}
		}
	}

	if req.State != "none" {
		if _, err := tx.ExecContext(
			ctx,
			batchEntry,
			ksuid.New().String(),
			req.VariantID,
			req.MeasureID,
			req.EntityID,
			req.Entity,
			req.State,
			req.Quantity,
		); err != nil {
			tx.Rollback()
			return id, errors.Wrap(err, "Error processing inventory")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "Error committing transaction")
	}

	id = &inventory_id

	return id, nil
}
