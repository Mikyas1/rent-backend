package properties

import (
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/utils/errors"
)

type PropertyRepositoryStub struct {
	Properties []properties.Property
}

func (r PropertyRepositoryStub) AddProperty(property *properties.Property) *errors.RestError {
	return nil
}

func (r PropertyRepositoryStub) GetTopProperties(p entities.Pagination, fDto properties.PropertyFilterCriteria) (*properties.PropertyListReturnDto, *errors.RestError) {
	return nil, nil
}

func (r PropertyRepositoryStub) GetPropertiesByOwner(ownerId uuid.UUID) ([]properties.Property, *errors.RestError) {
	return r.Properties, nil
}
func (r PropertyRepositoryStub) GetPropertyById(propertyId uuid.UUID) (*properties.Property, *errors.RestError) {
	return &r.Properties[1], nil
}
func (r PropertyRepositoryStub) RemoverProperty(propertyId, ownerId uuid.UUID) *errors.RestError {
	return nil
}

func (r PropertyRepositoryStub) RejectProperty(propertyId uuid.UUID) *errors.RestError {
	return nil
}
func (r PropertyRepositoryStub) ApproveProperty(propertyId uuid.UUID) *errors.RestError {
	return nil
}
func (r PropertyRepositoryStub) GetPendingProperties() ([]properties.Property, *errors.RestError) {
	return nil, nil
}

func NewPropertyRepositoryStub(properties []properties.Property) properties.Repository {
	if len(properties) <= 0 {
		panic("provide at list one property")
	}
	return PropertyRepositoryStub{Properties: properties}
}
