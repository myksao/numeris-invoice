package organisation

import (
	"context"
	"invoice/internal/organisation/domain"
)

type Repo interface {
	Create(ctx context.Context, org *domain.OrganisationReq) (id domain.CreateOrgRes, err error)
	RetrieveByID(ctx context.Context, id string) (*domain.Organisation, error)
}
