package main

import (
	"github.com/Izzy4999/fibre_test/initializers"
	"github.com/Izzy4999/fibre_test/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Cashier{})
	initializers.DB.AutoMigrate(&model.Category{})
	initializers.DB.AutoMigrate(&model.Discount{})
	initializers.DB.AutoMigrate(&model.Order{})
	initializers.DB.AutoMigrate(&model.Payment{})
	initializers.DB.AutoMigrate(&model.PaymentType{})
	initializers.DB.AutoMigrate(&model.Product{})
}
