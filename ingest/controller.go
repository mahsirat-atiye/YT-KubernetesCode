package ingest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	CreateBook(ctx *gin.Context)
}
type controller struct {
	svc Service
}

func NewController(service Service) Controller {
	return &controller{
		svc: service,
	}
}

func (c controller) CreateBook(ctx *gin.Context) {
	var req createBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}

	arg := CreateBookParams{
		Name:   req.Name,
		Author: req.Author,
	}

	book, err := c.svc.CreateBook(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, book)
	return
}
