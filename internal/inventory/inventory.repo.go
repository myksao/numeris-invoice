package inventory

import (
	"context"
	"invoice/internal/inventory/domain"
)

type Repo interface {
	Process(ctx context.Context, req *domain.InventoryProcess) (id *string, err error)
}
