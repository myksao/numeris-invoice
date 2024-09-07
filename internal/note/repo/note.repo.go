package repo

import (
	"context"
	"database/sql"
	"invoice/internal/note/domain"
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

func (repo *Repo) Create(ctx context.Context, req *domain.NoteReq) (id utils.DefaultCreateRes, err error) {
	if err := repo.validator.Struct(req); err != nil {
		repo.logger.Error("Error validating Note", zap.Error(err))
		return id, err
	}

	if err := repo.db.QueryRowxContext(
		ctx,
		createNote,
		ksuid.New().String(),
		req.EntityID,
		req.Entity,
		req.Note,
		req.CreatedBy,
		req.Command,
	).StructScan(&id); err != nil {
		repo.logger.Error("Error creating Note", zap.Error(err))
		return id, errors.Wrap(err, "Error creating Note")
	}

	return id, nil
}

func (repo *Repo) RetrieveByEntityID(ctx context.Context, ent *domain.NoteEntityIDReq) ([]*domain.Note, error) {
	notes := []*domain.Note{}
	if err := repo.db.SelectContext(ctx, &notes, retrieveNoteByEntityID, ent.ID, ent.Entity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving Notes")
	}

	return notes, nil
}

func (repo *Repo) RetrieveByEntity(ctx context.Context, ent *domain.NoteEntityReq) ([]*domain.Note, error) {
	notes := []*domain.Note{}
	if err := repo.db.SelectContext(ctx, &notes, retrieveNoteByEntity, ent.Entity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving Notes")
	}

	return notes, nil
}

func (repo *Repo) RetrieveByID(ctx context.Context, id string) (*domain.Note, error) {
	note := &domain.Note{}
	if err := repo.db.GetContext(ctx, note, retrieveNoteByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error retrieving Note")
	}
	return note, nil
}
