package repo

import (
	"context"
	"encoding/json"
	"invoice/internal/invoice/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestInvoice(t *testing.T) {
	t.Parallel()

	logger := zap.NewNop()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repo := NewRepo(logger, sqlxDB, nil, nil)

	t.Run("Create", func(t *testing.T) {

		name := "name"
		ref := "title"
		currency_id := "content"
		customer_id := "customer_id"
		outlet_id := "outlet_id"
		due_date := "due_date"
		total := 100.00
		subtotal := 100.00
		discount := 100.00
		reminder, _ := json.Marshal([]string{"* * * * *"})
		bank_account_id := "bank_account_id"
		status := "status"
		created_by := "created_by"
		// Mock transaction begin
		mock.ExpectBegin()

		invoice_id := ksuid.New().String()

		rows := sqlmock.NewRows([]string{"id", "name", "ref", "currency_id",
			"customer_id", "outlet_id", "due_date", "total", "subtotal",
			"discount", "reminder", "bank_account_id", "status", "created_by",
		}).AddRow(
			invoice_id,
			name, ref, currency_id,
			customer_id, outlet_id, due_date, total, subtotal,
			discount, reminder, bank_account_id, status, created_by,
		)

		req := &domain.InvoiceReq{
			Name:          name,
			Ref:           ref,
			CurrencyID:    currency_id,
			CustomerID:    customer_id,
			OutletID:      outlet_id,
			DueDate:       due_date,
			Total:         total,
			SubTotal:      subtotal,
			Discount:      &discount,
			Reminder:      reminder,
			BankAccountID: bank_account_id,
			Status:        status,
			CreatedBy:     created_by,
		}

		mock.ExpectQuery(createInvoice).WithArgs(
			invoice_id,
			req.Name,
			req.Ref,
			req.DueDate,
			req.Total,
			req.Status,
			req.Reminder,
			req.CreatedBy,
			req.BankAccountID,
			req.CurrencyID,
			req.CustomerID,
			req.SubTotal,
			req.Discount,
			req.OutletID,
		).WillReturnRows(rows)

		// Mock transaction commit
		mock.ExpectCommit()

		create, err := repo.Create(context.Background(), req)

		require.NoError(t, err)
		require.NotNil(t, create)
		require.Equal(t, req.Name, req.Ref)
	})
}
