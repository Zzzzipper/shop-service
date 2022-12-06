-- name: AddPartner :one
INSERT INTO partners (
  full_name,
  url,
  role
) VALUES (
  $1, 
  $2,
  $3
)
RETURNING *;

-- name: DeletePartner :one
DELETE FROM partners
WHERE id = $1
RETURNING *;

-- name: AddMerchant :one
INSERT INTO merchants (
  full_name,
  url,
  partner_id
) VALUES (
  $1, 
  $2,
  $3
)
RETURNING *;

-- name: DeleteMerchant :one
DELETE FROM merchants
WHERE id = $1
RETURNING *;


-- name: AddShop :one
INSERT INTO shops (
  full_name,
  url,
  merchant_id,
  login,
  password
) VALUES (
  $1, 
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: DeleteShop :one
DELETE FROM shops
WHERE id = $1
RETURNING *;

-- name: AddTerminal :one
INSERT INTO terminals (
  full_name,
  url,
  shop_id,
  login,
  password,
  token
) VALUES (
  $1, 
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: DeleteTerminal :one
DELETE FROM terminals
WHERE id = $1
RETURNING *;

-- name: Auth :one
SELECT
  p.id as partner_id, p.full_name as partner_full_name, p.url as partner_url, p.role as partner_role,
	m.id as merchant_id, m.full_name as merchant_full_name, m.url as merchant_url, 
	s.id as shop_id, s.full_name as shop_full_name, s.url as shop_url,
	t.id as terminal_id, t.full_name as terminal_full_name, t.url as terminal_url, t.login as terminal_login
FROM 
	terminals t, shops s, merchants m, partners p
WHERE
	t.shop_id=s.id 
AND
	s.merchant_id=m.id
AND 
	m.partner_id=p.id
AND
	t.login=$1
AND
	t.password=$2
LIMIT 1;
