package repo

import (
	"context"
	"database/sql"
	"invoice/internal/outlet/domain"

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

func (repo *Repo) Create(ctx context.Context, outlet *domain.OutletReq) (id domain.CreateOutletRes, err error) {
	if err := repo.validator.Struct(outlet); err != nil {
		repo.logger.Error("Error validating outlet", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createOutlet,
		ksuid.New().String(),
		outlet.Name,
		outlet.IsDefault,
		outlet.OrgID,
		outlet.Address,
	).StructScan(&id); err != nil {
		return id, errors.Wrap(err, "Error creating organisation")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.Outlet, error) {
	outlet := &domain.Outlet{}
	if err := repo.db.GetContext(ctx, outlet, retrieveOutletByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return outlet, nil
}

func (repo *Repo) RetrieveByOrgID(ctx context.Context, orgID string, limit, offset string) ([]*domain.Outlet, error) {
	outlets := []*domain.Outlet{}
	if err := repo.db.SelectContext(
		ctx,
		&outlets,
		retrieveOutletByOrgID,
		orgID,
		limit,
		offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return outlets, nil
}
