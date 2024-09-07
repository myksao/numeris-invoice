package domain

type Organisation struct {
	ID        string `json:"id" db:"id" validate:"required"`
	Name      string `json:"name" db:"name" validate:"required"`
	Reference string `json:"reference" db:"reference" validate:"required"`
	Address   string `json:"address" db:"address" validate:"required"`
	CreatedAt string `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string `json:"updated_at" db:"updated_at" validate:"required"`
}

type OrganisationReq struct {
	Name      string `json:"name" binding:"required" validate:"required"`
	Reference string `json:"reference" binding:"required" validate:"required"`
	Address   string `json:"address" binding:"required" validate:"required"`
}

type CreateOrgRes struct {
	ID string `json:"id"`
}
