package routes

import (
	"github.com/Izzy4999/fibre_test/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/cashiers/login", nil)
	app.Get("/cashiers/:cashierId/logout", nil)
	app.Post("/cashiers/:cashierId/passcode", nil)

	app.Post("/cashiers", controller.CreateCashier)
	app.Get("/cashiers", controller.CashierList)
	app.Get("/cashiers/:cashierId", controller.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)
	app.Put("/cashiers/:cashierId", controller.UpdateCashier)
}
