package property

import (
	"github.com/google/uuid"
	"rent/src/adaptors/properties"
	"rent/src/adaptors/storage"
	properties2 "rent/src/entities/properties"
	"testing"
)

func createCorrectPropertyDto(imageNames ...string) AddPropertyDto {
	var imgDtos []ImageDto
	for _, name := range imageNames {
		imgDtos = append(imgDtos, ImageDto{Name: name, Image: nil})
	}

	return AddPropertyDto{
		Owner:        uuid.New(),
		Price:        2000,
		Bedrooms:     2,
		Bathrooms:    1,
		Furnished:    true,
		Features:     []properties2.Feature{properties2.Garage, properties2.Internet},
		PropertyType: properties2.Studio,
		County:       properties2.Ethiopia,
		Region:       properties2.AddisAbabaR,
		City:         properties2.AddisAbaba,
		Images:       imgDtos,
	}
}

func CreateProperties() []properties2.Property {
	res := []properties2.Property{
		properties2.Property{
			OwnerId:        uuid.New(),
			Approved:       true,
			Price:          200,
			Bedrooms:       3,
			Bathrooms:      2,
			AddressId:      uuid.New(),
			PropertyType:   properties2.Studio,
			PropertyStatus: properties2.Approved,
		},
	}
	return res
}

func TestAddPropertyDto_Validate(t *testing.T) {
	dto := createCorrectPropertyDto()
	err := dto.Validate()
	if err != nil {
		t.Errorf("Validateor should pass")
	}
}

func TestDefaultPropertyService_AddProperty(t *testing.T) {
	service := NewDefaultPropertyService(properties.NewPropertyRepositoryStub(CreateProperties()), storage.NewStorageStub())
	imgNames := []string{"img1", "img2"}
	dto := createCorrectPropertyDto(imgNames...)
	prop, err := service.AddProperty(dto)
	if err != nil {
		t.Errorf("Property should have been added and no error should be generated")
	}

	if prop == nil {
		t.Errorf("Property should be created and returned")
	}

	if prop.UUID.String() == "00000000-0000-0000-0000-000000000000" {
		t.Errorf("Property should have UUID as an id")
	}

	for i, iName := range prop.Images {
		if iName != imgNames[i] {
			t.Errorf("Images Name expected %s got %s from created property", imgNames[i], iName)
		}
	}

	if prop.PropertyStatus != properties2.PendingApproval {
		t.Errorf("New added property should be pending approval")
	}

}
