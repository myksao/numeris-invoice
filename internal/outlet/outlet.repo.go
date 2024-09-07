package outlet

import (
	"context"
	"invoice/internal/outlet/domain"
)

type Repo interface {
	Create(ctx context.Context, org *domain.OutletReq) (id domain.CreateOutletRes, err error)
	RetrieveByID(ctx context.Context, id string) (*domain.Outlet, error)
	RetrieveByOrgID(ctx context.Context, orgID string, limit, offset string) ([]*domain.Outlet, error)
}
