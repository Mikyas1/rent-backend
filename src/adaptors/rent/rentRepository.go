package rent

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rent/src/entities/properties"
	"rent/src/entities/rent"
	"rent/src/entities/users"
	"rent/src/logger"
	"rent/src/utils/errors"
)

type RentRepository struct {
	db *gorm.DB
}

func (r RentRepository) CreateRentRequest(request *rent.RentRequest) (*rent.RentRequest, *errors.RestError) {
	result := r.db.Create(request)
	if result.Error != nil {
		logger.Error("Error while saving Rent Request " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}
	return request, nil
}

func (r RentRepository) MyRentRequestsAsOwner(ownerId uuid.UUID) ([]rent.RentRequest, *errors.RestError) {
	var rq []rent.RentRequest
	result := r.db.Where("owner_id = ?", ownerId.String()).
		Where("active = ?", true).Where("status <> ?", string(rent.Rejected)).
		Preload("Property.Address").Preload("Requester").Find(&rq)
	if result.Error != nil {
		logger.Error("Error while getting Rent Requests by Owner id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}
	return rq, nil
}

func (r RentRepository) MyRentRequestsAsRenter(renterId uuid.UUID) ([]rent.RentRequest, *errors.RestError) {
	var rq []rent.RentRequest
	result := r.db.Where("requester_id = ?", renterId.String()).
		Preload("Property.Address").Preload("Owner").Find(&rq)
	if result.Error != nil {
		logger.Error("Error while getting Rent Requests by Renter id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}
	return rq, nil
}

func (r RentRepository) RentRequestDetail(requestId uuid.UUID) (*rent.RentRequest, *errors.RestError) {
	var rq rent.RentRequest
	result := r.db.Where("id = ?", requestId.String()).
		Preload("Property.Address").Preload("Requester.RenterProfile").
		Preload("Owner").First(&rq)
	if result.Error != nil {
		logger.Error("Error while getting Rent Request by Rent id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}
	return &rq, nil
}

func (r RentRepository) RejectRentRequest(ownerId, rqId uuid.UUID) *errors.RestError {
	var rq rent.RentRequest
	result := r.db.Where("id = ?", rqId.String()).Where("owner_id = ?", ownerId.String()).
		Find(&rq)
	if result.Error != nil {
		logger.Error("Error while getting Rent Request by Rent id and Owner id " + result.Error.Error())
		return errors.GormError(result.Error, "Rent Request")
	}

	result = r.db.Model(&rq).Where("id = ?", rqId.String()).Where("owner_id = ?", ownerId.String()).
		Updates(map[string]interface{}{"active": false, "status": string(rent.Rejected)})
	if result.Error != nil {
		logger.Error("Error while Rejecting Rent Request" + result.Error.Error())
		return errors.GormError(result.Error, "Rent Request")
	}

	return nil
}

func (r RentRepository) AcceptRentRequest(ownerId, rqId uuid.UUID) (*rent.Rent, *errors.RestError) {
	var rq rent.RentRequest
	result := r.db.Where("id = ?", rqId.String()).Where("owner_id = ?", ownerId.String()).
		Find(&rq)
	if result.Error != nil {
		logger.Error("Error while getting Rent Request by Rent id and Owner id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}

	result = r.db.Model(&rq).Where("id = ?", rqId.String()).Where("owner_id = ?", ownerId.String()).
		Updates(map[string]interface{}{"active": false, "status": string(rent.Accepted)})
	if result.Error != nil {
		logger.Error("Error while Rejecting Rent Request" + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}

	rnt := rq.CreateRent()
	result = r.db.Create(&rnt)
	if result.Error != nil {
		logger.Error("Error while saving Rent " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent")
	}

	return &rnt, nil
}

func (r RentRepository) MyRentsAsOwner(ownerId uuid.UUID) ([]rent.Rent, *errors.RestError) {
	var rts []rent.Rent
	result := r.db.Where("owner_id = ?", ownerId.String()).
		Preload("Requester").Preload("RentRequest.Property.Address").Find(&rts)
	if result.Error != nil {
		logger.Error("Error while getting Rent by Owner id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent")
	}
	return rts, nil
}

func (r RentRepository) MyRentsAsRenter(renterId uuid.UUID) ([]rent.Rent, *errors.RestError) {
	var rts []rent.Rent
	result := r.db.Where("requester_id = ?", renterId.String()).
		Preload("Owner").Preload("RentRequest.Property.Address").Find(&rts)
	if result.Error != nil {
		logger.Error("Error while getting Rent by Owner id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent")
	}
	return rts, nil
}

func (r RentRepository) RentDetail(rentId uuid.UUID) (*rent.Rent, *errors.RestError) {
	var rq rent.Rent
	result := r.db.Where("id = ?", rentId.String()).
		Preload("RentRequest.Property.Address").Preload("Owner").
		Preload("Requester.RenterProfile").
		First(&rq)
	if result.Error != nil {
		logger.Error("Error while getting Rent by Rent id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent")
	}
	return &rq, nil
}

func (r RentRepository) GetStatistics() (map[string]interface{}, *errors.RestError) {
	var total_users int64
	var total_properties int64
	var total_rent_requests int64
	var total_rents int64

	result := r.db.Model(&users.User{}).
		Count(&total_users)
	if result.Error != nil {
		logger.Error("Error counting total users " + result.Error.Error())
		return nil, errors.GormError(result.Error, "User")
	}

	result = r.db.Model(&properties.Property{}).
		Count(&total_properties)
	if result.Error != nil {
		logger.Error("Error counting total properties " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}

	result = r.db.Model(&rent.RentRequest{}).
		Count(&total_rent_requests)
	if result.Error != nil {
		logger.Error("Error counting total Rent Requests " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rent Request")
	}

	result = r.db.Model(&rent.RentRequest{}).
		Count(&total_rents)
	if result.Error != nil {
		logger.Error("Error counting total Rents " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Rents")
	}

	return map[string]interface{}{
		"total_users":         total_users,
		"total_properties":    total_properties,
		"total_rent_requests": total_rent_requests,
		"total_rents":         total_rents,
	}, nil

	return nil, nil
}

func NewRentRepository(db *gorm.DB) rent.Repository {
	return RentRepository{
		db: db,
	}
}
