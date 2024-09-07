package domain

type Inventory struct {
	ID           string `json:"id" db:"id" validate:"required"`
	VariantID    string `json:"variant_id" db:"variant_id" validate:"required"`
	OpeningStock string `json:"opening_stock" db:"opening_stock" validate:"required"`
	AddedStock   string `json:"added_stock" db:"added_stock" validate:"required"`
	IssuedStock  string `json:"issued_stock" db:"issued_stock" validate:"required"`
	StockBalance string `json:"stock_balance" db:"stock_balance" validate:"required"`
}

type InventoryReq struct {
	VariantID string `json:"variant_id" binding:"required" validate:"required"`
}

type InventoryProcess struct {
	VariantID string  `json:"variant_id"  validate:"required"`
	MeasureID string  `json:"measure_id"  validate:"required"`
	Quantity  float64 `json:"quantity" validate:"required" `
	EntityID  string  `json:"entity_id" validate:"omitempty" binding:"omitempty"`
	Entity    string  `json:"entity" validate:"omitempty" binding:"omitempty"`                       //"invoice" | ""
	State     string  `json:"state" validate:"required,oneof=none added returned issued waste hold"` //"none" | "added" | "issued" | "waste" | "hold" | "returned"
}
