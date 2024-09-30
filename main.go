package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-book-api/routes"
)

func main() {
	fmt.Println("Hello There")

	app := fiber.New()

	routes.Routes(app)

	err := app.Listen(":3434")
	if err != nil {
		return
	}
}
