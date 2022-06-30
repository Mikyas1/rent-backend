package properties

import (
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/utils/errors"
)

type Repository interface {
	AddProperty(property *Property) *errors.RestError
	GetTopProperties(p entities.Pagination, fDto PropertyFilterCriteria) (*PropertyListReturnDto, *errors.RestError) // rentable properties, this can get advanced
	GetPropertiesByOwner(ownerId uuid.UUID) ([]Property, *errors.RestError)
	GetPropertyById(propertyId uuid.UUID) (*Property, *errors.RestError)
	RemoverProperty(propertyId, ownerId uuid.UUID) *errors.RestError
	RejectProperty(propertyId uuid.UUID) *errors.RestError
	ApproveProperty(propertyId uuid.UUID) *errors.RestError
	GetPendingProperties() ([]Property, *errors.RestError)
}
