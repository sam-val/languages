package db

import "context"

const listMoviesByFilmmakerID = ` 
	SELECT DISTINCT ON (m.id) m.* from "Movie" as m
	INNER JOIN "MovieCredits" as mc
	ON m.id = mc.movie_id 
	INNER JOIN "Filmmaker" as fm
	ON fm.id = mc.filmmaker_id
	WHERE fm.id = $1
	ORDER BY m.id
	LIMIT $2
	OFFSET $3;
`
type ListMoviesByFilmmakerIDParams struct {
	FilmmakerID int64 `json:"filmmaker_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMoviesByFilmmakerID(ctx context.Context, arg ListMoviesByFilmmakerIDParams) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, listMoviesByFilmmakerID, arg.FilmmakerID, arg.Limit, arg.Offset)
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