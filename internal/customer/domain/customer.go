package domain

type Customer struct {
	ID        string `json:"id" db:"id" validate:"required"`
	Name      string `json:"name" db:"name" validate:"required"`
	MobileNo  string `json:"mobile_no" db:"mobile_no" validate:"required"`
	Address   string `json:"address" db:"address" validate:"required"`
	OutletID  string `json:"outlet_id" db:"outlet_id" validate:"required"`
	Email     string `json:"email" db:"email" validate:"required"`
	CreatedAt string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at" db:"updated_at" validate:"required"`
}

type CustomerReq struct {
	Name     string `json:"name" binding:"required" validate:"required"`
	MobileNo string `json:"mobile_no" binding:"required" validate:"required"`
	Address  string `json:"address" binding:"required" validate:"required"`
	OutletID string `json:"outlet_id" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required" validate:"required"`
}

type CustomerOutletReq struct {
	ID string `uri:"id" binding:"required" validate:"required"`
}
