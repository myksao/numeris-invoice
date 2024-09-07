package repo

import (
	"context"
	"database/sql"
	currency_domin "invoice/internal/currency/domain"
	currency_repo "invoice/internal/currency/repo"
	"invoice/internal/variant/domain"
	"invoice/pkg/utils"

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
	currency  *currency_repo.Repo
}

func NewRepo(logger *zap.Logger, db *sqlx.DB, currency *currency_repo.Repo) *Repo {
	return &Repo{
		db:        db,
		logger:    logger,
		validator: validator.New(),
		currency:  currency,
	}
}

func (repo *Repo) Create(ctx context.Context, req *domain.VariantReq) (id utils.DefaultCreateRes, err error) {
	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return utils.DefaultCreateRes{}, errors.Wrap(err, "Error starting transaction")
	}
	if err := repo.validator.Struct(req); err != nil {
		repo.logger.Error("Error validating item", zap.Error(err))
		return id, err
	}

	if err := tx.QueryRowxContext(
		ctx,
		createVariant,
		ksuid.New().String(),
		req.Name,
		req.Description,
		req.ItemID,
		req.OutletID,
	).StructScan(&id); err != nil {
		tx.Rollback()
		return id, errors.Wrap(err, "Error creating item")
	}

	for _, measure := range req.Measure {
		err := repo.CreateMeasure(ctx, tx, &domain.MeasureReq{
			Entity:   "variant",
			EntityID: id.ID,
			Unit:     measure.Unit,
			Quantity: measure.Quantity,
			Currency: measure.Currency,
		})

		if err != nil {
			tx.Rollback()
			return utils.DefaultCreateRes{}, errors.Wrap(err, "Error creating measure")
		}
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.VariantRes, error) {
	variant := &domain.VariantRes{}
	if err := repo.db.GetContext(ctx, variant, retrieveVariantByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving item")
	}
	return variant, nil
}

func (repo *Repo) RetrieveByItemID(ctx context.Context, req *domain.VariantItemReq, page *utils.PaginationReq) ([]*domain.VariantItem, error) {
	variants := []*domain.VariantItem{}
	if err := repo.db.SelectContext(
		ctx,
		&variants,
		retrieveVariantByItemID,
		req.ID,
		page.Limit,
		page.Offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving item")
	}
	return variants, nil
}

func (repo *Repo) CreateMeasure(ctx context.Context, tx *sqlx.Tx, req *domain.MeasureReq) error {
	var id utils.DefaultCreateRes
	if err := tx.QueryRowxContext(
		ctx,
		createMeasure,
		ksuid.New().String(),
		req.Entity,
		req.EntityID,
		req.Unit,
		req.Quantity,
	).StructScan(&id); err != nil {
		return errors.Wrap(err, "Error creating measure")
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "Error committing transaction")
	}

	if _, err := repo.db.ExecContext(
		ctx,
		createInventory,
		ksuid.New().String(),
		req.EntityID,
		id.ID,
	); err != nil {
		return errors.Wrap(err, "Error creating inventory")
	}

	for _, currency := range req.Currency {
		currency_measure := make([]*currency_domin.CurrencyMeasureReq, 0)
		currency_measure = append(currency_measure, &currency_domin.CurrencyMeasureReq{
			CurrencyID: currency.CurrencyID,
			MeasureID:  id.ID,
			Price:      currency.Price,
		})

		err := repo.currency.CreateCurrencyVariantMeasure(ctx, currency_measure)

		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "Error creating currency variant measure")
		}
	}

	return nil
}

/**
 * RetrieveMeasureByID - (Used by inventory repo)
 * @param ctx context.Context
 * @param id string
 * @return *domain.Measure, error
 */
func (repo *Repo) RetrieveMeasureByID(ctx context.Context, id string) (measure domain.Measure, err error) {
	if err := repo.db.GetContext(
		ctx,
		&measure,
		retrieveMeasureByID,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return domain.Measure{}, nil
		}
		return domain.Measure{}, errors.Wrap(err, "Error retrieving measures")
	}

	return measure, nil
}

func (repo *Repo) RetrieveMeasureByVariantID(ctx context.Context, id string) (measure []*domain.MeasureRes, err error) {
	if err := repo.db.SelectContext(
		ctx,
		&measure,
		retrieveMeasureByVariantID,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving measures")
	}

	return measure, nil
}
