package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iam-vl/reserve-hotel/api"
	"github.com/iam-vl/reserve-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "reserve-hotel"
	userColl = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	user := types.User{
		FirstName: "Vassily",
		LastName:  "La",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var vl types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&vl); err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":1111", "Listening port on the API server")
	app := fiber.New()

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HandleGetUsers)        // Get all users
	apiV1.Get("/user/:id", api.HandleGetUserById) // Get user by Id
	app.Listen(*listenAddr)
}
