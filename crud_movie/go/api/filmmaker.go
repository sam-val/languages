package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
)

type createFilmmakerRequest struct {
	Name    string `json:"name" binding:"required"`
	// example: 2009-11-10T00:00:00Z
	Dob 	string `json:"dob" binding:"required"`
}

func (server *Server) createFilmmaker(ctx *gin.Context){
	var payload createFilmmakerRequest
	var dob time.Time
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dob, err := time.Parse(time.RFC3339,payload.Dob)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("bad DOB format")))
		return
	}

	arg := db.CreateFilmmakerParams{
		Name: payload.Name,
		Dob: dob,
	}
	filmmaker, err := server.store.CreateFilmmaker(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(filmmaker))
}

type getFilmmakerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFilmmaker(ctx *gin.Context) {
	var uri getFilmmakerRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	filmmaker, err := server.store.GetFilmmaker(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(filmmaker))
}

type deleteFilmmakerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteFilmmaker(ctx *gin.Context) {
	var req deleteFilmmakerRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	rows, err := server.store.DeleteFilmmaker(ctx, req.ID)
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

type listFilmmakersRequest struct {
	PageID int32 `form:"page_id" binding:"min=1"`
	PageSize int32 `form:"page_size" binding:"min=5,max=10"`
}

func (server *Server) listFilmmakers(ctx *gin.Context) {
	req := listFilmmakersRequest{
		PageID: 1,
		PageSize: 5,
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMoviesParams{
		Limit: req.PageSize,
		Offset: (req.PageID-1)*req.PageSize, 
	}

	filmmakers, err := server.store.ListMovies(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(filmmakers))
}
