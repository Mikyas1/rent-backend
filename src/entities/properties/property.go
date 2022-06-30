package properties

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"rent/src/entities"
	"rent/src/entities/users"
)

type Property struct {
	entities.Model

	OwnerId  uuid.UUID  `json:"owner_id" gorm:"type:char(36);index;not null"`
	Owner    users.User `json:"owner" gorm:"foreignKey:OwnerId;references:UUID"`
	Approved bool       `json:"approved" gorm:"default:false"`

	Price        float32  `json:"price" gorm:"not null"`
	PropertySize *float32 `json:"property_size"`
	LandArea     *float32 `json:"land_area"`
	Bedrooms     int      `json:"bedrooms" gorm:"index,not null"`
	Bathrooms    int      `json:"bathrooms" gorm:"index,not null"`
	GarageNo     *int     `json:"garage_no" gorm:"index"`
	Floor        *int     `json:"floor" gorm:"index"`
	YearBuilt    *int     `json:"year_built" gorm:"index"`
	Description  string   `json:"description"`
	Furnished    bool     `json:"furnished"`

	AddressId uuid.UUID `json:"address_id" gorm:"type:char(36);index;not null"`
	Address   Address   `json:"address" gorm:"foreignKey:AddressId;references:UUID"`

	Features       Features       `json:"features" gorm:"type:features[]"`
	PropertyType   PropertyType   `json:"property_type" gorm:"not null,index,type:property_types"`
	PropertyStatus PropertyStatus `json:"property_status" gorm:"not null,index,type:property_status"`

	Images pq.StringArray `json:"images" gorm:"type:text[]"`

	MaxFamilySize  *int `json:"max_family_size"`
	MaxPetSize     *int `json:"max_pet_size"`
	MaxVehicleSize *int `json:"max_vehicle_size"`
	MaxKidsSize    *int `json:"max_kids_size"`

	BluePrint pq.StringArray `json:"blue_print" gorm:"type:text[]"`
	OtherDocs pq.StringArray `json:"other_docs" gorm:"type:text[]"`
}

type Address struct {
	entities.Model
	Longitude *float64 `json:"longitude" gorm:"type:double precision;precision"`
	Latitude  *float64 `json:"latitude" gorm:"type:double precision;precision"`
	County    Country  `json:"county" gorm:"not null,type:countries"`
	Region    Region   `json:"region" gorm:"not null,type:regions"`
	City      City     `json:"city" gorm:"not null,type:cities"`
	Woreda    *int     `json:"woreda"`
	Kebele    *int     `json:"kebele"`

	Area        *string `json:"area"`
	Description string  `json:"description"`
}

type PropertyFilterCriteria struct {
	MinPrice     *float32      `json:"min_price" validate:"required,min=1,max=999999"`
	MaxPrice     *float32      `json:"max_price" validate:"required,min=1,max=999999"`
	Bedrooms     *int          `json:"bedrooms" validate:"min=0"`
	Bathrooms    *int          `json:"bathrooms" validate:"min=0"`
	Furnished    *bool         `json:"furnished"`
	PropertyType *PropertyType `json:"property_type" validate:"required"`
	County       *Country      `json:"county" validate:"required"`
	Region       *Region       `json:"region" validate:"required"`
	City         *City         `json:"city" validate:"required"`
	Features     *Features     `json:"features"`
}

func (f PropertyFilterCriteria) GetQuery() string {
	query := " True "
	if f.MinPrice != nil {
		query = query + fmt.Sprintf(" AND properties.price > %v ", *f.MinPrice)
	}
	if f.MaxPrice != nil {
		query = query + fmt.Sprintf(" AND properties.price < %v ", *f.MaxPrice)
	}
	if f.Bedrooms != nil {
		query = query + fmt.Sprintf(" AND properties.bedrooms = %v ", *f.Bedrooms)
	}
	if f.Bathrooms != nil {
		query = query + fmt.Sprintf(" AND properties.bathrooms = %v ", *f.Bathrooms)
	}
	if f.Furnished != nil {
		query = query + fmt.Sprintf(" AND properties.furnished = %v ", *f.Furnished)
	}
	if f.PropertyType != nil {
		query = query + fmt.Sprintf(" AND properties.property_type = '%s' ", *f.PropertyType)
	}
	if f.County != nil {
		query = query + fmt.Sprintf(" AND addresses.county = '%v' ", *f.County)
	}
	if f.Region != nil {
		query = query + fmt.Sprintf(" AND addresses.region = '%v' ", *f.Region)
	}
	if f.City != nil {
		query = query + fmt.Sprintf(" AND addresses.city = '%v' ", *f.City)
	}

	return query
}

type PropertyListReturnDto struct {
	Total        int64      `json:"total"`
	Properties   []Property `json:"properties"`
	Page         int        `json:"page"`
	LimitPerPage int        `json:"limit_per_page"`
}
