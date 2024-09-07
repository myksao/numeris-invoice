package repo

const (
	createOutlet = `INSERT INTO outlet (id, name, is_default, org_id, address) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	retrieveOutletByID = `SELECT
    	out.id, out.name, out.is_default, 
		json_build_object(
			'id', org_id,
			'name', org.name
        ) as org, out.address, out.created_at, out.updated_at
		FROM outlet out
        JOIN organisation org ON out.org_id = org.id
		WHERE out.id = $1`

	retrieveOutletByOrgID = `SELECT * FROM outlet WHERE org_id = $1 ORDER BY name ASC LIMIT $2 OFFSET $3`
)
