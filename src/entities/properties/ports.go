package properties

import (
	"rent/src/entities"
	"rent/src/utils/errors"
)

type Repository interface {
	AddProperty(property *Property) *errors.RestError
	GetTopProperties(p entities.Pagination) (*[]Property, *errors.RestError) // rentable properties, this can get advanced
	// get properties by owner
	// get property by id
	// delete property
	// edit property
}
