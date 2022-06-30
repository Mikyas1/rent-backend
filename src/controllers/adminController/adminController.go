package adminController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rent/src/services/admin"
	"rent/src/utils/errors"
	"rent/src/utils/http_utils"
)

type AdminController struct {
	AdminService admin.AdminService
}

func (ac AdminController) GetStatistics(c *fiber.Ctx) error {
	res, err := ac.AdminService.GetStatistics()
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (ac AdminController) GetPendingProperties(c *fiber.Ctx) error {
	res, err := ac.AdminService.GetPendingProperties()
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (ac AdminController) ApproveProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	propId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	err := ac.AdminService.ApproveProperty(propId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, "ok")
}

func (ac AdminController) RejectProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	propId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	err := ac.AdminService.RejectProperty(propId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, "ok")
}

func NewAdminController(adminService admin.AdminService) AdminController {
	return AdminController{
		adminService,
	}
}
