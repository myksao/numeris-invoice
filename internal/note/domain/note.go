package domain

import "encoding/json"

type Note struct {
	ID        string          `json:"id" db:"id" validate:"required"`
	Entity    string          `json:"entity" db:"entity" validate:"required"`
	EntityID  string          `json:"entity_id" db:"entity_id" validate:"required"`
	Note      string          `json:"note" db:"note" validate:"required"`
	Command   json.RawMessage `json:"command" db:"command" validate:"required"`
	CreatedBy string          `json:"created_by" db:"created_by" validate:"required"`
	CreatedAt string          `json:"created_at" db:"created_at" validate:"required"`
	UpdatedAt string          `json:"updated_at" db:"updated_at" validate:"required"`
}

type NoteReq struct {
	Entity    string          `json:"entity" binding:"required" validate:"required"`
	EntityID  string          `json:"entity_id" binding:"required" validate:"required"`
	Note      string          `json:"note" binding:"required" validate:"required"`
	Command   json.RawMessage `json:"command" binding:"required" validate:"required"`
	CreatedBy string          `json:"created_by" binding:"required" validate:"required"`
}

type NoteEntityIDReq struct {
	Entity string `uri:"entity" binding:"required" validate:"required"`
	ID     string `uri:"id" binding:"required" validate:"required"`
}

type NoteEntityReq struct {
	Entity string `uri:"entity" binding:"required" validate:"required"`
}
