package repo

const (
	fetchInventory  = `SELECT id FROM numeris.inventory WHERE variant_id = $1 and measure_id = $2 FOR UPDATE`
	createInventory = `INSERT INTO inventory (id, variant_id, measure_id) VALUES ($1, $2, $3) RETURNING id`
	addInventory    = `UPDATE inventory SET added_stock = added_stock + $1, stock_balance = opening_stock + (added_stock + $1 ) - issued_stock WHERE variant_id = $2 and measure_id = $3 RETURNING id`
	issueInventory  = `UPDATE inventory SET issued_stock = issued_stock + $1, stock_balance =  (opening_stock + added_stock) - (issued_stock + $1) WHERE variant_id = $2 and measure_id = $3 RETURNING id`
	openInventory   = `UPDATE inventory SET opening_stock = stock_balance, added_stock = 0 , issued_stock = 0 WHERE variant_id = $1 and measure_id = $2 AND opening_stock >= 0 AND opening_stock <= stock_balance RETURNING id`
	batchEntry      = `INSERT INTO batch (id, variant_id, measure_id,  entity_id, entity, state, stock) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
)
