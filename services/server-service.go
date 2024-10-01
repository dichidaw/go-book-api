package services

import (
	dm "go-book-api/datamodels"
	r "go-book-api/repositories"
)

type BookService interface {
	GetAllBooks() ([]dm.Book, error)
}

type UserService interface {
	GetAllUsers() ([]dm.User, error)
}

type BorrowingService interface {
	GetBorrowingHistories() ([]dm.Borrowing, error)
}

type bookService struct {
	repo r.BookRepository
}

type userService struct {
	repo r.UserRepository
}

type borrowingService struct {
	repo r.BorrowingRepository
}

func NewBookService(repo r.BookRepository) BookService {
	return &bookService{repo}
}

func NewUserService(repo r.UserRepository) UserService {
	return &userService{repo}
}

func NewBorrowingService(repo r.BorrowingRepository) BorrowingService {
	return &borrowingService{repo}
}

func (s *bookService) GetAllBooks() ([]dm.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *userService) GetAllUsers() ([]dm.User, error) {
	return s.repo.GetAllUsers()
}

func (s *borrowingService) GetBorrowingHistories() ([]dm.Borrowing, error) {
	return s.repo.GetBorrowingHistories()
}
