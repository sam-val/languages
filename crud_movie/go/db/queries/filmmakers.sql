-- name: GetFilmmaker :one
SELECT * FROM "Filmmaker"
WHERE id = $1
LIMIT 1;

-- name: ListFilmmakers :many
SELECT * FROM "Filmmaker"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateFilmmaker :one
INSERT INTO "Filmmaker" (
  name, dob
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateFilmmaker :one
UPDATE "Filmmaker"
SET name = $2, dob = $3
WHERE id = $1
RETURNING *;

-- name: DeleteFilmmaker :exec
DELETE FROM "Filmmaker"
WHERE id = $1;
