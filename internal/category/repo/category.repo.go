package repo

import (
	"context"
	"database/sql"
	"invoice/internal/category/domain"
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
}

func NewRepo(logger *zap.Logger, db *sqlx.DB) *Repo {
	return &Repo{
		db:        db,
		logger:    logger,
		validator: validator.New(),
	}
}

func (repo *Repo) Create(ctx context.Context, req *domain.CategoryReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(req); err != nil {
		repo.logger.Error("Error validating category", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createCategory,
		ksuid.New().String(),
		req.Name,
		req.OutletID,
	).StructScan(&id); err != nil {
		return id, errors.Wrap(err, "Error creating category")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.Category, error) {
	category := &domain.Category{}
	if err := repo.db.GetContext(ctx, category, retrieveCategoryByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving category")
	}
	return category, nil
}

func (repo *Repo) RetrieveByOutletID(ctx context.Context, req *domain.CategoryOutletReq, page *utils.PaginationReq) ([]*domain.Category, error) {
	categories := []*domain.Category{}
	if err := repo.db.SelectContext(
		ctx,
		&categories,
		retrieveCategoriesByOutletID,
		req.ID,
		page.Limit,
		page.Offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving categories")
	}
	return categories, nil
}
