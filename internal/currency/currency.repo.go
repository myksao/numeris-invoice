package currency

import (
	"context"
	"invoice/internal/currency/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, curr *domain.CurrencyReq) (id utils.DefaultCreateRes, err error)
	Retrieve(ctx context.Context, page *utils.PaginationReq) ([]*domain.Currency, error)
	CreateCurrencyVariantMeasure(ctx context.Context, measure []*domain.CurrencyMeasureReq) error
}
