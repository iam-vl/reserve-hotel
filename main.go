package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":1111", "Listening port on the API server")
	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	app.Get("/foo", HandleFoo)
	apiV1.Get("/user", HandleUser)
	app.Listen(*listenAddr)
}

func HandleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Working just fine!"})
}

func HandleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "James Foo"})
}
