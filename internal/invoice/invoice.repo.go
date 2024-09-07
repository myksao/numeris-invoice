package invoice

import (
	"context"
	"invoice/internal/invoice/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, invoice *domain.InvoiceReq) (id utils.DefaultCreateRes, err error)
	UpdateStatus(ctx context.Context, status *domain.InvoiceStatusReq, uid *string) (id *string, err error)
	RetrieveByID(ctx context.Context, id string) (*domain.InvoiceRes, error)
	FetchBoqs(ctx context.Context, id string) ([]*domain.InvoiceBoq, error)
	RetrieveSummary(ctx context.Context, id string) (*domain.InvoiceSummary, error)
}
