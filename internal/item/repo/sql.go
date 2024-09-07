package repo

const (
	createItem             = `INSERT INTO "item" (id, name, description, category_id, sku, outlet_id, created_by, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, false) RETURNING id`
	retrieveItemByID       = `SELECT * FROM "item" WHERE id = $1`
	retrieveItemByOutletID = `SELECT * FROM "item" WHERE outlet_id = $1 LIMIT $2 OFFSET $3`
)
