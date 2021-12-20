package ingest

import (
	db "awesomeProject1/database"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceCreateBook(t *testing.T) {
	//When testing service you need to mock the repo
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)
	testService := NewService(mockRepo)

	//Mocking if createBook in service
	mockRepo.EXPECT().CreateBook("a", "b").Return(db.Book{
		Name:   "a",
		Author: "b",
	}, nil).Times(1)

	// Here we test the service by itself
	//Test input for testing service.

	arg := CreateBookParams{
		Name:   "a",
		Author: "b",
	}

	t.Run("Create book in service", func(t *testing.T) {

		result, err := testService.CreateBook(arg)
		require.NoError(t, err)

		book := db.Book{
			Name:   "a",
			Author: "b",
		}
		assert.Equal(t, book.Name, result.Name)
		assert.Equal(t, book.Author, result.Author)

	})
}
