package repo

import (
	"context"
	"database/sql"
	"invoice/internal/user/domain"
	"invoice/pkg/utils"
	"log"

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

func (repo *Repo) Create(ctx context.Context, user *domain.UserReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(user); err != nil {
		repo.logger.Error("Error validating user", zap.Error(err))
		return id, err
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return id, errors.New("Error hashing password")
	}

	//  err = auth.CheckPassword(hashedPassword, password)
	// if err != nil {
	//     log.Fatalf("Invalid password: %v", err)
	// }

	if err := repo.db.QueryRowxContext(
		ctx,
		createUser,
		ksuid.New().String(),
		user.Name,
		user.Username,
		hash,
		user.Ref,
		user.OutletID,
	).StructScan(&id); err != nil {
		return id, errors.Wrap(err, "Error creating user")
	}

	return id, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	if err := repo.db.GetContext(ctx, user, retrieveUserByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving user")
	}
	return user, nil
}

func (repo *Repo) RetrieveByOutletID(ctx context.Context, id string, req *utils.PaginationReq) ([]*domain.User, error) {
	users := []*domain.User{}
	if err := repo.db.SelectContext(
		ctx,
		&users,
		retrieveUserByOutletID,
		id,
		req.Limit,
		req.Offset,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving user")
	}
	return users, nil
}
