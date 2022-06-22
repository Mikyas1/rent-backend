package property

import (
	"rent/src/entities/properties"
	"rent/src/utils/errors"
)

type PropertyService interface {
	AddProperty(dto AddPropertyDto) (*properties.Property, *errors.RestError)
	GetTopProperties()
	//EditProperty
	//DeleteProperty
}
