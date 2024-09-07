package repo

import (
	"context"
	"database/sql"
	"invoice/internal/bank.account/domain"
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

func (repo *Repo) Create(ctx context.Context, req *domain.BankAccountReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(req); err != nil {
		repo.logger.Error("Error validating Bank account", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createBankAccount,
		ksuid.New().String(),
		req.Name,
		req.OutletID,
		req.AccountNo,
		req.RoutingNo,
		req.AccountType,
		req.BankName,
		req.CurrencyID,
	).StructScan(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return id, errors.New("duplicate key value violates unique constraint")
		}
		return id, errors.Wrap(err, "Error creating bank account")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.BankAccount, error) {
	bankAccount := &domain.BankAccount{}
	if err := repo.db.GetContext(ctx, bankAccount, retrieveBankAccountByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving Bank account")
	}
	return bankAccount, nil
}

func (repo *Repo) RetrieveByOutletID(ctx context.Context, req *domain.BankAccountOutletReq, page *utils.PaginationReq) ([]*domain.BankAccount, error) {
	bankAccounts := []*domain.BankAccount{}
	if err := repo.db.SelectContext(
		ctx,
		&bankAccounts,
		retrieveBankAccountByOutletID,
		req.ID,
		page.Limit,
		page.Offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving Bank account")
	}
	return bankAccounts, nil
}
