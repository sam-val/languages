package tests

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/sam-val/languages/crud_movie/go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRole(t *testing.T) db.Role {
	name := util.RandomString(10)

	role, err := testQueries.CreateRole(context.Background(), name)
	require.NoError(t, err)

	require.Equal(t, role.Name, name)

	require.NotZero(t, role.CreatedAt)
	require.NotZero(t, role.ID)

	return role
}

func TestCreateRole(t *testing.T) {
	CreateRandomRole(t)
}

func TestGetRole(t *testing.T) {
	role := CreateRandomRole(t)
	
	result_role, err := testQueries.GetRole(context.Background(), role.ID)
	require.NoError(t, err)

	require.NotEmpty(t, result_role)
	require.Equal(t, role.Name, result_role.Name)
	require.Equal(t, role.CreatedAt, result_role.CreatedAt)
}

func TestDeleteRole(t *testing.T) {
	role := CreateRandomRole(t)
	rows, err := testQueries.DeleteRole(context.Background(), role.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)

	role2, err := testQueries.GetRole(context.Background(), role.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, role2)
}

func TestUpdateRole(t *testing.T) {
	role := CreateRandomRole(t)

	arg := db.UpdateRoleParams{
		ID: role.ID,
		Name: util.RandomOwner(),
	}

	role2, err := testQueries.UpdateRole(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, role2)

	require.Equal(t, role2.ID, arg.ID)
	require.Equal(t, role2.Name, arg.Name)
}

func TestListRoles(t *testing.T) {
	n := 8
	for i := 0; i <=n; i++ {
		CreateRandomRole(t)
	}
	arg := db.ListRolesParams{
		Limit: 3,
		Offset: 5,
	}

	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roles)
	require.Equal(t, int32(len(roles)), arg.Limit)

	for _, m := range roles {
		require.NotEmpty(t, m)
	}
}
