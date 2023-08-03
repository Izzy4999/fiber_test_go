package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/Izzy4999/fibre_test/initializers"
	"github.com/Izzy4999/fibre_test/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginParams struct {
	Name     string `json:"name"  validate:"required"`
	Passcode string `json:"passcode"  validate:"required,min=4,max=8"`
}

func Login(c *fiber.Ctx) error {
	v := validator.New()
	var cashierDetails model.Cashier

	cashier := &LoginParams{}

	err := c.BodyParser(cashier)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err,
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

	initializers.DB.Where("name=?", cashier.Name).Find(&cashierDetails)

	if cashierDetails.Id == 0 {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "invalid user or passcode",
		})
	}

	if cashier.Passcode != cashierDetails.Passcode {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "invalid user or passcode",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  cashierDetails.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	sess, err := initializers.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	sess.Set(os.Getenv("AUTH_KEY"), true)
	sess.Set("token", tokenString)

	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func Passcode(c *fiber.Ctx) error {
	return nil
}

func Logout(c *fiber.Ctx) error {
	sess, err := initializers.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "logged out",
		})
	}

	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func HealthCheck(c *fiber.Ctx) error {
	sess, err := initializers.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	auth := sess.Get(os.Getenv("AUTH_KEY"))

	if auth == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}
