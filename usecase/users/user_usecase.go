package users

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/darulfh/skuy_pay_be/dto"
	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"
	"github.com/darulfh/skuy_pay_be/usecase/auth"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetAllUsersUseCase(page, limit int, name string) ([]*model.UserResponse, error)
	GetUserByIDUseCase(userId string) (*model.UserResponse, error)
	GetUserByEmailUsecase(email string) (*model.UserResponse, error)
	UpdateUserImageByIDUseCase(userId string, payload *model.User) (*model.UserResponse, error)
	UpdateUserByIDUseCase(userId string, payload *model.User) (*model.UserResponse, error)
	DeleteUserByIDUseCase(userId string) error
	ChangePasswordUseCase(userID string, payload dto.Password) error
	CreatePINUseCase(userID string, payload dto.PIN) error
	ChangePINUseCase(userID string, payload dto.PIN) error
	CheckPINUseCase(userID string, payload dto.PIN) error
	TransferAmountUseCase(userID string, payload dto.TransactionTransferDto) (*model.Transaction, error)
	UpdateForgotPasswordUsecase(userID string, payload dto.Password) (*model.AuthResponse, error)
	GetUserByQueryUseCase(query string, page, limit int) ([]*model.User, error)
}

type userUsecase struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewUserUsecase(userRepository repository.UserRepository, transactionRepository repository.TransactionRepository) *userUsecase {
	return &userUsecase{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

func (uc *userUsecase) GetAllUsersUseCase(page, limit int, name string) ([]*model.UserResponse, error) {

	users, err := uc.userRepository.GetAllUsersRepository(page, limit, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	resp := make([]*model.UserResponse, 0, len(users))
	for _, user := range users {
		resp = append(resp, &model.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Balance:   user.Amount,
			Phone:     user.Phone,
			Image:     user.Image,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return resp, nil
}

func (uc *userUsecase) GetUserByIDUseCase(userId string) (*model.UserResponse, error) {

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Pin:       user.Pin,
		Balance:   user.Amount,
		Image:     user.Image,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}

	return resp, nil
}

func (uc *userUsecase) GetUserByEmailUsecase(email string) (*model.UserResponse, error) {
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Balance:   user.Amount,
		Image:     user.Image,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}

	return resp, nil
}
func (uc *userUsecase) UpdateUserImageByIDUseCase(userId string, payload *model.User) (*model.UserResponse, error) {

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	user.Image = payload.Image
	user.UpdatedAt = time.Now()

	updatedUser, err := uc.userRepository.UpdateUserByIDRepository(userId, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	resp := &model.UserResponse{
		ID:        updatedUser.ID,
		Image:     updatedUser.Image,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return resp, nil
}

func (uc *userUsecase) UpdateUserByIDUseCase(userId string, payload *model.User) (*model.UserResponse, error) {

	lowercasePayload := &model.User{
		Name:    strings.ToLower(payload.Name),
		Email:   strings.ToLower(payload.Email),
		Address: strings.ToLower(payload.Address),
		Phone:   payload.Phone,
	}

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	user.Name = lowercasePayload.Name
	user.Email = lowercasePayload.Email
	user.Address = lowercasePayload.Address
	user.Phone = lowercasePayload.Phone
	user.UpdatedAt = time.Now()

	updatedUser, err := uc.userRepository.UpdateUserByIDRepository(userId, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	resp := &model.UserResponse{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Address:   updatedUser.Address,
		Phone:     updatedUser.Phone,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return resp, nil
}

func (uc *userUsecase) DeleteUserByIDUseCase(userId string) error {
	err := uc.userRepository.DeleteUserByIDRepository(userId)
	if err != nil {
		return errors.New("user not found")
	}
	return err
}

func (uc *userUsecase) ChangePasswordUseCase(userID string, payload dto.Password) error {
	user, err := uc.userRepository.GetUserByIDRepository(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if !auth.ComparePasswords(user.Password, payload.Password) {
		return errors.New("current password is incorrect")
	}

	if auth.ComparePasswords(user.Password, payload.Newpassword) {
		return errors.New("new Password must be different from the current Password")
	}

	hashedPassword, _ := auth.HashPassword(payload.Newpassword)

	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	_, err = uc.userRepository.UpdateUserByIDRepository(userID, user)
	if err != nil {
		return fmt.Errorf("failed to update user password: %v", err)
	}

	return nil
}

func (uc *userUsecase) CreatePINUseCase(userID string, payload dto.PIN) error {
	user, err := uc.userRepository.GetUserByIDRepository(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Pin != "" {
		return errors.New("PIN already exists for this user")
	}

	hashedPin := HashPin(payload.Pin)
	user.Pin = hashedPin
	user.UpdatedAt = time.Now()

	_, err = uc.userRepository.UpdateUserByIDRepository(userID, user)
	if err != nil {
		return fmt.Errorf("failed to create PIN: %v", err)
	}

	return nil
}

func (uc *userUsecase) ChangePINUseCase(userID string, payload dto.PIN) error {
	user, err := uc.userRepository.GetUserByIDRepository(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if !CompareHashPin(user.Pin, payload.Pin) {
		return errors.New("current PIN is incorrect")
	}
	if CompareHashPin(user.Pin, payload.NewPIN) {
		return errors.New("new PIN must be different from the current PIN")
	}

	hashedNewPIN := HashPin(payload.NewPIN)
	user.Pin = hashedNewPIN
	user.UpdatedAt = time.Now()

	_, err = uc.userRepository.UpdateUserByIDRepository(userID, user)
	if err != nil {
		return fmt.Errorf("failed to change PIN: %v", err)
	}

	return nil
}

func (uc *userUsecase) CheckPINUseCase(userID string, payload dto.PIN) error {
	user, err := uc.userRepository.GetUserByIDRepository(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if !CompareHashPin(user.Pin, payload.Pin) {
		return errors.New("incorrect")
	}

	return nil
}

func (uc *userUsecase) TransferAmountUseCase(userID string, payload dto.TransactionTransferDto) (*model.Transaction, error) {
	user, err := uc.userRepository.GetUserByPhone(payload.PhoneNumber)
	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error user id not found")
	}
	myprofile, err := uc.userRepository.GetUserByIDRepository(userID)
	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error user id not found")
	}
	if myprofile.Amount < payload.Amount {
		return &model.Transaction{}, fmt.Errorf("your balance is not enough")
	}
	td := model.TransactionTransfer{
		Phone: user.Phone, UserID: user.ID, Note: payload.Note,
	}
	tf := model.Transaction{
		ID:            uuid.New().String(),
		UserID:        userID,
		Status:        model.STATUS_SUCCESSFUL,
		ProductType:   "transfer",
		TotalPrice:    payload.Amount,
		Price:         payload.Amount,
		AdminFee:      0,
		ProductDetail: td,
	}
	response, err := uc.transactionRepository.CreateTransactionByUserIdRepository(&tf)
	if err != nil {
		return &model.Transaction{}, err
	}

	myprofile.Amount -= response.TotalPrice
	_, err = uc.userRepository.UpdateUserAmountByIDRepository(userID, myprofile)

	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error update amount")
	}

	user.Amount += response.Price
	_, err = uc.userRepository.UpdateUserAmountByIDRepository(user.ID, user)

	if err != nil {
		return &model.Transaction{}, fmt.Errorf("error update amount")
	}
	return response, nil
}

func (uc *userUsecase) UpdateForgotPasswordUsecase(userID string, payload dto.Password) (*model.AuthResponse, error) {
	user, err := uc.userRepository.GetUserByIDRepository(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	hashedPassword, _ := auth.HashPassword(payload.Newpassword)

	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	_, err = uc.userRepository.UpdateUserByIDRepository(userID, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user password: %v", err)
	}

	token, err := middlewares.CreateToken(*user)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	resp := &model.AuthResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		Token: token,
	}

	return resp, nil
}

func (uc *userUsecase) GetUserByQueryUseCase(query string, page, limit int) ([]*model.User, error) {
	users, err := uc.userRepository.GetUserByQueryRepository(query, page, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CompareHashPin(hashPin string, pin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPin), []byte(pin))
	return err == nil
}

func HashPin(pin string) string {
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPin)
}
