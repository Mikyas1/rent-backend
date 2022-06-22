package http_utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rent/src/utils/errors"
)

func RespondJson(c *fiber.Ctx, statusCode int, body interface{}) error {
	c.Set("Content-Type", "application/json")
	if statusCode != 0 {
		c.Status(statusCode)
	}
	return c.JSON(body)
}

func RespondError(c *fiber.Ctx, error errors.RestError) error {
	if err := RespondJson(c, error.Status, error); err != nil {
		return err
	}
	return nil
}

func RespondSuccess(c *fiber.Ctx, body interface{}) error {
	if err := RespondJson(c, http.StatusOK, body); err != nil {
		return err
	}
	return nil
}
