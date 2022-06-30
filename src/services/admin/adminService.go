package admin

import (
	"github.com/google/uuid"
	properties2 "rent/src/entities/properties"
	rent2 "rent/src/entities/rent"
	"rent/src/utils/errors"
)

type DefaultAdminService struct {
	propertyRepo properties2.Repository
	rentRepo     rent2.Repository
}

func (s DefaultAdminService) GetStatistics() (map[string]interface{}, *errors.RestError) {
	return s.rentRepo.GetStatistics()
}

func (s DefaultAdminService) GetPendingProperties() ([]properties2.Property, *errors.RestError) {
	return s.propertyRepo.GetPendingProperties()
}

func (s DefaultAdminService) ApproveProperty(pId uuid.UUID) *errors.RestError {
	return s.propertyRepo.ApproveProperty(pId)
}

func (s DefaultAdminService) RejectProperty(pId uuid.UUID) *errors.RestError {
	return s.propertyRepo.RejectProperty(pId)
}

func NewDefaultAdminService(propertyRepo properties2.Repository, rentRepo rent2.Repository) AdminService {
	return DefaultAdminService{
		propertyRepo: propertyRepo,
		rentRepo:     rentRepo,
	}
}
