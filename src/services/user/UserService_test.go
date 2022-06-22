package user

import (
	"rent/src/adaptors/users"
	users2 "rent/src/entities/users"
	"rent/src/utils/hash"
	"testing"
)

func CreateTestUser(email, password string) *users2.User {
	age := 19
	phoneNo := "+251911111111"

	var user users2.User
	user.FirstName = "test"
	user.Lastname = "test"
	user.Email = email
	user.PhoneNo = &phoneNo
	user.Password, _ = hash.HashPassword(password)
	user.Age = &age
	user.Status = users2.Active
	return &user
}

func CreateCorrectRegisterDto() RegisterDto {
	age := 19
	phoneNo := "+251911111111"

	return RegisterDto{
		FirstName: "test_first_name",
		Lastname:  "test_last_name",
		Email:     "test@test.com",
		PhoneNo:   &phoneNo,
		Password:  "test",
		Age:       &age,
	}
}

func CreateInCorrectRegisterDto() RegisterDto {
	age := 10
	phoneNo := "test"

	return RegisterDto{
		FirstName: "test_first_name",
		Lastname:  "test_last_name",
		Email:     "test",
		PhoneNo:   &phoneNo,
		Password:  "test",
		Age:       &age,
	}

}

func TestRegisterDto_Validate_raise_error(t *testing.T) {
	dto := CreateInCorrectRegisterDto()
	err := dto.Validate()

	if err == nil {
		t.Errorf("Validator should have raised error")
	}
}

func TestRegisterDto_Validate_accept_correct_dto(t *testing.T) {
	dto := CreateCorrectRegisterDto()

	err := dto.Validate()

	if err != nil {
		t.Errorf("Validator should not raise error")
	}
}

func TestLoginDto_Validate_raise_error(t *testing.T) {
	dto := LoginDto{
		Email:    "",
		Password: "",
	}

	err := dto.Validate()
	if err == nil {
		t.Errorf("Validator should have raised error")
	}
}

func TestLoginDto_Validate_accept_correct_dto(t *testing.T) {
	dto := LoginDto{
		Email:    "test@test.com",
		Password: "1234543",
	}

	err := dto.Validate()
	if err != nil {
		t.Errorf("Validator should not raise error")
	}
}

func TestDefaultUserService_RegisterUser(t *testing.T) {
	user := CreateTestUser("test@test.com", "12345")
	dto := CreateCorrectRegisterDto()
	service := NewDefaultUserService(users.NewUserRepositoryStub(*user))
	res, err := service.RegisterUser(dto)

	if err != nil {
		t.Errorf("Registeration should have completed")
	}

	if res == nil {
		t.Errorf("response should have been created and returned")
	}

	if res.Token == "" {
		t.Errorf("Token should have been created and returned")
	}

	if res.User.Email != dto.Email {
		t.Errorf("User should have been created with the provided email")
	}

	if res.User.Password == dto.Password {
		t.Errorf("Password should have been hashed and saved")
	}

	if res.User.UUID.String() == "00000000-0000-0000-0000-000000000000" {
		t.Errorf("user should have been assigned a UUID")
	}

}

func TestDefaultUserService_LoginUser_with_correct_credentials(t *testing.T) {
	user := CreateTestUser("test@test.com", "12345")
	dto := LoginDto{
		Email:    "test@test.com",
		Password: "12345",
	}
	service := NewDefaultUserService(users.NewUserRepositoryStub(*user))
	res, err := service.LoginUser(dto)

	if err != nil {
		t.Errorf("Logging should have completed")
	}

	if res == nil {
		t.Errorf("response should have been created and returned")
	}

	if res.Token == "" {
		t.Errorf("Token should have been created and returned")
	}

	if res.User.Email != dto.Email {
		t.Errorf("Logged userController should have the same email as the one with the input")
	}
}

func TestDefaultUserService_LoginUser_with_incorrect_credentials(t *testing.T) {
	user := CreateTestUser("test@test.com", "12345")
	dto := LoginDto{
		Email:    "test@test.com",
		Password: "abcde",
	}
	service := NewDefaultUserService(users.NewUserRepositoryStub(*user))
	res, err := service.LoginUser(dto)

	if err == nil {
		t.Errorf("Logging should not complete")
	}

	if res != nil {
		t.Errorf("response should have not been created and returned")
	}

}
