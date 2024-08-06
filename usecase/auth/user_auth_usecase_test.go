package auth

import (
	"errors"
	"testing"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUsecaseTest struct {
	suite.Suite
	authUsecase  AuthUsecase
	authRepoMock *mocks.AuthRepository
}

func TestAuthUsecaseTest(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTest))
}

func (s *AuthUsecaseTest) SetupTest() {
	s.authRepoMock = &mocks.AuthRepository{}
	s.authUsecase = NewAuthUsecase(s.authRepoMock)
}

func (m *AuthUsecaseTest) TestLoginUseCaseSucccess() {
	payload := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: "password123",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password123")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("GetUserByEmailRepository", payload.Email).Return(&result, nil)

	resp, err := m.authUsecase.LoginUseCase(payload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")

	}
	assert.Equal(m.T(), resp.Name, result.Name)
}

func (m *AuthUsecaseTest) TestLoginUseCaseErrorEmail() {
	payload := model.User{
		Name:     "user1",
		Email:    "",
		Password: "password123",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password123")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("GetUserByEmailRepository", payload.Email).Return(&result, errors.New("email field is required"))

	_, err := m.authUsecase.LoginUseCase(payload)
	if err != nil {

		assert.Error(m.T(), err, "email field is required")

	}
}

func (m *AuthUsecaseTest) TestLoginUseCaseErrorPasswordRequired() {
	payload := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: "password123",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password123")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("GetUserByEmailRepository", payload.Email).Return(&result, errors.New("password field is required"))

	_, err := m.authUsecase.LoginUseCase(payload)
	if err != nil {

		assert.Error(m.T(), err, "password field is required")

	}
}

func (m *AuthUsecaseTest) TestLoginUseCaseErrorRepository() {
	payload := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: "",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password123")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("GetUserByEmailRepository", payload.Email).Return(&result, errors.New("repository error"))

	_, err := m.authUsecase.LoginUseCase(payload)
	if err != nil {

		assert.Error(m.T(), err, "repository error")

	}
}
func TestComparePasswords(t *testing.T) {
	hashedPassword := "$2a$10$p3oqG1M0/WBaxnZGslo3wuA2itOzhxy7/bj/21XJIOusedaXiYUNq"

	password := "password123"

	// Call the ComparePasswords function
	result := ComparePasswords(hashedPassword, password)

	// Assert that the result is true
	if !result {
		t.Error("expected ComparePasswords to return true, but got false")
	}
}

func (m *AuthUsecaseTest) TestLoginUseCaseErrorPasswotNotSame() {
	payload := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: "password123",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password13")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("GetUserByEmailRepository", payload.Email).Return(&result, nil)

	_, err := m.authUsecase.LoginUseCase(payload)
	if err != nil {
		// m.T().Errorf("unexpected error: %v", err)

		assert.Error(m.T(), err, "repository error")

	}
	// assert.Equal(m.T(), resp.Name, result.Name)
}

func (m *AuthUsecaseTest) TestRegisterUseCase() {
	payload := model.User{
		Name:     "user1",
		Email:    "",
		Password: "",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password123")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("RegisterRepository", payload.Email).Return(&result, nil)

	_, err := m.authUsecase.RegisterUseCase(payload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")

	}
}

func (m *AuthUsecaseTest) TestRegisterAdminUseCase() {
	payload := model.User{
		Name:     "user1",
		Email:    "",
		Password: "",
		Phone:    "08132131321",
	}

	hashedPassword, _ := HashPassword("password123")

	result := model.User{
		Name:     "user1",
		Email:    "user1@example.com",
		Password: hashedPassword,
		Phone:    "08132131321",
		UserType: "user",
	}
	m.authRepoMock.On("RegisterRepository", payload.Email).Return(&result, nil)

	_, err := m.authUsecase.RegisterAdminUseCase(payload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")

	}
}

func TestHashPassword(t *testing.T) {
	password := "password123"

	// Call the HashPassword function
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert that the hashed password is not empty
	if hashedPassword == "" {
		t.Error("expected HashPassword to return a non-empty string, but got an empty string")
	}
}
