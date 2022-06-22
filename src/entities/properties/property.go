package properties

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"rent/src/entities"
	"rent/src/entities/users"
)

type Property struct {
	entities.Model

	OwnerId  uuid.UUID  `json:"owner_id" gorm:"type:char(36);index;not null"`
	Owner    users.User `json:"-" gorm:"foreignKey:OwnerId;references:UUID"`
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
	Address   Address   `json:"-" gorm:"foreignKey:AddressId;references:UUID"`

	Features       Features       `json:"features" gorm:"type:features[]"`
	PropertyType   PropertyType   `json:"property_type" gorm:"not null,index,type:property_types"`
	PropertyStatus PropertyStatus `json:"property_status" gorm:"not null,index,type:property_status"`

	Images pq.StringArray `json:"images" gorm:"type:text[]"`
}

type Address struct {
	entities.Model
	Longitude   *float64 `json:"longitude" gorm:"type:double precision;precision"`
	Latitude    *float64 `json:"latitude" gorm:"type:double precision;precision"`
	County      Country  `json:"county" gorm:"not null,type:countries"`
	Region      Region   `json:"region" gorm:"not null,type:regions"`
	City        City     `json:"city" gorm:"not null,type:cities"`
	Area        *string  `json:"area"`
	Description string   `json:"description"`
}
