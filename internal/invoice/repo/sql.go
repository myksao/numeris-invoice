package repo

const (
	createInvoice = `INSERT INTO invoice (id, name, ref, due_date, total, status, reminder, created_by, bank_account_id, currency_id, customer_id, sub_total, discount, outlet_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`

	createInvoiceBoq = `INSERT INTO invoice_boq (id, variant_id, invoice_id, measure_id, quantity, unit_price, total) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	updateInvoiceBoq = `UPDATE invoice_boq SET quantity = $1, unit_price = $2, total = $3 WHERE id = $4`

	retrieveInvoiceByID = `SELECT 
		inv.id, inv.name, inv.ref, 
		inv.due_date, inv.total, inv.status, inv.reminder, inv.sub_total, inv.discount, inv.created_by, inv.created_at, inv.updated_at,
		json_build_object(
			'id', cur.id,
			'name', cur.name,
			'code', cur.code,
			'symbol', cur.symbol
		) as currency,
		json_build_object(
			'id', cus.id,
			'name', cus.name,
			'email', cus.email,
			'phone', cus.mobile_no,
			'address', cus.address
		) as customer,
		json_build_object(
			'id', ba.id,
			'name', ba.name,
			'account_no', ba.account_no,
			'routing_no', ba.routing_no,
			'account_type', ba.account_type,
			'bank_name', ba.bank_name
		) as bank_account,
		json_build_object(
			'id', out.id,
			'name', out.name
		) as outlet
	FROM invoice inv 
	LEFT JOIN currency cur ON inv.currency_id = cur.id
	LEFT JOIN customer cus ON inv.customer_id = cus.id
	LEFT JOIN bank_account ba ON inv.bank_account_id = ba.id
	LEFT JOIN outlet out ON inv.outlet_id = out.id
	WHERE inv.id = $1`

	retrieveInvoiceBoqByInvoiceID = `SELECT * FROM invoice_boq WHERE invoice_id = $1`

	retrieveInvoiceSummary = `SELECT 
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	)  FROM invoice WHERE outlet_id = $1) as total_invoice,
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'draft' and outlet_id = $1) as total_draft, 
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'sent' and outlet_id = $1)  as total_sent,
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'paid' and outlet_id = $1 ) as total_paid,
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'overdue' and outlet_id = $1)  as total_overdue,
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'unpaid'  and outlet_id = $1) as total_unpaid,
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'pending'  and outlet_id = $1 )as total_pending,
	(SELECT json_build_object(
		'entry', COUNT(*),
		'amount', SUM(total)
	) FROM invoice WHERE status = 'active'  and outlet_id = $1) as total_active
	`

	updateInvoiceStatus = `UPDATE invoice SET status = $1 WHERE id = $2 RETURNING id`

	triggerOverDueInvoice = `select cron.schedule (
    'upate-invoice-status', 
    $1,
    $$ update invoice set status = 'overdue' where due_date < now() and status = 'sent'; $$
);`
)
