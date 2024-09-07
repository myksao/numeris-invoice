package item

import (
	"context"
	"invoice/internal/item/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, customer *domain.ItemReq) (id utils.DefaultCreateRes, err error)
	RetrieveByOutletID(ctx context.Context, req *domain.ItemOutletReq, page *utils.PaginationReq) ([]*domain.Item, error)
	RetrieveByID(ctx context.Context, id string) (*domain.Item, error)
}
