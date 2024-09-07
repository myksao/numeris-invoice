package repo

import (
	"context"
	"database/sql"
	"invoice/internal/organisation/domain"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

func (repo *Repo) Create(ctx context.Context, org *domain.OrganisationReq) (id domain.CreateOrgRes, err error) {
	if err := repo.validator.Struct(org); err != nil {
		repo.logger.Error("Error validating organisation", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createOrganisation,
		ksuid.New().String(),
		org.Name,
		org.Reference,
		org.Address,
	).StructScan(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return id, errors.New("duplicate key value violates unique constraint")
		}
		return id, errors.Wrap(err, "Error creating organisation")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.Organisation, error) {
	org := &domain.Organisation{}
	if err := repo.db.GetContext(ctx, org, retrieveOrganisationByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving organisation")
	}
	return org, nil
}
