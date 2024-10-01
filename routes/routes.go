package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-book-api/handlers"
)

func Routes(app *fiber.App, handler *handlers.RequestHandler) {
	books := app.Group("/books")
	books.Get("/", handler.ReadBooks)

	users := app.Group("/users")
	users.Get("/", handler.ReadUsers)

	borrowingHistories := app.Group("/borrowingHistories")
	borrowingHistories.Get("/", handler.ReadBorrowingHistories)
}
