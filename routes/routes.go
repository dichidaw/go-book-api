package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-book-api/handlers"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.ReadBooks)
	app.Get("/:id", handlers.ReadBooksByID)
	app.Post("/", handlers.AddBooks)
	app.Patch("/:id", handlers.UpdateBookByID)
	app.Delete("/:id", handlers.DeleteBookByID)
}
