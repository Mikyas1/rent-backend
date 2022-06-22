package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rent/src/entities"
	"strconv"
)

func GeneratePaginationFromRequest(c *fiber.Ctx) error {
	limit := 10
	page := 1
	sort := "updated_at desc"
	pageTemp := c.Query("page")
	if "" != pageTemp {
		p, err := strconv.Atoi(pageTemp)
		if err == nil {
			page = p
		}
	}
	limitTemp := c.Query("limit")
	if "" != limitTemp {
		l, err := strconv.Atoi(limitTemp)
		if err == nil {
			limit = l
		}
	}
	tempSort := c.Query("sort")
	if tempSort != "" {
		sort = tempSort
	}
	p := entities.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	c.Locals("pagination", p)
	return c.Next()
}
