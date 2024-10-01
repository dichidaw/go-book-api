package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	s "go-book-api/services"
)

type RequestHandler struct {
	bookService      s.BookService
	userService      s.UserService
	borrowingService s.BorrowingService
}

func NewHandler(bookService s.BookService, userService s.UserService, borrowingService s.BorrowingService) *RequestHandler {
	return &RequestHandler{
		bookService:      bookService,
		userService:      userService,
		borrowingService: borrowingService,
	}
}

func (h *RequestHandler) ReadBooks(c *fiber.Ctx) error {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		log.Error("Error fetching books: ", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(books)
}

func (h *RequestHandler) ReadUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		log.Error("Error fetching users: ", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(users)
}

func (h *RequestHandler) ReadBorrowingHistories(c *fiber.Ctx) error {
	borrowingHistories, err := h.borrowingService.GetBorrowingHistories()
	if err != nil {
		log.Error("Error fetching users: ", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(borrowingHistories)
}
