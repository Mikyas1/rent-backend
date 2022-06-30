package rent

import (
	"github.com/google/uuid"
	"rent/src/entities/rent"
	"rent/src/utils/errors"
)

type DefaultRentService struct {
	repo rent.Repository
}

func (s DefaultRentService) GetPaymentTypes() ([]rent.PaymentType, *errors.RestError) {
	return rent.AllPaymentTypes, nil
}

func (s DefaultRentService) CreateRentRequest(dto RentRequestDto) (*rent.RentRequest, *errors.RestError) {
	err := dto.Validate()
	if err != nil {
		return nil, err
	}
	if dto.RequesterId.String() == dto.OwnerId.String() {
		return nil, errors.NewForbiddenError("You can't sent request to yourself", "You can't sent request to yourself")
	}
	rq := dto.CreateRentRequest()
	return s.repo.CreateRentRequest(&rq)
}

func (s DefaultRentService) GetRentRequestsAsOwner(ownerId uuid.UUID) ([]rent.RentRequest, *errors.RestError) {
	return s.repo.MyRentRequestsAsOwner(ownerId)
}

func (s DefaultRentService) GetRentRequestsAsRenter(renterId uuid.UUID) ([]rent.RentRequest, *errors.RestError) {
	return s.repo.MyRentRequestsAsRenter(renterId)
}

func (s DefaultRentService) RentRequestDetail(requestId uuid.UUID) (*rent.RentRequest, *errors.RestError) {
	return s.repo.RentRequestDetail(requestId)
}

func (s DefaultRentService) RejectRentRequest(ownerId, rentId uuid.UUID) *errors.RestError {
	return s.repo.RejectRentRequest(ownerId, rentId)
}

func (s DefaultRentService) AcceptRentRequest(ownerId, rentId uuid.UUID) (*rent.Rent, *errors.RestError) {
	return s.repo.AcceptRentRequest(ownerId, rentId)
}

func (s DefaultRentService) GetRentsAsOwner(ownerId uuid.UUID) ([]rent.Rent, *errors.RestError) {
	return s.repo.MyRentsAsOwner(ownerId)
}

func (s DefaultRentService) GetRentsAsRenter(renterId uuid.UUID) ([]rent.Rent, *errors.RestError) {
	return s.repo.MyRentsAsRenter(renterId)
}

func (s DefaultRentService) RentDetail(rentId uuid.UUID) (*rent.Rent, *errors.RestError) {
	return s.repo.RentDetail(rentId)
}

func NewDefaultRentService(repo rent.Repository) RentService {
	return DefaultRentService{
		repo: repo,
	}
}
