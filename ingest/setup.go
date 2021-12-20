package ingest

import (
	"github.com/gin-gonic/gin"
)

const address = "0.0.0.0:8080"

func SetupService() {
	repo := NewRepo()
	service := NewService(repo)
	controller := NewController(service)

	router := gin.Default()
	router.POST("/book", controller.CreateBook)
	router.Run(address)

}
