package admin

import (
	"github.com/google/uuid"
	"rent/src/entities/properties"
	"rent/src/utils/errors"
)

type AdminService interface {
	GetStatistics() (map[string]interface{}, *errors.RestError)
	GetPendingProperties() ([]properties.Property, *errors.RestError)
	ApproveProperty(pId uuid.UUID) *errors.RestError
	RejectProperty(pId uuid.UUID) *errors.RestError
}
