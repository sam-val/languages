package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}	
	router := gin.Default()

	router.POST("movie/", server.createMovie)
	router.GET("movie/:id", server.getMovie)
	router.GET("movies/", server.listMovies)
	router.DELETE("movie/:id", server.deleteMovie)

	router.POST("filmmaker/", server.createFilmmaker)
	router.GET("filmmaker/:id", server.getFilmmaker)
	router.GET("filmmakers/", server.listFilmmakers)
	router.DELETE("filmmaker/:id", server.deleteFilmmaker)

	router.POST("role/", server.createRole)
	router.GET("role/:id", server.getRole)
	router.GET("roles/", server.listRoles)
	router.DELETE("role/:id", server.deleteRole)

	router.POST("credit", server.createCredit)

	server.router = router
	return server
}

func (server *Server) Start(path string) error {
	return server.router.Run(path)
}

func sucessResponse(object any) gin.H {
	return gin.H{"sucesss": true, "data": object}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
