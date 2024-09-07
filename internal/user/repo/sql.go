package repo

const (
	createUser             = `INSERT INTO "user" (id, name, username, password, ref, outlet_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	retrieveUserByID       = `SELECT * FROM "user" WHERE id = $1`
	retrieveUserByOutletID = `SELECT * FROM "user" WHERE outlet_id = $1 ORDER BY name ASC LIMIT $2 OFFSET $3`
)
