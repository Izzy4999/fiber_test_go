package main

import (
	"github.com/Izzy4999/fibre_test/initializers"
	routes "github.com/Izzy4999/fibre_test/routes"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":5000")
}
