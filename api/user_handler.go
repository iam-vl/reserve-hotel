package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/reserve-hotel/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Vassily",
		LastName:  "Pupkin",
	}
	return c.JSON(u)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("User")
}
