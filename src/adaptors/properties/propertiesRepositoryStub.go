package properties

import (
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/utils/errors"
)

type PropertyRepositoryStub struct {
}

func (r PropertyRepositoryStub) AddProperty(property *properties.Property) *errors.RestError {
	return nil
}

func (r PropertyRepositoryStub) GetTopProperties(p entities.Pagination) (*[]properties.Property, *errors.RestError) {
	return nil, nil
}

func NewPropertyRepositoryStub() properties.Repository {
	return PropertyRepository{}
}
