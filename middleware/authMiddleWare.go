package middleware

import (
	"fmt"
	"os"

	"github.com/Izzy4999/fibre_test/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := initializers.Store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized to access",
		})
	}

	key := os.Getenv("AUTH_KEY")

	if sess.Get(key) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized yet",
		})
	}

	token := sess.Get("token")

	if token == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "not token present",
		})
	}

	tokenString := fmt.Sprintf("%v", token)

	tokenVal, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, c.Status(400).JSON(fiber.Map{
				"message": "Invalid token",
				"success": false,
			})
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := tokenVal.Claims.(jwt.MapClaims); ok && tokenVal.Valid {
		c.Locals("id", claims["id"])
		fmt.Println(claims["id"])
	} else {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.Next()
}
