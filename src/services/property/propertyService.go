package property

import (
	"context"
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
	err = s.repo.AddProperty(pr)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (s DefaultPropertyService) GetTopProperties() {

}

func NewDefaultPropertyService(repo properties.Repository, storage entities.Storage) PropertyService {
	return DefaultPropertyService{
		repo:    repo,
		storage: storage,
	}
}
