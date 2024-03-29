// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: filmmakers.sql

package db

import (
	"context"
	"time"
)

const createFilmmaker = `-- name: CreateFilmmaker :one
INSERT INTO "Filmmaker" (
  name, dob
) VALUES (
  $1, $2
)
RETURNING id, name, dob, created_at
`

type CreateFilmmakerParams struct {
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
}

func (q *Queries) CreateFilmmaker(ctx context.Context, arg CreateFilmmakerParams) (Filmmaker, error) {
	row := q.db.QueryRowContext(ctx, createFilmmaker, arg.Name, arg.Dob)
	var i Filmmaker
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Dob,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFilmmaker = `-- name: DeleteFilmmaker :execrows
DELETE FROM "Filmmaker"
WHERE id = $1
`

func (q *Queries) DeleteFilmmaker(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteFilmmaker, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getFilmmaker = `-- name: GetFilmmaker :one
SELECT id, name, dob, created_at FROM "Filmmaker"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetFilmmaker(ctx context.Context, id int64) (Filmmaker, error) {
	row := q.db.QueryRowContext(ctx, getFilmmaker, id)
	var i Filmmaker
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Dob,
		&i.CreatedAt,
	)
	return i, err
}

const listFilmmakers = `-- name: ListFilmmakers :many
SELECT id, name, dob, created_at FROM "Filmmaker"
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListFilmmakersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFilmmakers(ctx context.Context, arg ListFilmmakersParams) ([]Filmmaker, error) {
	rows, err := q.db.QueryContext(ctx, listFilmmakers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Filmmaker{}
	for rows.Next() {
		var i Filmmaker
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Dob,
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

const updateFilmmaker = `-- name: UpdateFilmmaker :one
UPDATE "Filmmaker"
SET name = $2, dob = $3
WHERE id = $1
RETURNING id, name, dob, created_at
`

type UpdateFilmmakerParams struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
}

func (q *Queries) UpdateFilmmaker(ctx context.Context, arg UpdateFilmmakerParams) (Filmmaker, error) {
	row := q.db.QueryRowContext(ctx, updateFilmmaker, arg.ID, arg.Name, arg.Dob)
	var i Filmmaker
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Dob,
		&i.CreatedAt,
	)
	return i, err
}
