package storage

import (
	"book_service/storage/postgres"
	"book_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	BookCategory() repo.BookCategoryRepoI
	Book() repo.BookRepoI
}

type storagePG struct {
	bookCategory repo.BookCategoryRepoI
	book         repo.BookRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		bookCategory: postgres.NewBookCategoryRepo(db),
		book:         postgres.NewBookRepo(db),
	}

}
func (s *storagePG) BookCategory() repo.BookCategoryRepoI {
	return s.bookCategory
}
func (s *storagePG) Book() repo.BookRepoI {
	return s.book
}
