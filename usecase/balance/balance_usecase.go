package balance

import (
	"fmt"
	"strconv"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"

	"github.com/google/uuid"
)

type BalanceUsecase interface {
	GenerateVaUseCase(userId string, payload model.GenerateVirtualAgregator) (*model.VaNumber, error)
	CreateBalanceUseCase(payload *model.PartnerCallbackVirtualAggregator) (*model.Transaction, error)
	GetPayBalanceStatusUseCase(userId, vaId string) (*model.VaNumber, error)
}

type balanceUsecase struct {
	balanceRepository     repository.BalanceRepository
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
	virtualAgregatorOyApi repository.VirtualAgregatorOyApi
}

func NewBalanceUsecase(balanceRepository repository.BalanceRepository, userRepository repository.UserRepository, transactionRepository repository.TransactionRepository, virtualAgregatorOyApi repository.VirtualAgregatorOyApi) *balanceUsecase {
	return &balanceUsecase{
		balanceRepository:     balanceRepository,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
		virtualAgregatorOyApi: virtualAgregatorOyApi,
	}
}

func (uc *balanceUsecase) GenerateVaUseCase(userId string, payload model.GenerateVirtualAgregator) (*model.VaNumber, error) {
	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, err
	}

	payload.PatnerUserId = userId
	payload.Username = user.Name
	payload.IsOpen = true
	payload.SingleUse = false

	resp, err := uc.virtualAgregatorOyApi.GenerateVaApi(payload)
	if err != nil {
		return nil, err
	}

	insert := &model.VaNumber{
		UserId:         user.ID,
		VaNumber:       resp.VaNumber,
		VaStatus:       resp.VaStatus,
		Amount:         resp.Amount,
		ExpirationTime: resp.ExpirationTime,
		Name:           resp.Name,
	}

	existingVa, err := uc.balanceRepository.InsertVaRepository(insert)
	if err != nil {
		return nil, err
	}

	if existingVa != nil {
		return &model.VaNumber{
			UserId:         user.ID,
			VaNumber:       existingVa.VaNumber,
			VaStatus:       existingVa.VaStatus,
			BankCode:       payload.BankCode,
			Amount:         existingVa.Amount,
			ExpirationTime: existingVa.ExpirationTime,
			Name:           existingVa.Name,
		}, nil
	}

	return resp, nil
}

func (uc *balanceUsecase) CreateBalanceUseCase(payload *model.PartnerCallbackVirtualAggregator) (*model.Transaction, error) {
	userId := payload.PartnerUserId

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	user.Amount += payload.Amount
	user.UpdatedAt = time.Now()

	_, err = uc.userRepository.UpdateUserAmountByIDRepository(userId, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user amount: %v", err)
	}

	productDetail := &model.VaNumber{
		UserId:   userId,
		VaNumber: payload.VaNumber,
		Amount:   payload.Amount,
		Name:     payload.UsernameDisplay,
	}

	total := payload.Amount + model.ADMIN_FEE
	createTransaction := &model.Transaction{
		ID:            fmt.Sprintf("TOPUP-%s", uuid.New().String()),
		UserID:        user.ID,
		Description:   fmt.Sprintf("Top Up Saldo Skuypay : RP. %s", strconv.FormatFloat(payload.Amount, 'f', -1, 64)),
		ProductType:   "topup",
		Status:        model.STATUS_SUCCESSFUL,
		AdminFee:      model.ADMIN_FEE,
		Price:         payload.Amount,
		TotalPrice:    total,
		ProductDetail: productDetail,
	}

	resp, err := uc.transactionRepository.CreateTransactionByUserIdRepository(createTransaction)
	if err != nil {
		return nil, err
	}
	// mailsend := model.PayloadMail{
	// 	OrderId:       resp.ID,
	// 	CustomerName:  user.Name,
	// 	Status:        resp.Status,
	// 	RecipentEmail: user.Email,
	// 	TransactionAt: resp.UpdatedAt,
	// }
	// mail.SendingMail(mailsend)

	return resp, nil
}

func (uc *balanceUsecase) GetPayBalanceStatusUseCase(userId, vaId string) (*model.VaNumber, error) {

	resp, err := uc.virtualAgregatorOyApi.GetVaIdStatusVaApi(vaId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
