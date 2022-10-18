package repo

import "book_service/genproto/book_service"

type BookCategoryRepoI interface {
	Create(req *book_service.CreateBookCategory) (string, error)
	GetAll(req *book_service.GetAllBookCategoryRequest) (*book_service.GetAllBookCategoryResponse, error)
	GetById(id string) (*book_service.BookCategory, error)
	Update(req *book_service.BookCategory) (*book_service.MsgResponse, error)
	Delete(id string) (*book_service.MsgResponse, error)
}
