package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"go-book-api/handlers"
	"go-book-api/repositories"
	"go-book-api/routes"
	"go-book-api/services"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file: ", err)
		return
	}

	dbConnection := handlers.ConnectDB()
	defer handlers.CloseDBConn(dbConnection)

	// init handler
	bookRepo := repositories.NewBookRepository(dbConnection)
	bookService := services.NewBookService(bookRepo)

	userRepo := repositories.NewUserRepository(dbConnection)
	userService := services.NewUserService(userRepo)

	borrowingRepo := repositories.NewBorrowingRepository(dbConnection)
	borrowingService := services.NewBorrowingService(borrowingRepo)

	bookHandler := handlers.NewHandler(bookService, userService, borrowingService)

	app := fiber.New()
	routes.Routes(app, bookHandler)

	appPort := os.Getenv("APP_PORT")
	err = app.Listen(fmt.Sprintf(":%s", appPort))
	if err != nil {
		log.Fatal("Failed to listen on port ", err)
		return
	}
}
