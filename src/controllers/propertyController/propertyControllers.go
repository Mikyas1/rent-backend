package propertyController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rent/src/services/property"
	"rent/src/utils/errors"
	"rent/src/utils/http_utils"
)

type PropertyController struct {
	PropertyService property.PropertyService
}

func (pc PropertyController) AddProperty(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	dto := property.AddPropertyDto{}
	data, mErr := c.MultipartForm()
	if mErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Multipart form data expected", ""))
	}
	err := dto.Create(data.Value)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	err = dto.CreateImages(data.File)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	dto.Owner = userId

	prop, err := pc.PropertyService.AddProperty(dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, prop)

}

func NewPropertyController(propertyService property.PropertyService) PropertyController {
	return PropertyController{
		PropertyService: propertyService,
	}
}
