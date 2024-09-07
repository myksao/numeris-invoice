package domain

import "encoding/json"

type Invoice struct {
	ID            string          `json:"id" db:"id" validate:"required"`
	Name          string          `json:"name" db:"name" validate:"required"`
	Ref           string          `json:"ref" db:"ref" validate:"required"`
	CurrencyID    string          `json:"currency_id" db:"currency_id" validate:"required"`
	OutletID      *string         `json:"outlet_id" db:"outlet_id" validate:"required"`
	CustomerID    string          `json:"customer_id" db:"customer_id" validate:"required"`
	DueDate       string          `json:"due_date" db:"due_date" validate:"required"`
	Total         string          `json:"total" db:"total" validate:"required"`
	SubTotal      string          `json:"sub_total" db:"sub_total" validate:"required"`
	Discount      string          `json:"discount" db:"discount" validate:"required"`
	Reminder      json.RawMessage `json:"reminder" db:"reminder" validate:"required"`
	BankAccountID string          `json:"bank_account_id" db:"bank_account_id" validate:"required"`
	Status        string          `json:"status" db:"status" validate:"required"`
	CreatedBy     string          `json:"created_by" db:"created_by" validate:"required"`
	CreatedAt     string          `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt     string          `json:"updated_at" db:"updated_at" validate:"required"`
}

type InvoiceRes struct {
	ID            string          `json:"id" db:"id" validate:"required"`
	Name          string          `json:"name" db:"name" validate:"required"`
	Ref           string          `json:"ref" db:"ref" validate:"required"`
	CurrencyID    string          `json:"currency_id" db:"currency_id" validate:"required"`
	OutletID      *string         `json:"outlet_id" db:"outlet_id" validate:"required"`
	CustomerID    string          `json:"customer_id" db:"customer_id" validate:"required"`
	DueDate       string          `json:"due_date" db:"due_date" validate:"required"`
	Total         string          `json:"total" db:"total" validate:"required"`
	SubTotal      string          `json:"sub_total" db:"sub_total" validate:"required"`
	Discount      string          `json:"discount" db:"discount" validate:"required"`
	Reminder      json.RawMessage `json:"reminder" db:"reminder" validate:"required"`
	BankAccountID string          `json:"bank_account_id" db:"bank_account_id" validate:"required"`
	Status        string          `json:"status" db:"status" validate:"required"`
	CreatedBy     string          `json:"created_by" db:"created_by" validate:"required"`
	CreatedAt     string          `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt     string          `json:"updated_at" db:"updated_at" validate:"required"`
	Currency      json.RawMessage `json:"currency" db:"currency" validate:"required"`
	Customer      json.RawMessage `json:"customer" db:"customer" validate:"required"`
	BankAccount   json.RawMessage `json:"bank_account" db:"bank_account" validate:"required"`
	Outlet        json.RawMessage `json:"outlet" db:"outlet" validate:"required"`
}

type InvoiceBoq struct {
	ID        string  `json:"id" db:"id" validate:"required"`
	VariantID string  `json:"variant_id" db:"variant_id" validate:"required"`
	InvoiceID string  `json:"invoice_id" db:"invoice_id" validate:"required"`
	MeasureID string  `json:"measure_id" db:"measure_id" validate:"required"`
	Quantity  float64 `json:"quantity" db:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" db:"unit_price" validate:"required"`
	Total     float64 `json:"total" db:"total" validate:"required"`
	CreatedAt string  `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string  `json:"updated_at" db:"updated_at" validate:"required"`
}

type InvoiceReq struct {
	Name          string          `json:"name" binding:"required" validate:"required"`
	Ref           string          `json:"ref" binding:"required" validate:"required"`
	CurrencyID    string          `json:"currency_id" binding:"required" validate:"required"`
	CustomerID    string          `json:"customer_id" binding:"required" validate:"required"`
	OutletID      string          `json:"outlet_id" binding:"required" validate:"required"`
	DueDate       string          `json:"due_date" validate:"required" binding:"omitempty,datetime=2006-01-02"`
	Total         float64         `json:"total" binding:"required" validate:"required"`
	SubTotal      float64         `json:"sub_total" binding:"required" validate:"required"`
	Discount      *float64        `json:"discount" binding:"omitempty" validate:"omitempty"`
	Reminder      json.RawMessage `json:"reminder" binding:"required" validate:"required"`
	BankAccountID string          `json:"bank_account_id" binding:"required" validate:"required"`
	Status        string          `json:"status" binding:"required,oneof= draft active pending" validate:"required"`
	CreatedBy     string          `json:"created_by" binding:"required" validate:"required"`
	Boq           []InvoiceBoqReq `json:"boq" binding:"required" validate:"required"`
}

type InvoiceBoqReq struct {
	VariantID string  `json:"variant_id" binding:"required" validate:"required"`
	MeasureID string  `json:"measure_id" binding:"required" validate:"required"`
	Quantity  float64 `json:"quantity" binding:"required" validate:"required"`
	UnitPrice float64 `json:"unit_price" binding:"required" validate:"required"`
	Total     float64 `json:"total" binding:"required"  validate:"required"`
}

type UpdateInvoiceBoqReq struct {
	ID        string  `json:"id" binding:"omitempty" validate:"omitempty"`
	VariantID string  `json:"variant_id" binding:"required" validate:"required"`
	MeasureID string  `json:"measure_id" binding:"required" validate:"required"`
	Quantity  float64 `json:"quantity" binding:"required" validate:"required"`
	UnitPrice float64 `json:"unit_price" binding:"required" validate:"required"`
	Total     float64 `json:"total" binding:"required"  validate:"required"`
}

type InvoiceSummary struct {
	TotalInvoice json.RawMessage `json:"total_invoice" db:"total_invoice" validate:"required"`
	TotalDraft   json.RawMessage `json:"total_draft" db:"total_draft" validate:"required"`
	TotalSent    json.RawMessage `json:"total_sent" db:"total_sent" validate:"required"`
	TotalPaid    json.RawMessage `json:"total_paid" db:"total_paid" validate:"required"`
	TotalOverdue json.RawMessage `json:"total_overdue" db:"total_overdue" validate:"required"`
	TotalUnpaid  json.RawMessage `json:"total_unpaid" db:"total_unpaid" validate:"required"`
	TotalPending json.RawMessage `json:"total_pending" db:"total_pending" validate:"required"`
	TotalActive  json.RawMessage `json:"total_active" db:"total_active" validate:"required"`
}

type InvoiceStatusReq struct {
	Status    string `json:"status" validate:"required,oneof=active sent paid draft pending overdue"`
	UpdatedBy string `json:"updated_by" validate:"required"`
}
