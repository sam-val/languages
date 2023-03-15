-- name: GetRole :one
SELECT * FROM "Role"
WHERE id = $1
LIMIT 1;

-- name: ListRoles :many
SELECT * FROM "Role"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateRole :one
INSERT INTO "Role" (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateRole :one
UPDATE "Role"
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRole :execrows
DELETE FROM "Role"
WHERE id = $1;
