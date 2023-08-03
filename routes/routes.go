package routes

import (
	"github.com/Izzy4999/fibre_test/controller"
	"github.com/Izzy4999/fibre_test/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/login", controller.Login)
	app.Get("/cashiers/:cashierId/logout", controller.Logout)
	// app.Post("/cashiers/", nil)

	app.Post("/cashiers", controller.CreateCashier)
	app.Get("/allcashiers", middleware.NewMiddleware(), controller.CashierList)
	app.Get("/cashiers/me", middleware.NewMiddleware(), controller.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", middleware.NewMiddleware(), controller.DeleteCashier)
	app.Put("/cashiers/:cashierId", middleware.NewMiddleware(), controller.UpdateCashier)
}
