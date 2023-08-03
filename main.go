package main

import (
	"github.com/Izzy4999/fibre_test/initializers"
	routes "github.com/Izzy4999/fibre_test/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := fiber.New()
	initializers.Setup()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	routes.Setup(app)

	app.Listen(":5000")
}
