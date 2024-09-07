package domain

type BankAccount struct {
	ID          string `json:"id" db:"id" validate:"required"`
	Name        string `json:"name" db:"name" validate:"required"`
	OutletID    string `json:"outlet_id" db:"outlet_id" validate:"required"`
	CurrencyID  string `json:"currency_id" db:"currency_id" validate:"required"`
	AccountNo   string `json:"account_no" db:"account_no" validate:"required"`
	RoutingNo   string `json:"routing_no" db:"routing_no" validate:"required"`
	AccountType string `json:"account_type" db:"account_type" validate:"required"`
	BankName    string `json:"bank_name" db:"bank_name" validate:"required"`
	IsActive    bool   `json:"is_active" db:"is_active" validate:"required"`
	IsDeleted   bool   `json:"is_deleted" db:"is_deleted" validate:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" validate:"required"`
}

type BankAccountReq struct {
	Name        string `json:"name" binding:"required" validate:"required"`
	OutletID    string `json:"outlet_id" binding:"required" validate:"required"`
	CurrencyID  string `json:"currency_id" binding:"required" validate:"required"`
	AccountNo   string `json:"account_no" binding:"required" validate:"required"`
	RoutingNo   string `json:"routing_no" binding:"required" validate:"required"`
	AccountType string `json:"account_type" binding:"required" validate:"required"`
	BankName    string `json:"bank_name" binding:"required" validate:"required"`
}

type BankAccountOutletReq struct {
	ID string `uri:"id" binding:"required" validate:"required"`
}
