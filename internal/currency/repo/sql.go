package repo


const (
	createCurrency = `INSERT INTO currency (id, name, code, symbol) VALUES ($1, $2, $3, $4) RETURNING id`
	retrieveCurrencies = `SELECT * FROM currency ORDER BY name ASC LIMIT $1 OFFSET $2` 
	createCurrencyMeasure = `INSERT INTO currency_measure (id, currency_id, measure_id, price) VALUES ($1, $2, $3, $4)`
)