package properties

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/entities/rent"
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

func (r PropertyRepository) GetTopProperties(p entities.Pagination, fDto properties.PropertyFilterCriteria) (*properties.PropertyListReturnDto, *errors.RestError) {
	var ps []properties.Property

	query := fDto.GetQuery()
	result := r.db.Scopes(Paginate(p)).
		Joins("left join addresses on addresses.id = properties.address_id").
		Where("approved = ?", true).
		Where("property_status = ?", properties.Approved.String()).
		Where(query).
		Preload("Address").
		Find(&ps)
	if result.Error != nil {
		logger.Error("Error when getting Properties " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}

	var count int64
	result = r.db.Model(&properties.Property{}).
		Joins("left join addresses on addresses.id = properties.id").
		Where("approved = ?", true).
		Where("property_status = ?", properties.Approved.String()).
		Where(query).
		Count(&count)
	if result.Error != nil {
		logger.Error("Error counting properties " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}

	prs := properties.PropertyListReturnDto{
		Total:        count,
		Properties:   ps,
		Page:         p.Page,
		LimitPerPage: p.Limit,
	}
	return &prs, nil
}

func (r PropertyRepository) GetPropertiesByOwner(ownerId uuid.UUID) ([]properties.Property, *errors.RestError) {
	var res []properties.Property
	result := r.db.Where("owner_id = ?", ownerId.String()).
		Preload("Address").
		Find(&res)
	if result.Error != nil {
		logger.Error("Error when getting Properties by owner id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}
	return res, nil
}

func (r PropertyRepository) GetPropertyById(propertyId uuid.UUID) (*properties.Property, *errors.RestError) {
	var res properties.Property
	result := r.db.Where("id = ?", propertyId.String()).
		Preload("Owner").Preload("Address").
		First(&res)
	if result.Error != nil {
		logger.Error("Error when getting Property by id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}
	return &res, nil
}

func (r PropertyRepository) RemoverProperty(propertyId, ownerId uuid.UUID) *errors.RestError {
	result := r.db.Where("property_status <> ?", properties.Rented.String()).
		Where("id = ?", propertyId.String()).
		Where("owner_id = ?", ownerId.String()).
		Delete(&properties.Property{})
	if result.Error != nil {
		logger.Error("Error deleting Property by id and ownerId " + result.Error.Error())
		return errors.GormError(result.Error, "Property")
	}
	result = r.db.Where("property_id = ?", propertyId.String()).
		Delete(&rent.RentRequest{})
	if result.Error != nil {
		logger.Error("Error deleting Rent Requests by property id" + result.Error.Error())
		return errors.GormError(result.Error, "Rent Request")
	}
	return nil
}

func (r PropertyRepository) RejectProperty(propertyId uuid.UUID) *errors.RestError {
	result := r.db.Model(&properties.Property{}).Where("id = ?", propertyId.String()).
		Updates(map[string]interface{}{"approved": true, "property_status": string(properties.Rejected)})
	if result.Error != nil {
		logger.Error("Error while Approving Property" + result.Error.Error())
		return errors.GormError(result.Error, "Property")
	}
	return nil
}
func (r PropertyRepository) ApproveProperty(propertyId uuid.UUID) *errors.RestError {
	result := r.db.Model(&properties.Property{}).Where("id = ?", propertyId.String()).
		Updates(map[string]interface{}{"approved": true, "property_status": string(properties.Approved)})
	if result.Error != nil {
		logger.Error("Error while Approving Property" + result.Error.Error())
		return errors.GormError(result.Error, "Property")
	}
	return nil
}
func (r PropertyRepository) GetPendingProperties() ([]properties.Property, *errors.RestError) {
	var ps []properties.Property
	result := r.db.Where("property_status = ?", properties.PendingApproval.String()).
		Preload("Owner").Preload("Address").
		Find(&ps)
	if result.Error != nil {
		logger.Error("Error getting pending approval Properties " + result.Error.Error())
		return nil, errors.GormError(result.Error, "Property")
	}
	return ps, nil
}

func NewPropertyRepository(db *gorm.DB) properties.Repository {
	return PropertyRepository{
		db: db,
	}
}
