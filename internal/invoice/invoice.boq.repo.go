package invoice

import (
	"context"
	"invoice/internal/invoice/domain"
	"invoice/pkg/utils"
)

type BoqRepo interface {
	Create(ctx context.Context, invoice_id string, invoiceboq []*domain.UpdateInvoiceBoqReq) (id utils.DefaultCreateRes, err error)
}
