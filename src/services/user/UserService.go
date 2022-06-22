package user

import (
	"github.com/google/uuid"
	"rent/src/entities/users"
	"rent/src/utils/errors"
	"rent/src/utils/hash"
)

type DefaultUserService struct {
	repo users.Repository // dependent on the port not the implementation (Adapter)
}

func (s DefaultUserService) RegisterUser(dto RegisterDto) (*LoggedUser, *errors.RestError) {
	err := dto.Validate()
	if err != nil {
		return nil, err
	}

	var user users.User
	user.FirstName = dto.FirstName
	user.Lastname = dto.Lastname
	user.Email = dto.Email
	user.PhoneNo = dto.PhoneNo
	user.Age = dto.Age
	user.Password, _ = hash.HashPassword(dto.Password)
	user.Status = users.Active
	user.UUID = uuid.New()

	err = s.repo.AddRegularUser(&user)
	if err != nil {
		return nil, err
	}

	token, err := user.CreateClaim()
	if err != nil {
		return nil, err
	}

	return &LoggedUser{
		User:  user,
		Token: token,
	}, nil
}

func (s DefaultUserService) LoginUser(dto LoginDto) (*LoggedUser, *errors.RestError) {
	err := dto.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetRegularUser(dto.Email)
	if err != nil {
		return nil, err
	}

	if hash.CheckPasswordHash(dto.Password, user.Password) {
		token, err := user.CreateClaim()
		if err != nil {
			return nil, err
		}
		return &LoggedUser{
			User:  *user,
			Token: token,
		}, nil
	} else {
		return nil, errors.NewBadRequestError("Email and Password dont match", "")
	}
}

func NewDefaultUserService(repo users.Repository) UserService {
	return DefaultUserService{
		repo: repo,
	}
}
