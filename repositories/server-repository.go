package repositories

import (
	dm "go-book-api/datamodels"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]dm.Book, error)
}

type UserRepository interface {
	GetAllUsers() ([]dm.User, error)
}

type BorrowingRepository interface {
	GetBorrowingHistories() ([]dm.Borrowing, error)
}

type bookRepository struct {
	db *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

type borrowingRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func NewBorrowingRepository(db *gorm.DB) BorrowingRepository {
	return &borrowingRepository{db}
}

func (r *bookRepository) GetAllBooks() ([]dm.Book, error) {
	var books []dm.Book
	err := r.db.Preload("Writers").Preload("BorrowedBy").Find(&books).Error
	return books, err
}

func (r *userRepository) GetAllUsers() ([]dm.User, error) {
	var users []dm.User
	err := r.db.Preload("Borrows").Find(&users).Error
	return users, err
}

func (r *borrowingRepository) GetBorrowingHistories() ([]dm.Borrowing, error) {
	var borrowingHistories []dm.Borrowing
	err := r.db.Find(&borrowingHistories).Error
	return borrowingHistories, err
}
