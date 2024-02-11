CREATE DATABASE healthcare_db;
DROP DATABASE healthcare_db;
CREATE EXTENSION postgis;
SELECT * FROM rooms;
SELECT * FROM chats WHERE room_id = 1;
SELECT * FROM doctors WHERE id = 1;

INSERT INTO chats (room_id, sender_id, message, created_at, updated_at)
VALUES
    (1, 1, 'Hallo doc', NOW(), NOW()),
    (1, 210, 'Ada kendala apa mas', NOW(), NOW()),
    (1, 1, 'Aku sayang kamu doc', NOW(), NOW()),
    (1, 1, 'Minta no. hp', NOW(), NOW()),
    (1, 210, 'Awww :3', NOW(), NOW());

SELECT * FROM doctors;
SELECT * FROM doctors WHERE id = 2;
SELECT * FROM users WHERE email LIKE 'super.admin%';
SELECT * FROM authority_user_roles;
SELECT * FROM users WHERE email = 'randysteven12@gmail.com';
INSERT INTO sub_districts (name, created_at, updated_at)
VALUES
    ('Sub district test', NOW(), NOW());

INSERT INTO sub_districts (name)
VALUES 
    ('Sub Distrcit test');

SELECT * FROM sub_districts;
SELECT * FROM addresses;
SELECT * FROM pharmacies WHERE admin_pharmacy_id = 5;

SELECT * FROM orders;
ALTER TABLE orders
ADD COLUMN shipment_name VARCHAR;


SELECT * FROM admin_pharmacies WHERE id = 58;
SELECT * FROM operationals;

SELECT * FROM register_account_tokens;
INSERT INTO users (name, email, password, is_verify, phone_number, online_status, photo, created_at, updated_at)
VALUES 
    ('User Test 1', 'user.test@gmail.com', 'test_1234', true, '+628123456789', 'online', '', NOW(), NOW());

INSERT INTO users (name, email, password, is_verify, phone_number, online_status, photo, created_at, updated_at)
VALUES 
    ('User Test 2', 'user.test2@gmail.com', 'test_1234', true, '+628123456711', 'online', '', NOW(), NOW());

SELECT "drugs".name,"drugs".image,"drugs".unit_in_pack,c.name as
    category,m.name as manufacture,max(pd.selling_unit) as max_price,min(pd.selling_unit) as min_price,
             "drugs".updated_at
FROM "drugs" JOIN categories c ON "drugs".category_id = c.id
    FULL OUTER JOIN manufactures m ON "drugs".manufacture_id = m.id
    FULL OUTER JOIN pharmacy_drugs pd ON "drugs"."id" = pd.drug_id
    FULL OUTER JOIN pharmacies p ON pd.pharmacy_id = p.id
    FULL OUTER JOIN addresses a ON p.address_id = a.id
WHERE (drugs.name ILIKE '%%' OR drugs.content ILIKE '%%' OR drugs.generic_name ILIKE '%%'
           OR drugs.description ILIKE '%%') AND "drugs"."deleted_at" IS NULL
GROUP BY "drugs".name,"drugs".image,"c"."name","drugs".unit_in_pack,"m"."name", "drugs".updated_at
ORDER BY drugs.updated_at desc LIMIT 10;

SELECT `"drugs".ID`,
    `"drugs".name`,
    `"drugs".generic_name`,
    `"drugs".content`,
    `m.id as manufacture_id`,
    `c.id as category_id`,
    `f.ID as drug_form_id`,
    `f.name as form`,
    `"drugs".unit_in_pack`,
    `"drugs".weight`,
    `"drugs".height`,
    `"drugs".length`,
    `"drugs".width`,
    `"drugs".image`,`"drugs".unit_in_pack`,`c.name as category`,
    `m.name as manufacture`,`drugs.updated_at`,`max(pd.selling_unit) as max_selling_unit`,`min(pd.selling_unit) as min_selling_unit`
FROM "drugs" JOIN categories c ON "drugs".category_id = c.id
    FULL OUTER JOIN manufactures m ON "drugs".manufacture_id = m.id
    FULL OUTER JOIN pharmacy_drugs pd ON "drugs"."id" = pd.drug_id
    FULL OUTER JOIN pharmacies p ON pd.pharmacy_id = p.id
    FULL OUTER JOIN forms f ON "drugs".form_id = f.id
    FULL OUTER JOIN addresses a ON p.address_id = a.id
WHERE (drugs.name ILIKE '%%' OR drugs.content ILIKE '%%' OR
       drugs.generic_name ILIKE '%%' OR drugs.description ILIKE '%%') AND
    "drugs"."deleted_at" IS NULL
GROUP BY
    "drugs".ID,
    "drugs".name,
    "drugs".generic_name,
    "drugs".content,
    "drugs".manufacture_id,
    m.id,
    c.id,
    f.ID,
    f.name,
    "drugs".unit_in_pack,
    "drugs".weight,
    "drugs".height,
    "drugs".length,
    "drugs".width,
    "drugs".image,"drugs".unit_in_pack,c.name,
    m.name,drugs.updated_at;

ALTER TABLE addresses
ADD is_default BOOLEAN;

SELECT * FROM addresses;
SELECT * FROM drugs;
SELECT * FROM pharmacy_drugs;



SELECT * FROM users;
SELECT * FROM addresses;
SELECT * FROM orders;
SELECT * FROM manufactures;
SELECT * FROM forms;
SELECT * FROM pharmacies;
SELECT * FROM categories;
SELECT * FROM drugs;
SELECT * FROM pharmacies WHERE id = 31;
SELECT * FROM shipments;
SELECT * FROM pharmacy_drugs;

-- CREATE OR REPLACE PROCEDURE get_random_drugs()
-- BEGIN
--     SELECT id FROM drugs ORDER BY RANDOM() LIMIT 1;
-- END
--
-- SELECT get_random_drugs;
--
-- CREATE OR REPLACE PROCEDURE insert_pharmacy_drugs()
-- LANGUAGE plpgsql as $$
-- BEGIN
--
--     FOR i IN 1..25 LOOP
--
--         INSERT INTO pharmacy_drugs;
--
-- END;

INSERT INTO pharmacy_drugs 
SELECT * FROM pharmacy


SELECT * FROM payments WHERE id = 1;
INSERT INTO payments (file, status, created_at, updated_at)
VALUES
    ('', 'Unpaid', NOW(), NOW()),
    ('', 'Decline', NOW(), NOW()),
    ('', 'Paid', NOW(), NOW());

SELECT * FROM admin_pharmacies JOIN users ON
admin_pharmacies.user_id = users.id;

INSERT INTO orders (user_id, pharmacy_id,order_date,address_id,shipment_id,payment_id,order_status,shipping_cost,total_drugs_amount,total_amount,created_at,updated_at)
VALUES
    (1, 1, NOW(), 1, 2, 1, 'FAILED', 0, 0, 0, NOW(), NOW()),
    (1, 1, NOW(), 1, 2, 1, 'FAILED', 0, 0, 0, NOW(), NOW());

INSERT INTO pharmacy_drugs (pharmacy_id, drug_id, stock, created_at, updated_at)
VALUES
    (1, 1, 100, NOW(), NOW()),
    (2, 1, 109, NOW(), NOW()),
    (1, 2, 111, NOW(), NOW()),
    (2, 2, 123, NOW(), NOW()),
    (2, 3, 231, NOW(), NOW());

SELECT SUM(stock) FROM pharmacy_drugs WHERE drug_id = 1;

SELECT EXISTS (
    SELECT 1 FROM pharmacy_drugs
    WHERE drug_id = 5
)::INT;

SELECT EXISTS (
    SELECT 1 FROM carts
    WHERE drug_id = 1 AND user_id = 2
);

SELECT * FROM carts;
DELETE FROM carts WHERE id = 4;

SELECT * FROM users;
SELECT * FROM carts;
SELECT * FROM pharmacy_drugs;

INSERT INTO drugs (name, content, generic_name, 
    description, manufacture_id, form_id, category_id, 
    unit_in_pack, weight, height, 
    length, width, image, created_at, updated_at)
VALUES
    (
        'Test Product B',
        'Komposisi Lah ini',
        'Paracetamol',
        'Obat batuk',
        1,
        2,
        1,
        'Tablet',
        1,
        2,
        3,
        4,
        'image',
        NOW(),
        NOW()
    ),
    (
        'Test Product C',
        'Komposisi Lah ini',
        'Paracetamol',
        'Obat batuk',
        2,
        2,
        1,
        'Syrup Obat',
        1,
        2,
        3,
        4,
        'image',
        NOW(),
        NOW()
    ),
    (
        'Test Product D',
        'Komposisi Lah ini',
        'Paracetamol',
        'Obat batuk',
        1,
        2,
        1,
        'Tablet',
        1,
        2,
        3,
        4,
        'image',
        NOW(),
        NOW()
    ),
    (
        'Test Product G',
        'Komposisi Lah ini',
        'Paracetamol',
        'Obat batuk',
        2,
        2,
        1,
        'Syrup Obat',
        1,
        2,
        3,
        4,
        'image',
        NOW(),
        NOW()
    );

SELECT * FROM pharmacies p JOIN addresses a 
ON p.address_id = a.id;

UPDATE addresses SET longtitude = -6.200000
WHERE id = 5;

CREATE OR REPLACE PROCEDURE get_list_25km(curr_long numeric, curr_lat numeric) language plpgsql
    BEGIN
    RETURN QUERY 
        SELECT * FROM addresses
        WHERE ST_DWithin('POINT(-4.6314 54.0887)'::geography, ST_MakePoint(longtitude, latitude), 8046.72);
    END;
WHERE id = 2;
UPDATE addresses SET longtitude = -6.200000
WHERE id = 3;
SELECT * FROM addresses;
INSERT INTO addresses(
    detail, province_id, city_id, subdistrict_id, longtitude, latitude, created_at, updated_at, user_id
    )
VALUES
    ('Jln. Penderitaan', 1, 1, 1, 106.8296488, 6.2273378, NOW(), NOW(), 2),
    ('Jln. Penderitaan', 1, 1, 1, 106.8296488, 6.3561, NOW(), NOW(), 3);

SELECT * FROM "admin_pharmacies" JOIN admin_pharmacy_jobs jobs ON admin_pharmacies.id = jobs.admin_pharmacy_id JOIN pharmacies.id = jobs.pharmacy_id LEFT JOIN "users" "User" ON "admin_pharmacies"."user_id" = "User"."id" AND "User"."deleted_at" IS NULL JOIN authority_user_roles aur ON "User".id = CAST(aur.user_id AS INTEGER) WHERE aur.role_id = '2' AND "admin_pharmacies"."deleted_at" IS NULL;

SELECT * FROM "admin_pharmacies" FULL OUTER JOIN admin_pharmacy_jobs jobs ON admin_pharmacies.id = jobs.admin_pharmacy_id JOIN pharmacies ON pharmacies.id = jobs.pharmacy_id JOIN "users" "User" ON "admin_pharmacies"."user_id" = "User"."id" AND "User"."deleted_at" IS NULL JOIN authority_user_roles aur ON "User".id = CAST(aur.user_id AS INTEGER) WHERE aur.role_id = '2' AND "admin_pharmacies"."deleted_at" IS NULL;

SELECT * FROM pharmacies;

CREATE TABLE "admin_pharmacy_job" (
    "admin_pharmacy_id" bigint,
    "id" bigserial NOT NULL,
    PRIMARY KEY ("admin_pharmacy_id","id"),
    CONSTRAINT "fk_admin_pharmacy_job_admin_pharmacy_job" FOREIGN KEY ("id") 
    REFERENCES "admin_pharmacy_jobs"("pharmacy_id"),CONSTRAINT "fk_admin_pharmacy_job_admin_pharmacy" FOREIGN KEY ("admin_pharmacy_id") REFERENCES "admin_pharmacies"("id"))

CREATE TABLE "admin_pharmacy_jobs" (
    "id" bigserial,"admin_pharmacy_id" bigint NOT NULL,
    "pharmacy_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL,"updated_at" timestamptz NOT NULL,"deleted_at" timestamptz,PRIMARY KEY ("id"),CONSTRAINT "fk_admin_pharmacy_jobs_admin_pharmacy" FOREIGN KEY ("admin_pharmacy_id") REFERENCES "admin_pharmacies"("id"),CONSTRAINT "fk_admin_pharmacy_jobs_pharmacy" FOREIGN KEY ("admin_pharmacy_id") REFERENCES "pharmacies"("id"))

CREATE TABLE "admin_pharmacy_jobs" ("id" bigserial,"admin_pharmacy_id" bigint NOT NULL,"pharmacy_id" bigint NOT NULL,"created_at" timestamptz NOT NULL,"updated_at" timestamptz NOT NULL,"deleted_at" timestamptz,PRIMARY KEY ("id"),CONSTRAINT "fk_admin_pharmacy_jobs_admin_pharmacy" FOREIGN KEY ("admin_pharmacy_id") REFERENCES "admin_pharmacies"("id"),CONSTRAINT "fk_admin_pharmacy_jobs_pharmacy" FOREIGN KEY ("pharmacy_id") REFERENCES "pharmacies"("id"))


CREATE DATABASE healthcare_db_test;

SELECT EXISTS (
    SELECT * FROM pharmacy_drugs
    WHERE drug_id = 1 AND pharmacy_id = 1
);
SELECT * FROM payments;
SELECT * FROM drugs;
SELECT * FROM pharmacies;

<<<<<<< HEAD
SELECT 
    d.id as 'id', 
    d.name as 'name',
    d.generic_name as 'generic_name',
    d.description as 'description'
    MIN(pd.selling_unit) as 'min_selling_unit', 
    MAX(pd.selling_unit) as 'max_selling_unit', SUM(pd.stock)
		FROM drugs d JOIN pharmacy_drugs pd ON
		d.id = pd.drug_id
		GROUP BY d.name;

SELECT * FROM drugs;

    SELECT p.id, a.longtitude, a.latitude, pd.selling_unit FROM pharmacies as p
         JOIN addresses as a ON p.address_id = a.id
                                          JOIN pharmacy_drugs as pd
                                          ON pd.pharmacy_id = p.id
        WHERE EXISTS (
            SELECT 1 FROM pharmacy_drugs
                     WHERE pharmacy_id = p.id AND drug_id IN (1)
        ) AND st_dwithin(ST_Makepoint(107.11926383462607, -41.78750959852032)::geography, ST_MakePoint(a.longtitude, a.latitude)::geography, 25000);

INSERT INTO addresses (detail, province_id, city_id, user_id, longtitude, latitude, created_at, updated_at, deleted_at, is_default)
VALUES
    ('Jln. User', 1, 1, 2, 107.11926383462607, -41.78750959852032, now(), now(), NULL, true);

SELECT * FROM users;

SELECT * FROM addresses AS a
WHERE st_dwithin(ST_Makepoint(107.11926383462607, -41.78750959852032)::geography, ST_MakePoint(a.longtitude, a.latitude)::geography, 25000)
ORDER BY a.longtitude, a.latitude LIMIT 1;

SELECT * FROM pharmacies;
INSERT INTO pharmacies (
    address_id, 
    name, 
    pharmaciest_name, 
    license_number, 
    phone_number, 
    admin_pharmacy_id, 
    created_at, 
    updated_at)
VALUES
    (57, 'Pharmacy Tidak Ta', 'Lalala', '99999999', '+621234561231', 1, NOW(), NOW());

INSERT INTO pharmacy_drugs (pharmacy_id, drug_id, selling_unit, stock, status, created_at, updated_at)
VALUES
    (46, 1, 12000, 199, 'active', NOW(), NOW()),
    (46, 2, 11500, 101, 'active', NOW(), NOW()),
    (46, 3, 12500, 121, 'active', NOW(), NOW());

    SELECT p.id FROM pharmacies as p
=======
SELECT * FROM pharmacies as p
WHERE EXISTS (
    SELECT 1 FROM pharmacy_drugs
             WHERE pharmacy_id = p.id AND drug_id IN (1, 2)
) LIMIT 1;

SELECT * FROM pharmacy_drugs;

SELECT * FROM pharmacies p JOIN addresses a ON p.address_id = a.id;
SELECT * FROM drugs;

SELECT * FROM users u JOIN addresses a ON u.id = a.user_id;
SELECT * FROM addresses;

SELECT * FROM pharmacies as p
>>>>>>> dev
         JOIN addresses as a ON p.address_id = a.id
        WHERE EXISTS (
            SELECT 1 FROM pharmacy_drugs
                     WHERE pharmacy_id = p.id AND drug_id IN (1, 2)
<<<<<<< HEAD
        ) LIMIT 1;



SELECT * FROM pharmacy_drugs;
SELECT * FROM carts;
SELECT * FROM addresses;

-- 1. Latitude: -41.787272, Longitude: 107.119107
-- 2. Latitude: -41.788410, Longitude: 107.118543
-- 3. Latitude: -41.786932, Longitude: 107.118956
-- 4. Latitude: -41.787918, Longitude: 107.119812
-- 5. Latitude: -41.788264, Longitude: 107.118308
-- 6. Latitude: -41.786545, Longitude: 107.118689
-- 7. Latitude: -41.786799, Longitude: 107.119477
-- 8. Latitude: -41.788127, Longitude: 107.119402
-- 9. Latitude: -41.786996, Longitude: 107.118203
-- 10. Latitude: -41.787611, Longitude: 107.119743

INSERT INTO addresses (detail, province_id, city_id, longtitude, latitude, created_at, updated_at, is_default)
VALUES 
    ('Jln. 1', 2, 1, 107.119107, -41.787272, NOW(), NOW(), false),
    ('Jln. 2', 2, 1, 107.118543, -41.788410, NOW(), NOW(), false),
    ('Jln. 3', 2, 1, 107.118956, -41.786932, NOW(), NOW(), false),
    ('Jln. 4', 2, 1, 107.119812, -41.787918, NOW(), NOW(), false),
    ('Jln. 5', 2, 1, 107.118308, -41.788264, NOW(), NOW(), false),
    ('Jln. 6', 2, 1, 107.118689, -41.786545, NOW(), NOW(), false),
    ('Jln. 7', 2, 1, 107.119477, -41.786799, NOW(), NOW(), false),
    ('Jln. 8', 2, 1, 107.119402, -41.788127, NOW(), NOW(), false),
    ('Jln. 9', 2, 1, 107.118203, -41.786996, NOW(), NOW(), false),
    ('Jln. 10', 2, 1, 107.119743, -41.787611, NOW(), NOW(), false);


BEGIN TRANSACTION buy_products;

    SELECT p.id FROM pharmacies as p
         JOIN addresses as a ON p.address_id = a.id
        WHERE EXISTS (
            SELECT 1 FROM pharmacy_drugs
                     WHERE pharmacy_id = p.id AND drug_id IN (1, 2)
        ) LIMIT 1;

    
SELECT * FROM categories WHERE id = 7;
SELECT * FROM addresses;

SELECT * FROM orders;
SELECT * FROM carts WHERE user_id = 2;
SELECT * FROM carts;

SELECT * FROM pharmacy_drugs as pd JOIN carts as c
ON pd.drug_id = c.drug_id
WHERE pd.pharmacy_id = 46 AND c.user_id = 2 AND pd.drug_id IN (1, 2);

	SELECT pd.drug_id, d.name, d.image, pd.selling_unit, pd.stock, c.quantity, c.quantity * pd.selling_unit FROM pharmacy_drugs as pd JOIN carts as c
		ON pd.drug_id = c.drug_id JOIN drugs as d ON d.id = pd.drug_id
		WHERE pd.pharmacy_id = 46 AND c.user_id = 2
		GROUP BY pd.drug_id, d.name, d.image, pd.selling_unit, pd.drug_id, pd.stock, c.quantity;

SELECT * FROM drugs;
UPDATE drugs SET weight = 100 WHERE id = 2;

SELECT * FROM users WHERE email = 'BBltDll@ICivapO.org';


SELECT * FROM orders o JOIN order_details od ON o.id = od.order_id
WHERE o.id = 13;

SELECT pd.drug_id FROM pharmacy_drugs as pd JOIN carts as c
ON pd.drug_id = c.drug_id
WHERE pd.pharmacy_id = 46 AND c.quantity < pd.stock;

SELECT pd1.id ,pd2.id, COALESCE(pd1.stock > pd2.stock, true)
FROM pharmacy_drugs as pd1
RIGHT JOIN pharmacy_drugs as pd2
USING(id);

SELECT pd1.pharmacy_id, pd1.drug_id, pd1.stock
FROM pharmacy_drugs pd1
JOIN pharmacy_drugs pd2 ON pd1.id <> pd2.id
WHERE pd2.pharmacy_id = 46 AND pd2.drug_id = 2 AND pd1.drug_id = 2 AND pd1.stock > pd2.stock
LIMIT 1;

SELECT * FROM pharmacy_drugs pd JOIN carts c
ON pd.drug_id = c.drug_id
WHERE pd.stock < c.quantity;

SELECT * FROM pharmacy_drugs WHERE drug_id = 2;

SELECT * FROM pharmacy_drugs WHERE drug_id = 3;

SELECT * FROM pharmacy_drugs WHERE pharmacy_id = 28;
SELECT * FROM pharmacy_drugs WHERE pharmacy_id = 26;
SELECT * FROM pharmacy_drugs WHERE pharmacy_id = 46;
SELECT * FROM pharmacy_drugs WHERE pharmacy_id IN (26, 28);

SELECT pd.drug_id as drug_id, pd.stock as stock, c.quantity as quantity FROM pharmacy_drugs as pd JOIN carts as c
			ON pd.drug_id = c.drug_id
			WHERE pd.stock < c.quantity AND pd.pharmacy_id = 46 AND c.user_id = 2;

UPDATE pharmacy_drugs
SET stock = 86 WHERE pharmacy_id = 46 AND drug_id = 2;

UPDATE pharmacy_drugs
SET stock = 10 WHERE pharmacy_id = 46 AND drug_id = 1;

UPDATE pharmacy_drugs
SET stock = 999 WHERE pharmacy_id = 29 AND drug_id = 2;

UPDATE pharmacy_drugs
SET stock = 999 WHERE pharmacy_id = 28 AND drug_id = 2;

UPDATE pharmacy_drugs
SET stock = 999 WHERE pharmacy_id = 29 AND drug_id = 1;

UPDATE pharmacy_drugs
SET stock = 999 WHERE pharmacy_id = 28 AND drug_id = 1;


SELECT * FROM pharmacy_drugs WHERE pharmacy_id IN (46, 28) AND drug_id = 2 AND drug_id = 1;
SELECT * FROM pharmacy_drugs WHERE pharmacy_id IN (46, 28, 29) AND drug_id = 2;

SELECT pd.drug_id, pd.stock, c.quantity FROM pharmacy_drugs as pd JOIN carts as c
			ON pd.drug_id = c.drug_id
			WHERE pd.stock < c.quantity AND pd.pharmacy_id = 46 AND c.user_id = 2;

SELECT * FROM pharmacy_drugs;

SELECT pd1.pharmacy_id
                        FROM pharmacy_drugs pd1
                        JOIN pharmacy_drugs pd2 ON pd1.id <> pd2.id
                        WHERE pd2.pharmacy_id = 46 AND pd2.drug_id = 2 AND pd1.drug_id = 2 AND pd1.stock > pd2.stock
                LIMIT 1;


SELECT pd.drug_id FROM pharmacy_drugs as pd JOIN carts as c
ON pd.drug_id = c.drug_id
WHERE pd.stock < c.quantity;

SELECT * FROM journals;

SELECT @stock = pd.stock FROM pharmacy_drugs as pd WHERE pd.pharmacy_id = 46 AND pd.drug_id = 1;
SELECT @quantity = c.quantity FROM carts as c WHERE c.drug_id = 1;
WHILE @stock <= @quantity
BEGIN
    SELECT 
END;

ALTER TABLE order_details
DROP COLUMN drug_id;

ALTER TABLE order_details
ADD COLUMN pharmacy_drug_id BIGINT REFERENCES pharmacy_drugs(id);
SELECT * FROM orders;
SELECT * FROM order_details;
DELETE FROM order_details;
DELETE FROM orders;

SELECT * FROM stock_mutations;

SELECT * FROM status_mutations;
INSERT INTO status_mutations (name, created_at, updated_at)
VALUES
    ('Pending', NOW(), NOW()),
    ('Accepted', NOW(), NOW()),
    ('Rejected', NOW(), NOW()),
    ('Canceled', NOW(), NOW());

SELECT * FROM journals;
=======
        ) AND st_dwithin(ST_Makepoint(-41.01234, 106.816222)::geography, ST_MakePoint(a.longtitude, a.latitude)::geography, 25000);

INSERT INTO pharmacy_drugs (pharmacy_id, drug_id, stock, selling_unit, status, created_at, updated_at)
VALUES
    (28, 1, 100, 11000, 'active', NOW(), NOW()),
    (28, 2, 123, 12000, 'active', NOW(), NOW()),
    (28, 3, 101, 14000, 'active', NOW(), NOW()),
    (29, 1, 122, 12000, 'active', NOW(), NOW()),
    (29, 2, 144, 13000, 'active', NOW(), NOW()),
    (30, 3, 122, 14000, 'active', NOW(), NOW());

BEGIN TRANSACTION buy_products;

      SELECT * FROM pharmacies as p
        WHERE EXISTS (
            SELECT 1 FROM pharmacy_drugs
                     WHERE pharmacy_id = p.id AND drug_id IN (1, 2)
        ) LIMIT 1;
>>>>>>> dev
