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
