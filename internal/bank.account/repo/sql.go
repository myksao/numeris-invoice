package repo

const (
	createBankAccount             = `INSERT INTO bank_account (id, name, outlet_id, account_no, routing_no, account_type, bank_name, currency_id, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, false) RETURNING id`
	retrieveBankAccountByOutletID = `SELECT * FROM bank_account WHERE outlet_id = $1 ORDER BY name ASC LIMIT $2 OFFSET $3`
	retrieveBankAccountByID       = `SELECT * FROM bank_account WHERE id = $1`
)
