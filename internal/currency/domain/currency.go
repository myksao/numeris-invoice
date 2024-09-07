package domain

type Currency struct {
	ID     string `json:"id" db:"id" validate:"required"`
	Name   string `json:"name" db:"name" validate:"required"`
	Code   string `json:"code" db:"code" validate:"required"`
	Symbol string `json:"symbol" db:"symbol" validate:"required"`
}

type CurrencyReq struct {
	Name   string `json:"name" binding:"required" validate:"required"`
	Code   string `json:"code" binding:"required" validate:"required"`
	Symbol string `json:"symbol" binding:"required" validate:"required"`
}

type CurrencyMeasure struct {
	ID         string `json:"id" db:"id" validate:"required"`
	CurrencyID string `json:"currency_id" db:"currency_id" validate:"required"`
	MeasureID  string `json:"measure_id" db:"measure_id" validate:"required"`
}

type CurrencyMeasureReq struct {
	CurrencyID string `json:"currency_id" binding:"required" validate:"required"`
	MeasureID  string `json:"measure_id" binding:"required" validate:"required"`
	Price      string `json:"price" binding:"required" validate:"required"`
}

type MeasureCurrencyRes struct {
	Currency Currency `json:"currency" db:"currency" validate:"required"`
	Price    string   `json:"price" db:"price" validate:"required"`
}
