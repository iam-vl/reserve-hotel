package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/reserve-hotel/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":1111", "Listening port on the API server")
	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HandleGetUsers)        // Get all users
	apiV1.Get("/user/:id", api.HandleGetUserById) // Get user by Id
	app.Listen(*listenAddr)
}
