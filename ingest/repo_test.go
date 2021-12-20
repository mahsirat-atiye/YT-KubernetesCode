package ingest

import (
	db "awesomeProject1/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepoCreateBook(t *testing.T) {

	repo := NewRepo()

	t.Run("Create book in service", func(t *testing.T) {

		result, err := repo.CreateBook("a", "b")
		require.NoError(t, err)
		book := db.Book{
			Name:   "a",
			Author: "b",
		}
		assert.Equal(t, book.Name, result.Name)
		assert.Equal(t, book.Author, result.Author)

	})
}
