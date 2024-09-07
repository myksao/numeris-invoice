package repo

const (
	createCategory               = `INSERT INTO category (id, name, outlet_id, is_deleted) VALUES ($1, $2, $3, false) RETURNING id`
	retrieveCategoryByID         = `SELECT * FROM category WHERE id = $1`
	retrieveCategoriesByOutletID = `SELECT * FROM category WHERE outlet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
)
