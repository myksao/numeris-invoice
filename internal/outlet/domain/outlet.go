package domain

import "encoding/json"

type Outlet struct {
	ID        string           `json:"id" db:"id" validate:"required"`
	Name      string           `json:"name" db:"name" validate:"required"`
	IsDefault bool             `json:"is_default" db:"is_default" validate:"required"`
	Address   string           `json:"address" db:"address" validate:"required"`
	OrgID     string           `json:"org_id" db:"org_id" validate:"required"`
	Org       *json.RawMessage `json:"org" validate:"required"`
	CreatedAt string           `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string           `json:"updated_at" db:"updated_at" validate:"required"`
}

type OutletReq struct {
	Name      string `json:"name" binding:"required" validate:"required"`
	IsDefault bool   `json:"is_default" binding:"required" validate:"required"`
	Address   string `json:"address" binding:"required" validate:"required"`
	OrgID     string `json:"org_id" binding:"required" validate:"required"`
}

type CreateOutletRes struct {
	ID string `json:"id"`
}
