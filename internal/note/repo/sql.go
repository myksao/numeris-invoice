package repo

const (
	createNote             = `INSERT INTO numeris.note (id, entity_id, entity, note, created_by, command) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	retrieveNoteByEntityID = `SELECT * FROM note WHERE entity_id = $1 and entity = $2`
	retrieveNoteByEntity   = `SELECT * FROM note WHERE  entity = $1`
	retrieveNoteByID       = `SELECT * FROM note WHERE id = $1`
)
