package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("User")
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("User")
}
