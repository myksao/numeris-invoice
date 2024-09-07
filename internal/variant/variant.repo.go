package variant

import (
	"context"
	"invoice/internal/variant/domain"
	"invoice/pkg/utils"

	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(ctx context.Context, customer *domain.VariantReq) (id utils.DefaultCreateRes, err error)
	RetrieveByID(ctx context.Context, id string) (*domain.VariantRes, error)
	RetrieveByItemID(ctx context.Context, req *domain.VariantItemReq, page *utils.PaginationReq) ([]*domain.VariantItem, error)
	CreateMeasure(ctx context.Context, tx *sqlx.Tx, req *domain.MeasureReq) error
	RetrieveMeasureByID(ctx context.Context, id string) (measure domain.Measure, err error)
	RetrieveMeasureByVariantID(ctx context.Context, id string) (measure []*domain.MeasureRes, err error)
}
