package datamodels

import "time"

type User struct {
	ID      uint        `json:"id" gorm:"primaryKey"`
	Name    string      `json:"name"`
	Borrows []Borrowing `json:"borrows" gorm:"foreignKey:UserID"`
}

type Book struct {
	ID         uint        `gorm:"primaryKey"`
	Title      string      `json:"title"`
	IsBorrowed bool        `json:"isBorrowed"`
	BorrowedBy []Borrowing `json:"borrowedBy" gorm:"foreignKey:BookID"`
	Writers    []Writer    `json:"writers" gorm:"many2many:book_writers;foreignKey:ID;joinForeignKey:BookID;References:ID;joinReferences:WriterID"`
}

type Borrowing struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	BookID     uint       `json:"bookID"`
	UserID     uint       `json:"userID"`
	BorrowedAt time.Time  `json:"borrowedAt"`
	ReturnedAt *time.Time `json:"returnedAt"`
}

type Writer struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Books []Book `json:"books" gorm:"many2many:book_writers;foreignKey:ID;joinForeignKey:WriterID;References:ID;joinReferences:BookID"`
}

type BookWriter struct {
	BookID   uint `gorm:"primaryKey"`
	WriterID uint `gorm:"primaryKey"`
}
