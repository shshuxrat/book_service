package repo

import "book_service/genproto/book_service"

type BookRepoI interface {
	Create(req *book_service.CreateBook) (string, error)
	GetAll(req *book_service.GetAllBookRequest) (*book_service.GetAllBookResponse, error)
	GetById(id string) (*book_service.GetBookByIdResponse, error)
	Update(req *book_service.UpdateBook) (*book_service.MsgRespons, error)
	Delete(id string) (*book_service.MsgRespons, error)
}
