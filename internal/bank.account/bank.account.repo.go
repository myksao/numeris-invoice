package bankaccount

import (
	"context"
	"invoice/internal/bank.account/domain"
	"invoice/pkg/utils"
)

type Repo interface {
	Create(ctx context.Context, req *domain.BankAccountReq) (id utils.DefaultCreateRes, err error)
	RetrieveByOutletID(ctx context.Context, req *domain.BankAccountOutletReq, page *utils.PaginationReq) ([]*domain.BankAccount, error)
	RetrieveByID(ctx context.Context, id string) (*domain.BankAccount, error)
}
