package objectStorageProxy

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rent/src/services/objectStorage"
	"rent/src/utils/http_utils"
)

type ObjectStorageHandler struct {
	ObjectService objectStorage.ObjectService
}

func (h ObjectStorageHandler) File(c *fiber.Ctx) error {
	objectName := c.Params("object_name")

	buf, contentType, err := h.ObjectService.GetObject(objectName)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	c.Set(fiber.HeaderContentType, *contentType)
	c.Set(fiber.HeaderContentLength, fmt.Sprintf("%d", len(buf)))

	return c.Send(buf)
}

func NewObjectStorageHandler(service objectStorage.ObjectService) ObjectStorageHandler {
	return ObjectStorageHandler{
		ObjectService: service,
	}
}
