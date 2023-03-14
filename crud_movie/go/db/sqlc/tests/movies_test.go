package tests

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/sam-val/languages/crud_movie/go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomMovie(t *testing.T) db.Movie {
	name := util.RandomOwner()
	year := util.RandomYear()

	arg := db.CreateMovieParams{
		Name: name,
		Year: year,
	}
	movie, err := testQueries.CreateMovie(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, movie.Name, arg.Name)
	require.Equal(t, movie.Year, arg.Year)

	require.NotZero(t, movie.CreatedAt)
	require.NotZero(t, movie.ID)

	return movie
}

func TestCreateMovie(t *testing.T) {
	CreateRandomMovie(t)
}

func TestGetMovie(t *testing.T) {
	movie := CreateRandomMovie(t)
	
	result_movie, err := testQueries.GetMovie(context.Background(), movie.ID)
	require.NoError(t, err)

	require.NotEmpty(t, result_movie)
	require.Equal(t, movie.Name, result_movie.Name)
	require.Equal(t, movie.Year, result_movie.Year)
	require.Equal(t, movie.CreatedAt, result_movie.CreatedAt)
}

func TestDeleteMovie(t *testing.T) {
	movie := CreateRandomMovie(t)
	err := testQueries.DeleteMovie(context.Background(), movie.ID)
	require.NoError(t, err)

	movie2, err := testQueries.GetMovie(context.Background(), movie.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, movie2)
}

func TestUpdateMovie(t *testing.T) {
	movie := CreateRandomMovie(t)

	arg := db.UpdateMovieParams{
		ID: movie.ID,
		Name: util.RandomOwner(),
		Year: util.RandomYear(),
	}

	movie2, err := testQueries.UpdateMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movie2)

	require.Equal(t, movie2.ID, arg.ID)
	require.Equal(t, movie2.Name, arg.Name)
	require.Equal(t, movie2.Year, arg.Year)
	require.Equal(t, movie2.CreatedAt, movie.CreatedAt)
}

func TestListMovies(t *testing.T) {
	n := 8
	for i := 0; i <=n; i++ {
		CreateRandomMovie(t)
	}
	arg := db.ListMoviesParams{
		Limit: 3,
		Offset: 5,
	}

	movies, err := testQueries.ListMovies(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movies)
	require.Equal(t, int32(len(movies)), arg.Limit)

	for _, m := range movies {
		require.NotEmpty(t, m)
	}
}
