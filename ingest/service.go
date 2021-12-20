package ingest

import (
	db "awesomeProject1/database"
	_ "database/sql"
)

type Service interface {
	CreateBook(arg CreateBookParams) (db.Book, error)
}
type service struct {
	repo Repo
}

func (s *service) CreateBook(arg CreateBookParams) (db.Book, error) {

	b, err := s.repo.CreateBook(arg.Name, arg.Author)
	return b, err
}

func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}
