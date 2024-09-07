package repo

const (
	createCustomer       = `INSERT INTO customer (id, name, email, mobile_no, address, outlet_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	retrieveCustomerByOutletID   = `SELECT * FROM customer WHERE outlet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	retrieveCustomerByID = `SELECT * FROM customer WHERE id = $1`
)
