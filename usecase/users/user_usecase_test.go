package users_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/darulfh/skuy_pay_be/dto"
	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository/mocks"
	"github.com/darulfh/skuy_pay_be/usecase/auth"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"
	"github.com/darulfh/skuy_pay_be/usecase/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserUseCaseTest struct {
	suite.Suite
	midleware           middlewares.Middlewares
	userUseCase         users.UserUsecase
	userRepoMock        *mocks.UserRepository
	transactionRepoMock *mocks.TransactionRepository
}

func TestUserUseCaseTest(t *testing.T) {
	suite.Run(t, new(UserUseCaseTest))
}

func (s *UserUseCaseTest) SetupTest() {
	s.userRepoMock = &mocks.UserRepository{}
	s.transactionRepoMock = &mocks.TransactionRepository{}
	s.userUseCase = users.NewUserUsecase(s.userRepoMock, s.transactionRepoMock)
}

func (s *UserUseCaseTest) TestGetAllUsers() {
	mockUsers := []*model.User{
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "1",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:    "arby",
			Email:   "arby@mail.com",
			Address: "Jl",
			Phone:   "08123456789",
		},
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "2",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:    "User2",
			Email:   "user2@example.com",
			Address: "Jl. Xyz No. 456",
			Phone:   "08123456788",
		},
	}

	s.userRepoMock.On("GetAllUsersRepository", 1, 10, "arby").Return(mockUsers, nil)

	resp, err := s.userUseCase.GetAllUsersUseCase(1, 10, "arby")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "arby", resp[0].Name)
	assert.Equal(s.T(), "arby@mail.com", resp[0].Email)
	assert.Equal(s.T(), "08123456789", resp[0].Phone)
	assert.Equal(s.T(), mockUsers[0].CreatedAt.Unix(), resp[0].CreatedAt.Unix())
	assert.Equal(s.T(), mockUsers[0].UpdatedAt.Unix(), resp[0].UpdatedAt.Unix())
}

func (s *UserUseCaseTest) TestGetAllUsers_Error() {
	s.userRepoMock.On("GetAllUsersRepository", 1, 10, "arby").Return(nil, errors.New("repository error"))

	resp, err := s.userUseCase.GetAllUsersUseCase(1, 10, "arby")
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), "failed to get all users: repository error", err.Error())
}

func (s *UserUseCaseTest) TestGetUserByID() {
	userID := "1"
	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:    "arby",
		Email:   "arby@mail.com",
		Address: "Jl",
		Phone:   "08123456789",
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)

	resp, err := s.userUseCase.GetUserByIDUseCase(userID)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), userID, resp.ID)
	assert.Equal(s.T(), "arby", resp.Name)
	assert.Equal(s.T(), "arby@mail.com", resp.Email)
	assert.Equal(s.T(), "08123456789", resp.Phone)
	assert.Equal(s.T(), user.CreatedAt.Unix(), resp.CreatedAt.Unix())
	assert.Equal(s.T(), user.UpdatedAt.Unix(), resp.UpdatedAt.Unix())
}

func (s *UserUseCaseTest) TestGetUserByEmail() {
	email := "arby@mail.com"
	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:   "arby",
		Email:  "arby@mail.com",
		Phone:  "08123456789",
		Amount: 1000,
		Image:  "image.jpg",
	}

	s.userRepoMock.On("GetUserByEmail", email).Return(user, nil)

	resp, err := s.userUseCase.GetUserByEmailUsecase(email)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), user.ID, resp.ID)
	assert.Equal(s.T(), user.Name, resp.Name)
	assert.Equal(s.T(), user.Email, resp.Email)
	assert.Equal(s.T(), user.Phone, resp.Phone)
	assert.Equal(s.T(), user.Amount, resp.Balance)
	assert.Equal(s.T(), user.Image, resp.Image)
	assert.Equal(s.T(), user.UpdatedAt, resp.UpdatedAt)
	assert.Equal(s.T(), user.CreatedAt, resp.CreatedAt)
}

func (s *UserUseCaseTest) TestUpdateUserImageByIDUseCase() {
	userID := "1"
	payload := &model.User{
		Image: "new_image.jpg",
	}
	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:   "arby",
		Email:  "arby@mail.com",
		Phone:  "08123456789",
		Amount: 1000,
		Image:  "image.jpg",
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.AnythingOfType("*model.User")).Return(user, nil)

	resp, err := s.userUseCase.UpdateUserImageByIDUseCase(userID, payload)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), user.ID, resp.ID)
	assert.Equal(s.T(), payload.Image, resp.Image)
	assert.Equal(s.T(), user.UpdatedAt, resp.UpdatedAt)
}

func (s *UserUseCaseTest) TestUpdateUserByIDUseCase() {
	userID := "1"
	payload := &model.User{
		Name:    "usman",
		Email:   "usman@example.com",
		Address: "123 Street",
		Phone:   "08123456789",
	}
	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:    "arby",
		Email:   "arby@mail.com",
		Address: "Jl. Abc No. 123",
		Phone:   "08123456789",
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.AnythingOfType("*model.User")).Return(user, nil)

	resp, err := s.userUseCase.UpdateUserByIDUseCase(userID, payload)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), user.ID, resp.ID)
	assert.Equal(s.T(), strings.ToLower(payload.Name), resp.Name)
	assert.Equal(s.T(), strings.ToLower(payload.Email), resp.Email)
	assert.Equal(s.T(), strings.ToLower(payload.Address), resp.Address)
	assert.Equal(s.T(), payload.Phone, resp.Phone)
	assert.Equal(s.T(), user.UpdatedAt, resp.UpdatedAt)
}

func (s *UserUseCaseTest) TestDeleteUserByIDUseCase() {
	userID := "1"

	s.userRepoMock.On("DeleteUserByIDRepository", userID).Return(nil)

	err := s.userUseCase.DeleteUserByIDUseCase(userID)

	assert.NoError(s.T(), err)
}

func (s *UserUseCaseTest) TestChangePasswordUseCase() {
	userID := "1"
	currentPassword := "current_password"
	newPassword := "new_password"
	hashedCurrentPassword, _ := auth.HashPassword(currentPassword)
	hashedNewPassword, _ := auth.HashPassword(newPassword)

	payload := dto.Password{
		Password:    currentPassword,
		Newpassword: newPassword,
	}

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "arby",
		Email:    "arby@mail.com",
		Password: hashedCurrentPassword,
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.AnythingOfType("*model.User")).Return(user, nil)

	err := s.userUseCase.ChangePasswordUseCase(userID, payload)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), hashedNewPassword, hashedNewPassword)
	assert.WithinDuration(s.T(), time.Now(), user.UpdatedAt, time.Second)
}
func (s *UserUseCaseTest) TestCreatePINUseCase() {
	userID := "1"
	pin := "1234"
	hashedPin := users.HashPin(pin)

	payload := dto.PIN{
		Pin:    pin,
		NewPIN: pin,
	}

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "arby",
		Email: "arby@mail.com",
		Pin:   "",
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.AnythingOfType("*model.User")).Return(user, nil)

	err := s.userUseCase.CreatePINUseCase(userID, payload)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), hashedPin, hashedPin)
	assert.WithinDuration(s.T(), time.Now(), user.UpdatedAt, time.Second)
}

func (s *UserUseCaseTest) TestGetUserByQueryUseCase() {
	query := "arby"
	page := 1
	limit := 10

	mockUsers := []*model.User{
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "1",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:    "arby",
			Email:   "arby@mail.com",
			Address: "Jl. Abc No. 123",
			Phone:   "08123456789",
		},
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "2",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:    "User2",
			Email:   "user2@example.com",
			Address: "Jl. Xyz No. 456",
			Phone:   "08123456788",
		},
	}

	s.userRepoMock.On("GetUserByQueryRepository", query, page, limit).Return(mockUsers, nil)

	users, err := s.userUseCase.GetUserByQueryUseCase(query, page, limit)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), users)
	assert.Equal(s.T(), len(mockUsers), len(users))
	assert.Equal(s.T(), mockUsers[0].ID, users[0].ID)
	assert.Equal(s.T(), mockUsers[0].Name, users[0].Name)
	assert.Equal(s.T(), mockUsers[0].Email, users[0].Email)
	assert.Equal(s.T(), mockUsers[0].Address, users[0].Address)
	assert.Equal(s.T(), mockUsers[0].Phone, users[0].Phone)
	assert.Equal(s.T(), mockUsers[0].CreatedAt.Unix(), users[0].CreatedAt.Unix())
	assert.Equal(s.T(), mockUsers[0].UpdatedAt.Unix(), users[0].UpdatedAt.Unix())
}
func (s *UserUseCaseTest) TestUpdateForgotPasswordUseCase_Success() {
	userID := "2"
	password := "new_password"
	token := "valid_token"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "2",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:    "arby@gmail.com",
		Name:     "arbyusman",
		Password: "old_password",
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.Anything).Return(user, nil)

	response, err := s.userUseCase.UpdateForgotPasswordUsecase(userID, dto.Password{
		Newpassword: password,
	})

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), response)
	assert.Equal(s.T(), userID, response.ID)
	assert.Equal(s.T(), user.Email, response.Email)
	assert.Equal(s.T(), user.Name, response.Name)
	assert.Equal(s.T(), token, token)

	assert.Equal(s.T(), password, password)
	assert.True(s.T(), user.UpdatedAt.After(user.CreatedAt))

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
	s.userRepoMock.AssertCalled(s.T(), "UpdateUserByIDRepository", userID, user)
}

func (s *UserUseCaseTest) TestUpdateForgotPasswordUseCase_UserNotFound() {
	userID := "user_id"
	password := "new_password"

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(nil, errors.New("user not found"))

	response, err := s.userUseCase.UpdateForgotPasswordUsecase(userID, dto.Password{
		Newpassword: password,
	})

	assert.Error(s.T(), err)
	assert.Nil(s.T(), response)
	assert.EqualError(s.T(), err, "user not found")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

func (s *UserUseCaseTest) TestUpdateForgotPasswordUseCase_UpdateError() {
	userID := "2"
	password := "new_password"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "2",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:    "test@gmail.com",
		Name:     "arbyusman",
		Password: "old_password",
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.Anything).Return(nil, errors.New("update error"))

	response, err := s.userUseCase.UpdateForgotPasswordUsecase(userID, dto.Password{
		Newpassword: password,
	})

	assert.Error(s.T(), err)
	assert.Nil(s.T(), response)
	assert.EqualError(s.T(), err, "failed to update user password: update error")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
	s.userRepoMock.AssertCalled(s.T(), "UpdateUserByIDRepository", userID, user)
}

func (s *UserUseCaseTest) TestChangePINUseCase_Success() {
	userID := "user_id"
	pin := "1234"
	newPIN := "5678"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        userID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email: "arby@gmail.com",
		Name:  "arbyusman",
		Pin:   users.HashPin(pin),
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.Anything).Return(user, nil)

	err := s.userUseCase.ChangePINUseCase(userID, dto.PIN{
		Pin:    pin,
		NewPIN: newPIN,
	})

	assert.NoError(s.T(), err)
	pin = users.HashPin(newPIN)

	assert.Equal(s.T(), pin, pin)
	assert.True(s.T(), user.UpdatedAt.After(user.CreatedAt))

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
	s.userRepoMock.AssertCalled(s.T(), "UpdateUserByIDRepository", userID, user)
}

func (s *UserUseCaseTest) TestChangePINUseCase_UserNotFound() {
	userID := "user_id"
	pin := "1234"
	newPIN := "5678"

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(nil, errors.New("user not found"))

	err := s.userUseCase.ChangePINUseCase(userID, dto.PIN{
		Pin:    pin,
		NewPIN: newPIN,
	})

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "user not found")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

func (s *UserUseCaseTest) TestChangePINUseCase_IncorrectCurrentPIN() {
	userID := "user_id"
	pin := "1234"
	newPIN := "5678"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        userID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email: "arby@gmail.com",
		Name:  "arbyusman",
		Pin:   users.HashPin("4321"),
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)

	err := s.userUseCase.ChangePINUseCase(userID, dto.PIN{
		Pin:    pin,
		NewPIN: newPIN,
	})

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "current PIN is incorrect")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

func (s *UserUseCaseTest) TestChangePINUseCase_SameNewPINAsCurrent() {
	userID := "user_id"
	pin := "1234"
	newPIN := "1234"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        userID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},

		Email: "arby@gmail.com",
		Name:  "arbyusman",
		Pin:   users.HashPin(pin),
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)

	err := s.userUseCase.ChangePINUseCase(userID, dto.PIN{
		Pin:    pin,
		NewPIN: newPIN,
	})

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "new PIN must be different from the current PIN")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

func (s *UserUseCaseTest) TestChangePINUseCase_UpdateError() {
	userID := "user_id"
	pin := "1234"
	newPIN := "5678"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        userID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},

		Email: "arby@gmail.com",
		Name:  "arbyusman",
		Pin:   users.HashPin(pin),
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)
	s.userRepoMock.On("UpdateUserByIDRepository", userID, mock.Anything).Return(nil, errors.New("update error"))

	err := s.userUseCase.ChangePINUseCase(userID, dto.PIN{
		Pin:    pin,
		NewPIN: newPIN,
	})

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "failed to change PIN: update error")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
	s.userRepoMock.AssertCalled(s.T(), "UpdateUserByIDRepository", userID, user)
}

func (s *UserUseCaseTest) TestCheckPINUseCase_CorrectPIN() {
	userID := "user_id"
	pin := "1234"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "2",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},

		Email: "arby@gmail.com",
		Name:  "arbyusman",
		Pin:   users.HashPin(pin),
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)

	err := s.userUseCase.CheckPINUseCase(userID, dto.PIN{
		Pin: pin,
	})

	assert.NoError(s.T(), err)

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

func (s *UserUseCaseTest) TestCheckPINUseCase_IncorrectPIN() {
	userID := "user_id"
	pin := "1234"

	user := &model.User{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "2",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},

		Email: "arby@gmail.com",
		Name:  "arbyusman",
		Pin:   users.HashPin("4321"),
	}

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(user, nil)

	err := s.userUseCase.CheckPINUseCase(userID, dto.PIN{
		Pin: pin,
	})

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "incorrect")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

func (s *UserUseCaseTest) TestCheckPINUseCase_UserNotFound() {
	userID := "user_id"
	pin := "1234"

	s.userRepoMock.On("GetUserByIDRepository", userID).Return(nil, errors.New("user not found"))

	err := s.userUseCase.CheckPINUseCase(userID, dto.PIN{
		Pin: pin,
	})

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "user not found")

	s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
}

// func (s *UserUseCaseTest) TestTransferAmountUseCase() {
// 	userID := "user123"
// 	phoneNumber := "08123456789"
// 	amount := 100.0
// 	note := "Transfer test"

// 	mockUser := &model.User{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "2",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		Phone:  phoneNumber,
// 		Amount: 200.0,
// 	}

// 	mockTransaction := &model.Transaction{
// 		ID:          "transaction123",
// 		UserID:      userID,
// 		Status:      model.STATUS_SUCCESSFUL,
// 		ProductType: "transfer",
// 		TotalPrice:  amount,
// 		Price:       amount,
// 		AdminFee:    0,
// 		ProductDetail: model.TransactionTransfer{
// 			Phone:  phoneNumber,
// 			UserID: mockUser.ID,
// 			Note:   note,
// 		},
// 	}

// 	s.userRepoMock.On("GetUserByPhone", phoneNumber).Return(mockUser, nil)
// 	s.userRepoMock.On("GetUserByIDRepository", userID).Return(mockUser, nil)
// 	s.transactionRepoMock.On("CreateTransactionByUserIdRepository", mock.AnythingOfType("*model.Transaction")).Return(mockTransaction, nil)
// 	s.userRepoMock.On("UpdateUserAmountByIDRepository", userID, mock.AnythingOfType("*model.User")).Return(nil)

// 	resp, err := s.userUseCase.TransferAmountUseCase(userID, dto.TransactionTransferDto{
// 		PhoneNumber: phoneNumber,
// 		Amount:      amount,
// 		Note:        note,
// 	})

// 	assert.NoError(s.T(), err)
// 	assert.Equal(s.T(), mockTransaction, resp)

// 	// s.userRepoMock.AssertCalled(s.T(), "GetUserByPhone", phoneNumber)
// 	// s.userRepoMock.AssertCalled(s.T(), "GetUserByIDRepository", userID)
// 	// s.transactionRepoMock.AssertCalled(s.T(), "CreateTransactionByUserIdRepository", mock.AnythingOfType("*model.Transaction"))
// 	// s.userRepoMock.AssertCalled(s.T(), "UpdateUserAmountByIDRepository", userID, mock.AnythingOfType("*model.User"))
// }
