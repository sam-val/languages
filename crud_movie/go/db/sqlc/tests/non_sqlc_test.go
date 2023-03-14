package tests

import (
	"context"
	"testing"

	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestListMoviesByFilmmakerID(t *testing.T) {
	movie := CreateRandomMovie(t)
	filmmaker := CreateRandomFilmmaker(t)
	n := 3 
	for i:=0;i<n;i++ {
		arg := db.CreateMovieCreditParams{
			MovieID: movie.ID,
			FilmmakerID: filmmaker.ID,
			RoleID: CreateRandomRole(t).ID,
		}
		_, err := testQueries.CreateMovieCredit(context.Background(), arg)
		require.NoError(t, err)
	}
	
	arg := db.ListMoviesByFilmmakerIDParams{
		FilmmakerID: filmmaker.ID,
		Limit: 5,
		Offset: 0,
	}
	movies, err := testQueries.ListMoviesByFilmmakerID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movies)

	require.Equal(t, 1, len(movies))
	require.Equal(t, movie.ID, movies[0].ID)
	require.Equal(t, movie.Name, movies[0].Name)
}