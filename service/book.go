package service

import (
	"book_service/genproto/book_service"
	"book_service/pkg/logger"
	"book_service/storage"
	"context"

	"github.com/jmoiron/sqlx"
)

type bookService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	book_service.UnimplementedBookServiceServer
}

func NewBookService(logger logger.LoggerI, db *sqlx.DB) *bookService {
	return &bookService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}
func (s *bookService) Create(c context.Context, req *book_service.CreateBook) (*book_service.BookId, error) {
	bookId, err := s.storage.Book().Create(req)
	if err != nil {
		return nil, err
	}
	return &book_service.BookId{
		Id: bookId,
	}, nil
}

func (s *bookService) GetAll(c context.Context, req *book_service.GetAllBookRequest) (*book_service.GetAllBookResponse, error) {
	bookList, err := s.storage.Book().GetAll(req)
	if err != nil {
		return nil, err
	}

	return bookList, nil
}

func (s *bookService) GetById(c context.Context, req *book_service.BookId) (*book_service.GetBookByIdResponse, error) {
	bookdata, err := s.storage.Book().GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return bookdata, nil
}

func (s *bookService) Update(c context.Context, req *book_service.UpdateBook) (*book_service.MsgRespons, error) {
	msg, err := s.storage.Book().Update(req)
	if err != nil {
		return nil, err
	}

	return &book_service.MsgRespons{
		Msg: msg.Msg,
	}, nil
}
func (s *bookService) Delete(c context.Context, req *book_service.BookId) (*book_service.MsgRespons, error) {
	msg, err := s.storage.Book().Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &book_service.MsgRespons{
		Msg: msg.Msg,
	}, nil
}
