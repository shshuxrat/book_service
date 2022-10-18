package service

import (
	"book_service/genproto/book_service"
	"book_service/pkg/logger"
	"book_service/storage"
	"context"

	"github.com/jmoiron/sqlx"
)

type bookCategoryService struct {
	logger  logger.LoggerI
	storage storage.StorageI
	book_service.UnimplementedBookCategoryServiceServer
}

func NewBookCategoryService(logger logger.LoggerI, db *sqlx.DB) *bookCategoryService {
	return &bookCategoryService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}

func (s *bookCategoryService) Create(c context.Context, req *book_service.CreateBookCategory) (*book_service.BookCategoryId, error) {
	id, err := s.storage.BookCategory().Create(req)
	if err != nil {
		return nil, err
	}

	return &book_service.BookCategoryId{
		Id: id,
	}, nil
}
func (s *bookCategoryService) GetAll(c context.Context, req *book_service.GetAllBookCategoryRequest) (*book_service.GetAllBookCategoryResponse, error) {
	bookCategoryList, err := s.storage.BookCategory().GetAll(req)
	if err != nil {
		return nil, err
	}

	return bookCategoryList, nil
}
func (s *bookCategoryService) GetById(c context.Context, req *book_service.BookCategoryId) (*book_service.BookCategory, error) {
	bookCategory, err := s.storage.BookCategory().GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return bookCategory, err
}
func (s *bookCategoryService) Update(c context.Context, req *book_service.BookCategory) (*book_service.MsgResponse, error) {
	msg, err := s.storage.BookCategory().Update(req)
	if err != nil {
		return nil, err
	}
	return &book_service.MsgResponse{
		Msg: msg.Msg,
	}, nil
}
func (s *bookCategoryService) Delete(c context.Context, req *book_service.BookCategoryId) (*book_service.MsgResponse, error) {
	msg, err := s.storage.BookCategory().Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &book_service.MsgResponse{
		Msg: msg.Msg,
	}, nil
}
