package property

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"mime/multipart"
	"rent/src/entities/properties"
	"rent/src/utils/conv"
	"rent/src/utils/errors"
	"rent/src/utils/img"
)

type ImageDto struct {
	Name  string
	Image *multipart.FileHeader
}

type AddPropertyDto struct {
	Owner        uuid.UUID `json:"user_id"`
	Price        float32   `json:"price" validate:"required,min=1,max=999999"`
	PropertySize *float32  `json:"property_size" validate:"omitempty,min=1"`
	LandArea     *float32  `json:"land_area" validate:"omitempty,min=1"`
	Bedrooms     int       `json:"bedrooms" validate:"min=0"`
	Bathrooms    int       `json:"bathrooms" validate:"min=0"`
	GarageNo     *int      `json:"garage_no" validate:"omitempty,min=0"`
	Floor        *int      `json:"floor" validate:"omitempty,min=0"`
	YearBuilt    *int      `json:"year_built" validate:"omitempty,min=1890,max=2022"`
	Description  string    `json:"description"`
	Furnished    bool      `json:"furnished"`

	Features     properties.Features     `json:"features"`
	PropertyType properties.PropertyType `json:"property_type" validate:"required"`

	Images []ImageDto `json:"images"`

	Longitude          *float64           `json:"longitude" validate:"omitempty,longitude"`
	Latitude           *float64           `json:"latitude" validate:"omitempty,latitude"`
	County             properties.Country `json:"county" validate:"required"`
	Region             properties.Region  `json:"region" validate:"required"`
	City               properties.City    `json:"city" validate:"required"`
	Area               *string            `json:"area"`
	AddressDescription string             `json:"address_description"`
}

func (d *AddPropertyDto) Create(data map[string][]string) *errors.RestError {
	v := validator.New()
	temp := map[string]interface{}{}
	for key, val := range data {
		//temp[key] = strings.Join(val, " ")
		temp[key] = val
	}

	vlr := map[string]interface{}{
		"price":         "required,min=1,max=999999",
		"property_size": "omitempty,min=1",
		"land_area":     "omitempty,min=1",
		"bedrooms":      "min=0",
		"bathrooms":     "min=0",
		"garage_no":     "omitempty,min=0",
		"floor":         "omitempty,min=0",
		"year_built":    "omitempty,min=1890,max=2022",
		"furnished":     "required",
		"property_type": "required",
		"longitude":     "omitempty,longitude",
		"latitude":      "omitempty,latitude",
		"county":        "required",
		"region":        "required",
		"city":          "required",
	}

	res := v.ValidateMap(temp, vlr)
	if len(res) != 0 {
		for key, val := range res {
			return errors.NewBadRequestError(""+val.(error).Error(), key)
		}
	}

	conv.ToFloat32(&d.Price, temp["price"])
	conv.ToFloat32(d.PropertySize, temp["property_size"])
	conv.ToFloat32(d.LandArea, temp["land_area"])
	conv.ToFloat64(d.Longitude, temp["longitude"])
	conv.ToFloat64(d.Latitude, temp["latitude"])
	conv.ToInt(&d.Bedrooms, temp["bedrooms"])
	conv.ToInt(&d.Bathrooms, temp["bathrooms"])
	conv.ToInt(d.GarageNo, temp["garage_no"])
	conv.ToInt(d.Floor, temp["floor"])
	conv.ToInt(d.YearBuilt, temp["year_built"])
	conv.ToBool(&d.Furnished, temp["furnished"])

	if pt, ok := temp["property_type"]; ok {
		d.PropertyType = properties.PropertyType(pt.([]string)[0])
	}

	if ct, ok := temp["county"]; ok {
		d.County = properties.Country(ct.([]string)[0])
	}

	if rg, ok := temp["region"]; ok {
		d.Region = properties.Region(rg.([]string)[0])
	}

	if ct, ok := temp["city"]; ok {
		d.City = properties.City(ct.([]string)[0])
	}

	if desc, ok := temp["description"]; ok {
		d.Description = desc.([]string)[0]
	}

	if features, ok := temp["features"]; ok {
		for _, fet := range features.([]string) {
			d.Features = append(d.Features, properties.Feature(fet))
		}
	}

	if area, ok := temp["area"]; ok {
		d.Area = area.(*string)
	}

	if addDesc, ok := temp["address_description"]; ok {
		d.AddressDescription = addDesc.(string)
	}

	return nil
}

func (d *AddPropertyDto) CreateImages(data map[string][]*multipart.FileHeader) *errors.RestError {
	var images []*multipart.FileHeader
	for key, val := range data {
		if len(val) > 0 {
			if key == "images" {
				images = append(images, val...)
			}
		}
	}

	for _, val := range images {
		file, name, imgErr := img.ValidateImage(val)
		if imgErr != nil {
			return errors.NewBadRequestError("Image "+imgErr.Error(), "")
		}
		d.Images = append(d.Images, ImageDto{
			Name:  name,
			Image: file,
		})
	}

	return nil
}

func (d AddPropertyDto) Validate() *errors.RestError {
	v := validator.New()
	err := v.Struct(d)
	if err != nil {
		return errors.NewBadRequestError("Request error", err.Error())
	}
	return nil
}

func (d AddPropertyDto) CreatePropertyFromDto() *properties.Property {
	var p properties.Property
	p.UUID = uuid.New()
	p.OwnerId = d.Owner
	p.Price = d.Price
	p.PropertySize = d.PropertySize
	p.LandArea = d.LandArea
	p.Bedrooms = d.Bedrooms
	p.Bathrooms = d.Bathrooms
	p.GarageNo = d.GarageNo
	p.Floor = d.Floor
	p.YearBuilt = d.YearBuilt
	p.Description = d.Description
	p.Furnished = d.Furnished
	p.Features = d.Features
	p.PropertyType = d.PropertyType
	p.Address = properties.Address{
		Longitude:   d.Longitude,
		Latitude:    d.Latitude,
		County:      d.County,
		Region:      d.Region,
		City:        d.City,
		Area:        d.Area,
		Description: d.AddressDescription,
	}
	p.Address.UUID = uuid.New()
	//p.Images = pgtype.TextArray{pgtype.Text{String: "fds"}}

	p.PropertyStatus = properties.PendingApproval

	return &p
}
