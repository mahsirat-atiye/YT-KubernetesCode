package ingest

import (
	db "awesomeProject1/database"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestControllerCreateBook(t *testing.T) {
	//When testing controller you need to mock the service
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)
	testController := NewController(mockService)

	//Test output for mocking
	arg := CreateBookParams{
		Name:   "a",
		Author: "b",
	}

	//Mocking if createBook in service
	mockService.EXPECT().CreateBook(arg).Return(db.Book{
		Name:   "a",
		Author: "b",
	}, nil).Times(1)

	// Here we test the controller by itself
	//Test input for testing controller. this will be body of request.
	bookInputData, err := json.Marshal(gin.H{
		"name":   "a",
		"author": "b",
	})
	require.NoError(t, err)
	t.Run("Create book in controller", func(t *testing.T) {

		//input of the controller is a context. Let's create a context at first!
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		url := "/book/"
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bookInputData))
		require.NoError(t, err)
		ctx.Request = request

		testController.CreateBook(ctx)
		assert.Equal(t, http.StatusOK, ctx.Writer.Status())

	})
}
