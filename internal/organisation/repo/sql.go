package repo


const (
	createOrganisation = `INSERT INTO organisation (id, name, reference, address) VALUES ($1, $2, $3, $4) RETURNING id`
	retrieveOrganisationByID = `SELECT * FROM organisation WHERE id = $1`
) 