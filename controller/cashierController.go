package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Izzy4999/fibre_test/initializers"
	"github.com/Izzy4999/fibre_test/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Cashier struct {
	Name     string `json:"name" validate:"required"`
	Passcode string `json:"passcode" validate:"required,min=4,max=8"`
}

func CreateCashier(c *fiber.Ctx) error {
	v := validator.New()
	cashier := &Cashier{}

	err := c.BodyParser(cashier)
	if err != nil {
		return c.Status(401).JSON(&fiber.Map{
			"error": err,
		})
	}

	err = v.Struct(*cashier)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("%v %v %v", err.Field(), err.Tag(), err.Param()),
			})
		}
	}

	cashierDetails := model.Cashier{
		Name:      cashier.Name,
		Passcode:  cashier.Passcode,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	result := initializers.DB.Create(&cashierDetails)

	if result.Error != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    cashierDetails,
		"message": "Cashier created successfully",
	})
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}

func CashierList(c *fiber.Ctx) error {
	var cashier []model.Cashier
	var count int64
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	if limit == 0 {
		initializers.DB.Select("*").Offset(skip).Find(&cashier).Count(&count)
	} else {
		initializers.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Successful",
		"data":    cashier,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
