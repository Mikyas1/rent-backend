package rent

import (
	"github.com/google/uuid"
	"rent/src/entities/rent"
	"rent/src/utils/errors"
)

type RentService interface {
	GetPaymentTypes() ([]rent.PaymentType, *errors.RestError)
	CreateRentRequest(dto RentRequestDto) (*rent.RentRequest, *errors.RestError)
	GetRentRequestsAsOwner(ownerId uuid.UUID) ([]rent.RentRequest, *errors.RestError)
	GetRentRequestsAsRenter(renterId uuid.UUID) ([]rent.RentRequest, *errors.RestError)
	RentRequestDetail(requestId uuid.UUID) (*rent.RentRequest, *errors.RestError)

	RejectRentRequest(ownerId, rentId uuid.UUID) *errors.RestError
	AcceptRentRequest(ownerId, rentId uuid.UUID) (*rent.Rent, *errors.RestError)

	GetRentsAsOwner(ownerId uuid.UUID) ([]rent.Rent, *errors.RestError)
	GetRentsAsRenter(renterId uuid.UUID) ([]rent.Rent, *errors.RestError)
	RentDetail(rentId uuid.UUID) (*rent.Rent, *errors.RestError)
}
