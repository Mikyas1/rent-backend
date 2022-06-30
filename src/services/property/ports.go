package property

import (
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/utils/errors"
)

type PropertyService interface {
	AddProperty(dto AddPropertyDto) (*properties.Property, *errors.RestError)
	GetTopProperties(p entities.Pagination, fDto properties.PropertyFilterCriteria) (*properties.PropertyListReturnDto, *errors.RestError)
	GetOwnerProperties(ownerId uuid.UUID) ([]properties.Property, *errors.RestError)
	DeleteProperty(propertyId, ownerId uuid.UUID) *errors.RestError
	GetPropertyDetail(propertyId uuid.UUID) (*properties.Property, *errors.RestError)
	GetPropertyOptions() (map[string]interface{}, *errors.RestError)
	//EditProperty
}
