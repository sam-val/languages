// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: movies.sql

package db

import (
	"context"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO "Movie" (
  name, year
) VALUES (
  $1, $2
)
RETURNING id, name, year, created_at
`

type CreateMovieParams struct {
	Name string `json:"name"`
	Year int32  `json:"year"`
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie, arg.Name, arg.Year)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Year,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :execrows
DELETE FROM "Movie"
WHERE id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteMovie, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getMovie = `-- name: GetMovie :one
SELECT id, name, year, created_at FROM "Movie"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetMovie(ctx context.Context, id int64) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Year,
		&i.CreatedAt,
	)
	return i, err
}

const listMovies = `-- name: ListMovies :many
SELECT id, name, year, created_at FROM "Movie"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListMoviesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMovies(ctx context.Context, arg ListMoviesParams) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, listMovies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Movie{}
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Year,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :one
UPDATE "Movie"
SET name = $2, year = $3
WHERE id = $1
RETURNING id, name, year, created_at
`

type UpdateMovieParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Year int32  `json:"year"`
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, updateMovie, arg.ID, arg.Name, arg.Year)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Year,
		&i.CreatedAt,
	)
	return i, err
}
