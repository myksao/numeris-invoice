package domain

import (
	"encoding/json"
)

type Variant struct {
	ID          string `json:"id" db:"id" validate:"required"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"omitempty"`
	ItemID      string `json:"item_id" db:"item_id" validate:"required"`
	OutletID    string `json:"outlet_id" db:"outlet_id" validate:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" validate:"required"`
}

type VariantRes struct {
	ID          string           `json:"id" db:"id" validate:"required"`
	Name        string           `json:"name" db:"name" validate:"required"`
	Description string           `json:"description" db:"description" validate:"omitempty"`
	ItemID      string           `json:"item_id" db:"item_id" validate:"required"`
	OutletID    string           `json:"outlet_id" db:"outlet_id" validate:"required"`
	CreatedAt   string           `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   string           `json:"updated_at" db:"updated_at" validate:"required"`
	Item        *json.RawMessage `json:"item" db:"item" validate:"omitempty"`
	Outlet      *json.RawMessage `json:"outlet" db:"outlet" validate:"omitempty"`
	Inventory   *json.RawMessage `json:"inventory" db:"inventory" validate:"omitempty"`
}

type VariantItem struct {
	ID          string           `json:"id" db:"id" validate:"required"`
	Name        string           `json:"name" db:"name" validate:"required"`
	Description string           `json:"description" db:"description" validate:"omitempty"`
	ItemID      string           `json:"item_id" db:"item_id" validate:"required"`
	OutletID    string           `json:"outlet_id" db:"outlet_id" validate:"required"`
	CreatedAt   string           `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   string           `json:"updated_at" db:"updated_at" validate:"required"`
	Item        *json.RawMessage `json:"item" db:"item" validate:"omitempty"`
	Measure     *json.RawMessage `json:"measure" db:"measure" validate:"required"`
}

type VariantReq struct {
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"description" validate:"omitempty"`
	ItemID      string       `json:"item_id" validate:"required"`
	OutletID    string       `json:"outlet_id" validate:"required"`
	Measure     []MeasureReq `json:"measure" validate:"required"`
}

type VariantItemReq struct {
	ID string `uri:"id" validate:"required"`
}
