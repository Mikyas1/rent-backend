package properties

import (
	"gorm.io/gorm"
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/logger"
	"rent/src/utils/errors"
)

type PropertyRepository struct {
	db *gorm.DB
}

func (r PropertyRepository) AddProperty(property *properties.Property) *errors.RestError {
	result := r.db.Create(property)
	if result.Error != nil {
		logger.Error("Error while saving Regular User " + result.Error.Error())
		return errors.GormError(result.Error, "Property")
	}
	return nil
}

func Paginate(p entities.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.Limit
		return db.Offset(offset).Limit(p.Limit).Order(p.Sort)
	}
}

func (r PropertyRepository) GetTopProperties(p entities.Pagination) (*[]properties.Property, *errors.RestError) {
	var ps []properties.Property
	result := r.db.Scopes(Paginate(p)).
		Find(&ps)
	if result.Error != nil {
		logger.Error("Error when getting Properties " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}

	return nil, nil
}

func NewPropertyRepository(db *gorm.DB) properties.Repository {
	return PropertyRepository{
		db: db,
	}
}
