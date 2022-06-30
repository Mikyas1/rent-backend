package rent

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"rent/src/entities/rent"
	"rent/src/utils/errors"
)

type RentRequestDto struct {
	RequesterId  uuid.UUID        `json:"requester_id" validate:"required"`
	PropertyId   uuid.UUID        `json:"property_id" validate:"required"`
	OwnerId      uuid.UUID        `json:"owner_id" validate:"required"`
	RentDuration int              `json:"rent_duration" validate:"required,min=1,max=999"`
	PaymentType  rent.PaymentType `json:"payment_type" validate:"required"`
}

func (d RentRequestDto) Validate() *errors.RestError {
	v := validator.New()
	err := v.Struct(d)
	if err != nil {
		return errors.NewBadRequestError("Request error", err.Error())
	}
	return nil
}

func (d *RentRequestDto) CreateRentRequest() rent.RentRequest {
	var rnt rent.RentRequest
	rnt.UUID = uuid.New()
	rnt.RequesterId = d.RequesterId
	rnt.PropertyId = d.PropertyId
	rnt.OwnerId = d.OwnerId
	rnt.Active = true
	rnt.RentDuration = d.RentDuration
	rnt.Status = rent.Pending
	rnt.PaymentType = d.PaymentType
	return rnt
}
