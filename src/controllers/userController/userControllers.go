package userController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rent/src/services/user"
	"rent/src/utils/errors"
	"rent/src/utils/http_utils"
)

type UserController struct {
	UserService user.UserService
}

func (uc UserController) RegisterUser(c *fiber.Ctx) error {
	dto := user.RegisterDto{}
	if err := c.BodyParser(&dto); err != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Bad Request", ""))
	}
	res, err := uc.UserService.RegisterUser(dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (uc UserController) LoginUser(c *fiber.Ctx) error {
	dto := user.LoginDto{}
	if err := c.BodyParser(&dto); err != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Bad Request", ""))
	}
	res, err := uc.UserService.LoginUser(dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (uc UserController) CreateRenterProfile(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	dto := user.RenterProfileDto{}
	if err := c.BodyParser(&dto); err != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Bad Request", ""))
	}
	res, err := uc.UserService.AddRenterProfile(userId, dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (uc UserController) GetRenterProfile(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	res, err := uc.UserService.GetRenterProfile(userId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func NewUserController(userService user.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}
