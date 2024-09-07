package user

import (
	"context"
	"invoice/internal/user/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, user *domain.UserReq) (id utils.DefaultCreateRes, err error)
	RetrieveByID(ctx context.Context, id string) (*domain.User, error)
	RetrieveByOutletID(ctx context.Context, id string, req *utils.PaginationReq) ([]*domain.User, error)
}
