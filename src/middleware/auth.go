package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"rent/src/utils/errors"
	"rent/src/utils/http_utils"
	"strings"
)

func GetBearerToken(c *fiber.Ctx) error {
	headerValue := c.Get("authorization")
	var token string

	if len(headerValue) > 0 {
		components := strings.SplitN(headerValue, " ", 2)

		if len(components) == 2 && components[0] == "Bearer" {
			token = components[1]
		}
	}
	c.Locals("token", token)
	return c.Next()
}

func ValidateUser(c *fiber.Ctx) error {
	token, tokenErr := jwt.ParseWithClaims(c.Locals("token").(string), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if tokenErr != nil {
		return http_utils.RespondError(c, *errors.NewUnauthorizedError("Unauthorized to access this resource", ""))
	}
	payload := token.Claims.(*jwt.StandardClaims)

	userId, err := uuid.Parse(payload.Id)
	if err != nil {
		return http_utils.RespondError(c, *errors.NewUnauthorizedError("Unauthorized to access this resource", ""))
	}

	c.Locals("id", userId)
	return c.Next()
}
