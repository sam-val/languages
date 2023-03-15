package tests

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/stretchr/testify/require"
)

func CreateRandomCredit(t *testing.T) db.MovieCredit {
	movie := CreateRandomMovie(t)
	filmmaker := CreateRandomFilmmaker(t)
	role := CreateRandomRole(t)

	arg := db.CreateMovieCreditParams{
		MovieID: movie.ID,
		FilmmakerID: filmmaker.ID,
		RoleID: role.ID,
	}
	credit, err := testQueries.CreateMovieCredit(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, credit.MovieID, arg.MovieID)
	require.Equal(t, credit.FilmmakerID, arg.FilmmakerID)
	require.Equal(t, credit.RoleID, arg.RoleID)

	require.NotZero(t, credit.CreatedAt)
	require.NotZero(t, credit.ID)

	return credit
}

func TestCreateCredit(t *testing.T) {
	CreateRandomCredit(t)
}

func TestGetCredit(t *testing.T) {
	credit := CreateRandomCredit(t)
	
	credit2, err := testQueries.GetMovieCredit(context.Background(), credit.ID)
	require.NoError(t, err)

	require.NotEmpty(t, credit2)
	require.Equal(t, credit2.FilmmakerID, credit.FilmmakerID)
	require.Equal(t, credit2.MovieID, credit.MovieID)
	require.Equal(t, credit2.RoleID, credit.RoleID)
	require.Equal(t, credit2.CreatedAt, credit.CreatedAt)
}

func TestDeleteCredit(t *testing.T) {
	credit := CreateRandomCredit(t)
	rows, err := testQueries.DeleteMovieCredit(context.Background(), credit.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)

	credit2, err := testQueries.GetMovieCredit(context.Background(), credit.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, credit2)
}

func TestUpdateCredit(t *testing.T) {
	credit := CreateRandomCredit(t)

	filmmaker := CreateRandomFilmmaker(t)
	movie := CreateRandomMovie(t)
	role := CreateRandomRole(t)
	arg := db.UpdateMovieCreditParams{
		ID: credit.ID,
		FilmmakerID: filmmaker.ID,
		MovieID: movie.ID,
		RoleID: role.ID,

	}

	credit2, err := testQueries.UpdateMovieCredit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, credit2)

	require.Equal(t, credit2.ID, arg.ID)
	require.Equal(t, credit2.MovieID, movie.ID)
	require.Equal(t, credit2.FilmmakerID, filmmaker.ID)
	require.Equal(t, credit2.RoleID, role.ID)
}
