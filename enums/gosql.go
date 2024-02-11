package enums

const (
	QueryCheckProductExistsInPharmacyProduct = `
		SELECT EXISTS (
			SELECT 1 FROM pharmacy_drugs
			WHERE drug_id = ?
		)
	`

	CheckProductExistInCartUser = `
		SELECT EXISTS (
			SELECT 1 FROM carts
			WHERE drug_id = ? AND user_id = ?
		)
	`

	GetAdminPharmacyProducts = `
		SELECT * FROM pharmacies p JOIN pharmacy_drugs pd ON
		p.id = pd.pharmacy_id JOIN drugs d ON 
		pd.drug_id = d.id;
	`

	CheckPaymentFromUser = `
		SELECT EXISTS (
			SELECT 1 FROM orders
			WHERE user_id = ? AND payment_id = ?
		)
	`

	GetProducts = `
		SELECT *, MIN(pd.selling_unit), MAX(pd.selling_unit)
		FROM drugs d JOIN pharmacy_drugs pd ON
		d.id = pd.drug_id
		GROUP BY d.id`

	SeedingAuthorityUserRoles = `
		INSERT INTO authority_user_roles (user_id, role_id)
		SELECT id, 3 FROM users
		LIMIT 25;
	`

	SeedingAuthorityAdminRoles = `
		INSERT INTO authority_user_roles (user_id, role_id)
		SELECT id, 2 FROM users
		OFFSET 50 LIMIT 10;
	`

	SeedingAuthoritySuperAdminRoles = `
		INSERT INTO authority_user_roles (user_id, role_id)
		SELECT id, 1 FROM users
		OFFSET 60 LIMIT 1;
	`

	SeedingAuthorityDoctorRoles = `
		INSERT INTO authority_user_roles (user_id, role_id)
		SELECT id, 4 FROM users
		OFFSET 25 LIMIT 25;
	`

	GetRadius25km = `
		st_dwithin(ST_Makepoint(%s, %s)::geography, ST_MakePoint(a.longtitude, a.latitude)::geography, 25000)
	`

	QueryToGetPharmacyThatHasProducts = `
	SELECT p.id FROM pharmacies as p
		JOIN addresses as a ON p.address_id = a.id
	   	WHERE EXISTS (
		   SELECT 1 FROM pharmacy_drugs
				WHERE pharmacy_id = p.id AND drug_id IN (?)
   		) AND 
		st_dwithin(
			ST_Makepoint(?, ?)::geography, 
			ST_MakePoint(a.longtitude, a.latitude)::geography, 25000) AND p.deleted_at IS NULL;
	`

	GetCartForCheckout = `
	SELECT c.id as id, pd.id as pharmacy_drug_id, pd.drug_id as drug_id, d.name as drug_name, 
	       d.image as drug_image, pd.selling_unit as selling_unit, c.quantity as quantity,
			d.weight as weight, d.height as height, d.length as length, d.width as width,
	       c.quantity * pd.selling_unit as price
	FROM pharmacy_drugs as pd JOIN carts as c
		ON pd.drug_id = c.drug_id JOIN drugs as d ON d.id = pd.drug_id
		WHERE pd.pharmacy_id = ? AND c.user_id = ? 
		GROUP BY c.id, pd.id, pd.drug_id, d.name, d.image, pd.selling_unit, pd.drug_id, pd.stock, c.quantity,
		         d.weight, d.height, d.length, d.width
	`

	GetPharmacyThatHasMoreStock = `
		SELECT pd1.pharmacy_id
			FROM pharmacy_drugs pd1
			JOIN pharmacy_drugs pd2 ON pd1.id <> pd2.id
			WHERE pd2.pharmacy_id = ? AND pd2.drug_id = ? AND pd1.drug_id = ? AND pd1.stock > pd2.stock
		LIMIT 1;
	`

	GetProductStockLessThanCartQuantity = `
		SELECT pd.drug_id as drug_id, pd.stock as stock, c.quantity as quantity FROM pharmacy_drugs as pd JOIN carts as c
			ON pd.drug_id = c.drug_id
			WHERE pd.stock < c.quantity AND pd.pharmacy_id = ? AND c.user_id = ?;
	`

	InsertIntoStockMutation = `
		INSERT INTO stock_mutations (from_pharmacy_id, to_pharmacy_id, drug_id, quantity, status_mutation_id, created_at, updated_at)
		VALUES 
			(?, ?, ?, ?, ?, NOW(), NOW())
	`

	InsertIntoJournals = `
		INSERT INTO journals (drug_id, from_pharmacy_id, to_pharmacy_id, status, quantity, final_stock, created_at, updated_at) VALUES 
		(?, ?, ?, ?, ?, ?, NOW(), NOW())
	`
)
