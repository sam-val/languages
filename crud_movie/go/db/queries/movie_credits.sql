-- name: GetMovieCredit :one
SELECT * FROM "MovieCredits"
WHERE id = $1
LIMIT 1;

-- name: CreateMovieCredit :one
INSERT INTO "MovieCredits" (
  movie_id, filmmaker_id, role_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateMovieCredit :one
UPDATE "MovieCredits"
SET movie_id = $2, filmmaker_id = $3, role_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteMovieCredit :exec
DELETE FROM "MovieCredits"
WHERE id = $1;
