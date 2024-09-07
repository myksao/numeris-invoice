package domain

type User struct {
	ID        string `json:"id" db:"id" validate:"required"`
	Name      string `json:"name" db:"name" validate:"required"`
	Username  string `json:"username" db:"username" validate:"required"`
	Password  string `json:"password" db:"password" validate:"required"`
	Ref       string `json:"ref" db:"ref" validate:"required"`
	OutletID  string `json:"outlet_id" db:"outlet_id" validate:"required"`
	CreatedAt string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at" db:"updated_at" validate:"required"`
}

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Ref      string `json:"ref" binding:"required"`
	OutletID string `json:"outlet_id" binding:"required"`
}

type UserOutletReq struct {
	Offset string `form:"offset" binding:"omitempty,number,gte=0"`
	Limit  string `form:"limit" binding:"omitempty,number,gte=1,required_with=offset"`
}
