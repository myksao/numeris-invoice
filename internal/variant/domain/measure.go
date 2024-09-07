package domain

import (
	"encoding/json"
)

type Measure struct {
	ID        string `json:"id" db:"id" validate:"required"`
	Entity    string `json:"entity" db:"entity" validate:"required"`
	EntityID  string `json:"entity_id" db:"entity_id" validate:"required"`
	Unit      string `json:"unit" db:"unit" validate:"required"`
	Quantity  string `json:"quantity" db:"quantity" validate:"required"`
	IsActive  bool   `json:"is_active" db:"is_active" validate:"required"`
	CreatedAt string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at" db:"updated_at" validate:"required"`
}

type MeasureReq struct {
	Entity   string            `json:"entity" binding:"omitempty" validate:"required"`
	EntityID string            `json:"entity_id" binding:"omitempty" validate:"required"`
	Unit     string            `json:"unit" binding:"required" validate:"required"`
	Quantity string            `json:"quantity" binding:"required" validate:"required"`
	Currency []MeasureCurrency `json:"currency" binding:"required" validate:"required"`
}

type MeasureCurrency struct {
	CurrencyID string `json:"currency_id" db:"currency_id" validate:"required"`
	Price      string `json:"price" db:"price" validate:"required"`
}

type MeasureRes struct {
	ID        string          `json:"id" db:"id" validate:"required"`
	Entity    string          `json:"entity" db:"entity" validate:"required"`
	EntityID  string          `json:"entity_id" db:"entity_id" validate:"required"`
	Unit      string          `json:"unit" db:"unit" validate:"required"`
	Quantity  string          `json:"quantity" db:"quantity" validate:"required"`
	IsActive  bool            `json:"is_active" db:"is_active" validate:"required"`
	CreatedAt string          `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string          `json:"updated_at" db:"updated_at" validate:"required"`
	Currency  json.RawMessage `json:"currency" db:"currency" validate:"required"`
}
