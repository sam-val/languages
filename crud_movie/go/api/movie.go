package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
)

type createMovieRequest struct {
	Name    string `json:"name" binding:"required"`
	Year int `json:"year" binding:"required,min=1890"`
}

func (server *Server) createMovie(ctx *gin.Context){
	var payload createMovieRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMovieParams{
		Name: payload.Name,
		Year: int32(payload.Year),
	}
	movie, err := server.store.CreateMovie(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(movie))
}

type getMovieRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getMovie(ctx *gin.Context) {
	var uri getMovieRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	movie, err := server.store.GetMovie(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(movie))
}

type deleteMovieRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteMovie(ctx *gin.Context) {
	var req deleteMovieRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	rows, err := server.store.DeleteMovie(ctx, req.ID)
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

type listMoviesRequest struct {
	FilmmakerID int64 `form:"filmmaker_id"`
	PageID int32 `form:"page_id" binding:"min=1"`
	PageSize int32 `form:"page_size" binding:"min=5,max=10"`
}

func (server *Server) listMovies(ctx *gin.Context) {
	req := listMoviesRequest{
		FilmmakerID: 0,
		PageID: 1,
		PageSize: 5,
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var movies []db.Movie
	var query_func func() ([]db.Movie, error)
	if req.FilmmakerID != 0 {
		arg := db.ListMoviesByFilmmakerIDParams{
			FilmmakerID: req.FilmmakerID,
			Limit: req.PageSize,
			Offset: (req.PageID-1)*req.PageSize,
		}
		query_func = func() ([]db.Movie, error) {
			return server.store.ListMoviesByFilmmakerID(ctx, arg)
		}
	} else {
		arg := db.ListMoviesParams{
			Limit: req.PageSize,
			Offset: (req.PageID-1)*req.PageSize, 
		}
		query_func = func() ([]db.Movie, error) {
			return server.store.ListMovies(ctx, arg)
		}
	}
	movies, err := query_func()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(movies))
}

type listMoviesByFimmakerIDRequest struct {
	FilmmakerID int64 `form:"filmmaker_id" binding:"min=1"`
	PageID int32 `form:"page_id" binding:"min=1"`
	PageSize int32 `form:"page_size" binding:"min=5,max=10"`
}

func (server *Server) listMoviesByFilmmakerID(ctx *gin.Context) {
	req := listMoviesByFimmakerIDRequest{
		PageID: 1,
		PageSize: 5,
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMoviesByFilmmakerIDParams{
		FilmmakerID: req.FilmmakerID,
		Limit: req.PageSize,
		Offset: (req.PageID-1)*req.PageSize, 
	}
	movies, err := server.store.ListMoviesByFilmmakerID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sucessResponse(movies))
}
