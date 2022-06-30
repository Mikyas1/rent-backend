package propertyController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/entities/properties"
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

func (pc PropertyController) GetTopProperties(c *fiber.Ctx) error {
	p := c.Locals("pagination").(entities.Pagination)
	dto := properties.PropertyFilterCriteria{}
	if err := c.BodyParser(&dto); err != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Bad Request", ""))
	}

	ps, err := pc.PropertyService.GetTopProperties(p, dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, ps)
}

func (pc PropertyController) GetOwnerProperties(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	props, err := pc.PropertyService.GetOwnerProperties(userId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, props)
}

func (pc PropertyController) GetPropertyDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	propId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	prop, err := pc.PropertyService.GetPropertyDetail(propId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, prop)
}

func (pc PropertyController) RemoveProperty(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	id := c.Params("id")
	propId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	err := pc.PropertyService.DeleteProperty(propId, userId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, "ok")
}

func (pc PropertyController) GetPropertyOptions(c *fiber.Ctx) error {
	res, err := pc.PropertyService.GetPropertyOptions()
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func NewPropertyController(propertyService property.PropertyService) PropertyController {
	return PropertyController{
		PropertyService: propertyService,
	}
}
