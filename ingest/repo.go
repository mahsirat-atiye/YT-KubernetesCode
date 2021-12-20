package ingest

import (
	db "awesomeProject1/database"
	"github.com/jmoiron/sqlx"
	"log"
)

//go:generate mockgen -destination=mock_repo.go -package ingest . Repo
type Repo interface {
	CreateBook(name, author string) (db.Book, error)
	CreateBooksTable()
}

type repo struct {
	db *sqlx.DB
}

func (r repo) CreateBook(name, author string) (db.Book, error) {
	bookInsertion := `INSERT INTO books (name, author) VALUES ($1, $2) RETURNING *;`
	row := r.db.QueryRowx(bookInsertion, name, author)
	var b db.Book
	err := row.StructScan(&b)
	return b, err

}

func NewRepo() Repo {
	// layer1
	storage, err := db.NewDB()
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	r := &repo{
		db: storage,
	}
	r.DropBooksTable()
	r.CreateBooksTable()
	return r
}

func (r repo) CreateBooksTable() {
	schema := `CREATE TABLE books (
    id bigserial PRIMARY KEY ,
    name varchar NOT NULL ,
    author varchar NOT NULL);`
	r.db.Exec(schema)
}

func (r repo) DropBooksTable() {
	schema := `DROP TABLE IF EXISTS books;`
	r.db.Exec(schema)

}
