package repo

import (
	"context"
	"database/sql"
	"fmt"
	"invoice/internal/currency/domain"
	"invoice/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

type Repo struct {
	db        *sqlx.DB
	logger    *zap.Logger
	validator *validator.Validate
}

func NewRepo(logger *zap.Logger, db *sqlx.DB) *Repo {
	return &Repo{
		db:        db,
		logger:    logger,
		validator: validator.New(),
	}
}

func (repo *Repo) Create(ctx context.Context, curr *domain.CurrencyReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(curr); err != nil {
		repo.logger.Error("Error validating currency", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createCurrency,
		ksuid.New().String(),
		curr.Name,
		curr.Code,
		curr.Symbol,
	).StructScan(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return id, errors.New("duplicate key value violates unique constraint")
		}
		return id, errors.Wrap(err, "Error creating currency")
	}

	return id, nil
}

func (repo *Repo) Retrieve(ctx context.Context, page *utils.PaginationReq) ([]*domain.Currency, error) {
	var currencies []*domain.Currency
	if err := repo.db.SelectContext(ctx, &currencies, retrieveCurrencies, page.Limit, page.Offset); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving currencies")
	}
	return currencies, nil
}

func (repo *Repo) CreateCurrencyVariantMeasure(ctx context.Context, measure []*domain.CurrencyMeasureReq) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "Error starting transaction")
	}

	for _, m := range measure {
		fmt.Printf("CurrencyID: %s, MeasureID: %s, Price: %s\n", m.CurrencyID, m.MeasureID, m.Price)
		if _, err := tx.ExecContext(
			ctx,
			createCurrencyMeasure,
			ksuid.New().String(),
			m.CurrencyID,
			m.MeasureID,
			m.Price,
		); err != nil {
			tx.Rollback()
			return errors.Wrap(err, "Error creating currency variant measure")
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "Error committing transaction")
	}

	return nil
}
