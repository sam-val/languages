-- name: GetMovie :one
SELECT * FROM "Movie"
WHERE id = $1
LIMIT 1;

-- name: ListMovies :many
SELECT * FROM "Movie"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateMovie :one
INSERT INTO "Movie" (
  name, year
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateMovie :one
UPDATE "Movie"
SET name = $2, year = $3
WHERE id = $1
RETURNING *;

-- name: DeleteMovie :execrows
DELETE FROM "Movie"
WHERE id = $1;
