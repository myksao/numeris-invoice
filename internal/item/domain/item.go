package domain

type Item struct {
	ID          string `json:"id" db:"id" validate:"required"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"omitempty"`
	CategoryID  string `json:"category_id" db:"category_id" validate:"required"`
	SKU         string `json:"sku" db:"sku" validate:"required"`
	OutletID    string `json:"outlet_id" db:"outlet_id" validate:"required"`
	CreatedBy   string `json:"created_by" db:"created_by" validate:"required"`
	IsDeleted   bool   `json:"is_deleted" db:"is_deleted" validate:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" validate:"required"`
}

type ItemReq struct {
	Name        string `json:"name" binding:"required" validate:"required"`
	Description string `json:"description" binding:"required" validate:"required"`
	CategoryID  string `json:"category_id" binding:"required" validate:"required"`
	SKU         string `json:"sku" binding:"required" validate:"required"`
	OutletID    string `json:"outlet_id" binding:"required" validate:"required"`
	CreatedBy   string `json:"created_by" binding:"required" validate:"required"`
}

type ItemOutletReq struct {
	ID string `uri:"id" binding:"required" validate:"required"`
}
