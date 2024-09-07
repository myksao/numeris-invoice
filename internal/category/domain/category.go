package domain

type Category struct {
	ID          string  `json:"id" db:"id" validate:"required"`
	Name        string  `json:"name" db:"name" validate:"required"`
	Description *string `json:"description" db:"description" validate:"omitempty"`
	OutletID    string  `json:"outlet_id" db:"outlet_id" validate:"required"`
	IsDeleted   bool    `json:"is_deleted" db:"is_deleted" validate:"required"`
	CreatedAt   string  `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at" validate:"required"`
}

type CategoryReq struct {
	Name     string `json:"name" binding:"required" validate:"required"`
	OutletID string `json:"outlet_id" binding:"required" validate:"required"`
}

type CategoryOutletReq struct {
	ID string `uri:"id" binding:"required" validate:"required"`
}
