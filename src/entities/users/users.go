package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"rent/src/entities"
	"rent/src/logger"
	"rent/src/utils/errors"
)

type BaseUser struct {
	entities.Model
	FirstName string  `json:"first_name" gorm:"not null" faker:"first_name"`
	Lastname  string  `json:"last_name" gorm:"not null" faker:"last_name"`
	Email     string  `json:"email" gorm:"unique;index;not null" faker:"email,unique"`
	PhoneNo   *string `json:"phone_no" gorm:"unique" faker:"e_164_phone_number,unique"`
	Password  string  `json:"-" gorm:"not null"`
	Age       *int    `json:"age"`
}

type RenterProfile struct {
	entities.Model
	FamilySize   int    `json:"family_size" gorm:"not null"`
	PetSize      int    `json:"pet_size"`
	Vehicles     int    `json:"vehicles"`
	SpecialNeeds string `json:"special_needs"`
}

type User struct {
	BaseUser
	Status          UserStatus     `json:"status" gorm:"not null,type:user_status"`
	RenterProfileId *uuid.UUID     `json:"renter_profile_id" gorm:"type:char(36);index"`
	RenterProfile   *RenterProfile `json:"renter_profile" gorm:"foreignKey:RenterProfileId;references:UUID"`
	RatingAsOwner   *int           `json:"rating_as_owner" gorm:"default:0"`
	RatingAsRenter  *int           `json:"rating_as_renter" gorm:"default:0"`
}

func (u User) CreateClaim() (string, *errors.RestError) {
	payload := jwt.StandardClaims{
		// TODO Add expire at time for phoneNo code claim
		//ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:      u.UUID.String(),
		Subject: "Regular User",
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		logger.Error("Error happened when creating claim for User " + err.Error())
		return "", errors.NewInternalServerError("Internal server error happened", "")
	}
	return token, nil
}

type AdminUser struct {
	BaseUser
	Status UserStatus `json:"status" gorm:"type:user_status"`
	Roles  AdminRoles `json:"roles" gorm:"type:admin_roles[]"`
}

func (u AdminUser) CreateClaim() (string, *errors.RestError) {
	payload := jwt.StandardClaims{
		// TODO Add expire at time for phoneNo code claim
		//ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:      u.UUID.String(),
		Subject: "Admin User",
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		logger.Error("Error happened when creating claim for User " + err.Error())
		return "", errors.NewInternalServerError("Internal server error happened", "")
	}
	return token, nil
}
