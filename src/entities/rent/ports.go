package rent

import (
	"github.com/google/uuid"
	"rent/src/utils/errors"
)

type Repository interface {
	CreateRentRequest(request *RentRequest) (*RentRequest, *errors.RestError)
	MyRentRequestsAsOwner(ownerId uuid.UUID) ([]RentRequest, *errors.RestError)
	MyRentRequestsAsRenter(renterId uuid.UUID) ([]RentRequest, *errors.RestError)
	RentRequestDetail(requestId uuid.UUID) (*RentRequest, *errors.RestError)
	RejectRentRequest(ownerId, rqId uuid.UUID) *errors.RestError
	AcceptRentRequest(ownerId, rentId uuid.UUID) (*Rent, *errors.RestError)
	MyRentsAsOwner(ownerId uuid.UUID) ([]Rent, *errors.RestError)
	MyRentsAsRenter(renterId uuid.UUID) ([]Rent, *errors.RestError)
	RentDetail(rentId uuid.UUID) (*Rent, *errors.RestError)
	GetStatistics() (map[string]interface{}, *errors.RestError)
}
