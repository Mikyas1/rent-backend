package messageController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rent/src/services/message"
	"rent/src/utils/errors"
	"rent/src/utils/http_utils"
)

type MessageController struct {
	MessageService message.MessageService
}

func (mc MessageController) GetConversations(c *fiber.Ctx) error {
	userId := c.Locals("id").(uuid.UUID)
	cns, err := mc.MessageService.GetConversationHeads(userId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, cns)
}

func (mc MessageController) GetMessages(c *fiber.Ctx) error {
	id := c.Params("id")
	conId, parsErr := uuid.Parse(id)
	if parsErr != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Id not found", ""))
	}
	mgs, err := mc.MessageService.GetMessages(conId)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, mgs)
}

func (mc MessageController) SendMessage(c *fiber.Ctx) error {
	senderId := c.Locals("id").(uuid.UUID)
	dto := message.SendMessageDto{}
	if err := c.BodyParser(&dto); err != nil {
		return http_utils.RespondError(c, *errors.NewBadRequestError("Bad Request", ""))
	}
	dto.SenderId = senderId
	err := mc.MessageService.SendMessage(dto)
	if err != nil {
		return http_utils.RespondError(c, *err)
	}
	return http_utils.RespondSuccess(c, "ok")
}

func NewMessageController(messageService message.MessageService) MessageController {
	return MessageController{
		MessageService: messageService,
	}
}
