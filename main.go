package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"go-book-api/handlers"
	"go-book-api/routes"
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

	app := fiber.New()
	routes.Routes(app)

	appPort := os.Getenv("APP_PORT")
	err = app.Listen(fmt.Sprintf(":%s", appPort))
	if err != nil {
		log.Fatal("Failed to listen on port ", err)
		return
	}
}
