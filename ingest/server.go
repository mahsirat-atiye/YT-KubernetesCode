package ingest

import "github.com/gin-gonic/gin"

type Server struct {
	ctrl   Controller
	router *gin.Engine
}

func NewServer(ctrl Controller) *Server {
	router := gin.Default()
	server := &Server{ctrl: ctrl}
	router.POST("/book", server.ctrl.CreateBook)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	server.router.Run(address)
	return nil
}
