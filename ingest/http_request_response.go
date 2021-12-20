package ingest

import "github.com/gin-gonic/gin"

type CreateBookParams struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type createBookRequest struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// Handlers
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
