package repo

import (
	"context"
	"database/sql"
	"invoice/internal/customer/domain"
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

func (repo *Repo) Create(ctx context.Context, customer *domain.CustomerReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(customer); err != nil {
		repo.logger.Error("Error validating customer", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createCustomer,
		ksuid.New().String(),
		customer.Name,
		customer.Email,
		customer.MobileNo,
		customer.Address,
		customer.OutletID,
	).StructScan(&id); err != nil {
		return id, errors.Wrap(err, "Error creating customer")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.Customer, error) {
	customer := &domain.Customer{}
	if err := repo.db.GetContext(ctx, customer, retrieveCustomerByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving customer")
	}
	return customer, nil
}

func (repo *Repo) RetrieveByOutletID(ctx context.Context, req *domain.CustomerOutletReq, page *utils.PaginationReq) ([]*domain.Customer, error) {
	customers := []*domain.Customer{}
	if err := repo.db.SelectContext(
		ctx,
		&customers,
		retrieveCustomerByOutletID,
		req.ID,
		page.Limit,
		page.Offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving customers")
	}
	return customers, nil
}
