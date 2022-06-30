package rentControllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rent/src/services/rent"
	"rent/src/utils/errors"
	"rent/src/utils/http_utils"
)

type RentController struct {
	RentService rent.RentService
}

func (rc RentController) GetPaymentTypes(c *fiber.Ctx) error {
	res, err := rc.RentService.GetPaymentTypes()
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (rc RentController) CreateRentRequest(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	dto := rent.RentRequestDto{}
	if err := c.BodyParser(&dto); err != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Bad Request", ""))
	}
	dto.RequesterId = userId
	res, err := rc.RentService.CreateRentRequest(dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (rc RentController) GetRentRequestAsOwner(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	res, err := rc.RentService.GetRentRequestsAsOwner(userId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (rc RentController) GetRentRequestAsRenter(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	res, err := rc.RentService.GetRentRequestsAsRenter(userId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (rc RentController) RentRequestDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	rqId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	res, err := rc.RentService.RentRequestDetail(rqId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, res)
}

func (rc RentController) RejectRentRequest(c *fiber.Ctx) error {
	ownerId := c.Locals("id").(uuid.UUID)
	id := c.Params("id")
	rqId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	err := rc.RentService.RejectRentRequest(ownerId, rqId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, "ok")
}

func (rc RentController) AcceptRentRequest(c *fiber.Ctx) error {
	ownerId := c.Locals("id").(uuid.UUID)
	id := c.Params("id")
	rqId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	rnt, err := rc.RentService.AcceptRentRequest(ownerId, rqId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, rnt)
}

func (rc RentController) AllRentsAsOwner(c *fiber.Ctx) error {
	ownerId := c.Locals("id").(uuid.UUID)
	rnts, err := rc.RentService.GetRentsAsOwner(ownerId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, rnts)
}

func (rc RentController) AllRentsAsRenter(c *fiber.Ctx) error {
	retnerId := c.Locals("id").(uuid.UUID)
	rnts, err := rc.RentService.GetRentsAsRenter(retnerId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, rnts)
}

func (rc RentController) GetRentDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	rId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	rnt, err := rc.RentService.RentDetail(rId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, rnt)
}

func NewRentController(service rent.RentService) RentController {
	return RentController{
		RentService: service,
	}
}
