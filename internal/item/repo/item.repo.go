package repo

import (
	"context"
	"database/sql"
	"invoice/internal/item/domain"
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

func (repo *Repo) Create(ctx context.Context, req *domain.ItemReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(req); err != nil {
		repo.logger.Error("Error validating item", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createItem,
		ksuid.New().String(),
		req.Name,
		req.Description,
		req.CategoryID,
		req.SKU,
		req.OutletID,
		req.CreatedBy,
	).StructScan(&id); err != nil {
		return id, errors.Wrap(err, "Error creating item")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.Item, error) {
	item := &domain.Item{}
	if err := repo.db.GetContext(ctx, item, retrieveItemByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving item")
	}
	return item, nil
}

func (repo *Repo) RetrieveByOutletID(ctx context.Context, req *domain.ItemOutletReq, page *utils.PaginationReq) ([]*domain.Item, error) {
	items := []*domain.Item{}
	if err := repo.db.SelectContext(
		ctx,
		&items,
		retrieveItemByOutletID,
		req.ID,
		page.Limit,
		page.Offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving items")
	}
	return items, nil
}
