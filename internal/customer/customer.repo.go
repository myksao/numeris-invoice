package customer

import (
	"context"
	"invoice/internal/customer/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, customer *domain.CustomerReq) (id utils.DefaultCreateRes, err error)
	RetrieveByOutletID(ctx context.Context, req *domain.CustomerOutletReq, page *utils.PaginationReq) ([]*domain.Customer, error)
	RetrieveByID(ctx context.Context, id string) (*domain.Customer, error)
}
