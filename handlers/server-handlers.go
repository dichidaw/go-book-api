package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ReadBooks(c *fiber.Ctx) error {
	err := c.SendString("'Hi', Fiber said")
	if err != nil {
		log.Fatal("Error in ReadBooks: ", err)
		return err
	}
	return nil
}

//
//func ReadBooksByID(c *fiber.Ctx) error {
//
//}
//
//func AddBooks(c *fiber.Ctx) error {
//
//}
//
//func UpdateBookByID(c *fiber.Ctx) error {
//
//}
//
//func DeleteBookByID(c *fiber.Ctx) error {
//
//}
