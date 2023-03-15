package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
)

type createCreditRequest struct {
	FilmmakerID    int64 `json:"filmmaker_id" binding:"required,min=1"`
	MovieID    int64 `json:"movie_id" binding:"required,min=1"`
	RoleID    int64 `json:"role_id" binding:"required,min=1"`
}

func (server *Server) createCredit(ctx *gin.Context){
	var payload createCreditRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMovieCreditParams{
		FilmmakerID: payload.FilmmakerID,
		MovieID: payload.MovieID,
		RoleID: payload.RoleID,
	}
	credit, err := server.store.CreateMovieCredit(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(credit))
}
