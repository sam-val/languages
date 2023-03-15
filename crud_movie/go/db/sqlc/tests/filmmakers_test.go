package tests

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/sam-val/languages/crud_movie/go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomFilmmaker(t *testing.T) db.Filmmaker {
	name := util.RandomOwner()
	dob := util.RandomDate()

	arg := db.CreateFilmmakerParams{
		Name: name,
		Dob: dob,
	}
	filmmaker, err := testQueries.CreateFilmmaker(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, filmmaker.Name, arg.Name)
	require.Equal(t, filmmaker.Dob, arg.Dob)

	require.NotZero(t, filmmaker.CreatedAt)
	require.NotZero(t, filmmaker.ID)

	return filmmaker
}

func TestCreateFilmmaker(t *testing.T) {
	CreateRandomFilmmaker(t)
}

func TestGetFilmmaker(t *testing.T) {
	filmmaker := CreateRandomFilmmaker(t)
	
	filmmaker2, err := testQueries.GetFilmmaker(context.Background(), filmmaker.ID)
	require.NoError(t, err)

	require.NotEmpty(t, filmmaker2)
	require.Equal(t, filmmaker.Name, filmmaker2.Name)
	require.Equal(t, filmmaker.Dob, filmmaker2.Dob)
	require.Equal(t, filmmaker.CreatedAt, filmmaker2.CreatedAt)
}

func TestDeleteFilmmaker(t *testing.T) {
	filmmaker := CreateRandomFilmmaker(t)
	rows, err := testQueries.DeleteFilmmaker(context.Background(), filmmaker.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)

	filmmaker2, err := testQueries.GetFilmmaker(context.Background(), filmmaker.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, filmmaker2)
}

func TestUpdateFilmmaker(t *testing.T) {
	filmmaker := CreateRandomFilmmaker(t)

	arg := db.UpdateFilmmakerParams{
		ID: filmmaker.ID,
		Name: util.RandomOwner(),
		Dob: util.RandomDate(),
	}

	filmmaker2, err := testQueries.UpdateFilmmaker(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, filmmaker2)

	require.Equal(t, filmmaker2.ID, arg.ID)
	require.Equal(t, filmmaker2.Name, arg.Name)
	require.Equal(t, filmmaker2.Dob, arg.Dob)
	require.Equal(t, filmmaker2.CreatedAt, filmmaker2.CreatedAt)
}

func TestListFilmmaker(t *testing.T) {
	n := 8
	for i := 0; i <=n; i++ {
		CreateRandomFilmmaker(t)
	}
	arg := db.ListFilmmakersParams{
		Limit: 3,
		Offset: 5,
	}

	filmmakers, err := testQueries.ListFilmmakers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, filmmakers)
	require.Equal(t, int32(len(filmmakers)), arg.Limit)

	for _, m := range filmmakers {
		require.NotEmpty(t, m)
	}
}
