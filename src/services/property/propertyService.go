package property

import (
	"context"
	"github.com/google/uuid"
	"os"
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/utils/errors"
)

type DefaultPropertyService struct {
	repo    properties.Repository
	storage entities.Storage
}

func (s DefaultPropertyService) AddProperty(dto AddPropertyDto) (*properties.Property, *errors.RestError) {
	err := dto.Validate()
	if err != nil {
		return nil, err
	}

	pr := dto.CreatePropertyFromDto()
	var images []string
	for _, img := range dto.Images {
		err := s.storage.SaveImage(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), img.Name, img.Image)
		if err != nil {
			return nil, err
		}
		images = append(images, img.Name)
	}
	pr.Images = images

	var otherDocs []string
	for _, img := range dto.OtherDocs {
		err := s.storage.SaveImage(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), img.Name, img.Image)
		if err != nil {
			return nil, err
		}
		otherDocs = append(otherDocs, img.Name)
	}
	pr.OtherDocs = otherDocs

	var bluePrints []string
	for _, img := range dto.BluePrint {
		err := s.storage.SaveImage(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), img.Name, img.Image)
		if err != nil {
			return nil, err
		}
		bluePrints = append(bluePrints, img.Name)
	}
	pr.BluePrint = bluePrints

	err = s.repo.AddProperty(pr)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (s DefaultPropertyService) GetTopProperties(p entities.Pagination, fDto properties.PropertyFilterCriteria) (*properties.PropertyListReturnDto, *errors.RestError) {
	return s.repo.GetTopProperties(p, fDto)
}

func (s DefaultPropertyService) GetOwnerProperties(ownerId uuid.UUID) ([]properties.Property, *errors.RestError) {
	return s.repo.GetPropertiesByOwner(ownerId)
}
func (s DefaultPropertyService) DeleteProperty(propertyId, ownerId uuid.UUID) *errors.RestError {
	pr, err := s.repo.GetPropertyById(propertyId)
	if err != nil {
		return err
	}
	if pr.PropertyStatus == properties.Rented {
		return errors.NewForbiddenError("You can't delete a rented property", "You can't delete a rented property")
	}
	return s.repo.RemoverProperty(propertyId, ownerId)
}

func (s DefaultPropertyService) GetPropertyDetail(propertyId uuid.UUID) (*properties.Property, *errors.RestError) {
	return s.repo.GetPropertyById(propertyId)
}

func (s DefaultPropertyService) GetPropertyOptions() (map[string]interface{}, *errors.RestError) {
	features := properties.AllFeatures
	regions := properties.AllRegions
	cities := properties.AllCities
	countries := properties.AllCountries
	propertyStatus := properties.AllPropertyStatus
	propertyTypes := properties.AllPropertyTypes

	return map[string]interface{}{"features": features,
		"regions":         regions,
		"cities":          cities,
		"countries":       countries,
		"property_status": propertyStatus,
		"property_types":  propertyTypes}, nil
}

func NewDefaultPropertyService(repo properties.Repository, storage entities.Storage) PropertyService {
	return DefaultPropertyService{
		repo:    repo,
		storage: storage,
	}
}
