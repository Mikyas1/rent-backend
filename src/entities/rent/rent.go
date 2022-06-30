package rent

import (
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/entities/properties"
	"rent/src/entities/users"
)

type RentRequest struct {
	entities.Model

	RequesterId uuid.UUID  `json:"requester_id" gorm:"type:char(36);uniqueIndex:requester_property;not null"`
	Requester   users.User `json:"requester" gorm:"foreignKey:RequesterId;references:UUID"`

	OwnerId uuid.UUID  `json:"owner_id" gorm:"type:char(36);index;not null"`
	Owner   users.User `json:"owner" gorm:"foreignKey:OwnerId;references:UUID"`

	PropertyId   uuid.UUID           `json:"property_id" gorm:"type:char(36);uniqueIndex:requester_property;not null"`
	Property     properties.Property `json:"property" gorm:"foreignKey:PropertyId;references:UUID"`
	Active       bool                `json:"active" gorm:"default:true"`
	RentDuration int                 `json:"rent_duration" gorm:"not null"`
	Status       RentRequestStatus   `json:"status" gorm:"not null,index,type:rent_request_status"`
	PaymentType  PaymentType         `json:"payment_type" gorm:"type:payment_types"`
}

func (rr *RentRequest) CreateRent() Rent {
	var rnt Rent
	rnt.UUID = uuid.New()
	rnt.RentRequestId = rr.UUID
	rnt.RentRequest = *rr
	rnt.Active = true
	rnt.Status = OnGoing
	rnt.OwnerId = rr.OwnerId
	rnt.RequesterId = rr.RequesterId
	return rnt
}

type Rent struct {
	entities.Model

	RequesterId uuid.UUID  `json:"requester_id" gorm:"type:char(36);index;not null"`
	Requester   users.User `json:"requester" gorm:"foreignKey:RequesterId;references:UUID"`

	OwnerId uuid.UUID  `json:"owner_id" gorm:"type:char(36);index;not null"`
	Owner   users.User `json:"owner" gorm:"foreignKey:OwnerId;references:UUID"`

	RentRequestId uuid.UUID   `json:"rent_request_id" gorm:"type:char(36);index;not null"`
	RentRequest   RentRequest `json:"rent_request" gorm:"foreignKey:RentRequestId;references:UUID"`
	Active        bool        `json:"active" gorm:"default:true"`
	Status        RentStatus  `json:"status" gorm:"not null,index,type:rent_status"`
}
