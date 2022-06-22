package errors

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func GormError(err error, record string) *RestError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewNotFoundError(fmt.Sprintf("%s not found", record), "")
	}
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return NewBadRequestError(fmt.Sprintf("%s already exists", record), "")
	}
	if strings.Contains(err.Error(), "invalid input value for enum") {
		return NewBadRequestError("Unsupported Enum provided", "")
	}
	return NewInternalServerError(err.Error(), "")
}
