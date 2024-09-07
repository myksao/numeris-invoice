package note

import (
	"context"
	"invoice/internal/note/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, req *domain.NoteReq) (id utils.DefaultCreateRes, err error)
	RetrieveByEntityID(ctx context.Context, ent *domain.NoteEntityIDReq) ([]*domain.Note, error)
	RetrieveByEntity(ctx context.Context, ent *domain.NoteEntityReq) ([]*domain.Note, error)
	RetrieveByID(ctx context.Context, id string) (*domain.Note, error)
}
