package api

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
)

type createRoleRequest struct {
	Name    string `json:"name" binding:"required"`
}

func (server *Server) createRole(ctx *gin.Context){
	var payload createRoleRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	role, err := server.store.CreateRole(ctx, payload.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(role))
}

type getRoleRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRole(ctx *gin.Context) {
	var uri getRoleRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	role, err := server.store.GetRole(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(role))
}

type deleteRoleRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteRole(ctx *gin.Context) {
	var req deleteRoleRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	rows, err := server.store.DeleteRole(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	if rows != 1 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(nil))
}

type listRolesRequest struct {
	PageID int32 `form:"page_id" binding:"min=1"`
	PageSize int32 `form:"page_size" binding:"min=5,max=10"`
}

func (server *Server) listRoles(ctx *gin.Context) {
	req := listRolesRequest{
		PageID: 1,
		PageSize: 5,
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRolesParams{
		Limit: req.PageSize,
		Offset: (req.PageID-1)*req.PageSize, 
	}

	roles, err := server.store.ListRoles(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(roles))
}
