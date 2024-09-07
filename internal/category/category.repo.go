package category

import (
	"context"
	"invoice/internal/category/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, req *domain.CategoryReq) (id utils.DefaultCreateRes, err error)
	RetrieveByOutletID(ctx context.Context, req *domain.CategoryOutletReq, page *utils.PaginationReq) ([]*domain.Category, error)
	RetrieveByID(ctx context.Context, id string) (*domain.Category, error)
}
