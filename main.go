package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ssss-tantalum/fiber-auth-jwt-sample/database"
	"github.com/ssss-tantalum/fiber-auth-jwt-sample/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetUpRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
