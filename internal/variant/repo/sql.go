package repo

const (
	createVariant       = `INSERT INTO numeris.variant (id, name, description, item_id, outlet_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	retrieveVariantByID = `SELECT 
		vari.id, vari.name, vari.description, vari.item_id, vari.outlet_id, vari.created_at, vari.updated_at,
		json_build_object('id', itm.id, 'name', itm.name, 'description', itm.description) as item,
		json_build_object('id', out.id, 'name', out.name, 'is_default', out.is_default) as outlet,
		json_build_object(
				'id', inv.id, 
				'opening_stock', inv.opening_stock,
				'added_stock', inv.added_stock,
				'issued_stock', inv.issued_stock,
				'stock_balance', inv.stock_balance
		) as inventory
	FROM variant vari 
	LEFT JOIN item itm ON vari.item_id = itm.id
	LEFT JOIN outlet out ON vari.outlet_id = out.id
	LEFT JOIN inventory inv ON vari.id = inv.variant_id
	WHERE vari.id = $1`
	retrieveVariantByItemID = `SELECT  vari.id, vari.name, vari.description, vari.item_id, vari.outlet_id, vari.created_at, vari.updated_at,
    							json_build_object('id', itm.id, 'name', itm.name, 'description', itm.description, 'category_id', itm.category_id, 'created_at', itm.created_at, 'updated_at', itm.updated_at) as item,
									(
										SELECT json_agg(
											json_build_object(
												'id', mea.id,
												'entity', mea.entity,
												'entity_id', mea.entity_id,
												'unit', mea.unit,
												'quantity', mea.quantity,
												'currency',
													(SELECT json_agg(
														json_build_object(
															'currency_id', cur.id,
															'price', cm.price
														)
													) FROM currency cur INNER JOIN numeris.currency_measure cm on cur.id = cm.currency_id and cm.measure_id = mea.id)
											)
										) FROM measure mea WHERE mea.entity_id = vari.id and mea.entity = 'variant'
									) as measure
											FROM variant vari
											INNER JOIN item itm ON vari.item_id = itm.id
											WHERE vari.item_id = $1 ORDER BY vari.created_at ASC LIMIT $2 OFFSET $3`

	createMeasure = `INSERT INTO measure (id, entity, entity_id, unit, quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	retrieveMeasureByID = `SELECT * FROM numeris.measure WHERE id = $1 and entity = 'variant'`

	retrieveMeasureByVariantID = `SELECT 
		mea.id, mea.entity, mea.entity_id, mea.unit, mea.quantity, mea.is_active, mea.created_at, mea.updated_at,
		json_agg(
			json_build_object(
				'currency_id', cur.id,
				'price', cmea.price
			)
		) as currency
	FROM measure mea
	LEFT JOIN currency_measure cmea ON cmea.measure_id = mea.id
	LEFT JOIN currency cur ON cmea.currency_id = cur.id
	WHERE mea.entity_id = $1 and mea.entity = 'variant'
	GROUP BY mea.id`

	createInventory = `INSERT INTO inventory (id, variant_id, measure_id) VALUES ($1, $2, $3) RETURNING id`
)
