package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	server.router = router
	return server
}

func (s *Server) Start(addr string) error {
	if err := s.router.Run(addr); err != nil {
		return err
	}
	return nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
