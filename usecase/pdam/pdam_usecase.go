package pdam

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"
)

type PdamUseCase interface {
	CreatePdamUseCase(payload *model.Pdam) (*model.Pdam, error)
	GetAllPdamUseCase(page, limit int) ([]*model.Pdam, error)
	GetPdamByIdUseCase(pdamId string) (*model.Pdam, error)
	UpdatePdamByIdUseCase(pdamId string, payload *model.Pdam) (*model.Pdam, error)
	DeletePdamByIDUseCase(userId string) error
	BillInquiryPdamUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	PayBillPdamUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	BillPdamStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error)
}

type pdamUseCase struct {
	pdamRepository        repository.PdamRepository
	userRepository        repository.UserRepository
	discountRepository    repository.DiscountRepository
	transactionRepository repository.TransactionRepository
	billerOyApi           repository.BillerOyApiRepository
}

func NewPdamUseCase(pdamRepository repository.PdamRepository, userRepository repository.UserRepository, discountRepository repository.DiscountRepository, transactionRepository repository.TransactionRepository, billerOyApiRepository repository.BillerOyApiRepository) *pdamUseCase {
	return &pdamUseCase{pdamRepository: pdamRepository, userRepository: userRepository, discountRepository: discountRepository, transactionRepository: transactionRepository, billerOyApi: billerOyApiRepository}
}

func (uc *pdamUseCase) CreatePdamUseCase(payload *model.Pdam) (*model.Pdam, error) {

	pdam, err := uc.pdamRepository.CreatePdamRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("error creating PDAM in database: %w", err)
	}
	return pdam, err

}

func (uc *pdamUseCase) GetAllPdamUseCase(page, limit int) ([]*model.Pdam, error) {
	pdam, err := uc.pdamRepository.GetAllPdamRepository(page, limit)

	if err != nil {
		return nil, err
	}

	return pdam, nil
}

func (uc *pdamUseCase) GetPdamByIdUseCase(pdamId string) (*model.Pdam, error) {

	pdam, err := uc.pdamRepository.GetPdamByIdRepository(pdamId)
	if err != nil {
		return nil, errors.New("PDAM not found")
	}

	return pdam, nil

}

func (uc *pdamUseCase) UpdatePdamByIdUseCase(pdamId string, payload *model.Pdam) (*model.Pdam, error) {
	pdam, err := uc.pdamRepository.GetPdamByIdRepository(pdamId)
	if err != nil {
		return nil, fmt.Errorf("failed to update PDAM: %v", err)
	}

	pdam.ProviderName = payload.ProviderName
	pdam.Address = payload.Address
	pdam.Type = payload.Type
	pdam.UpdatedAt = time.Now()

	updatedPdam, err := uc.pdamRepository.UpdatePdamByIdRepository(pdamId, pdam)
	if err != nil {
		return nil, fmt.Errorf("failed to update PDAM: %v", err)
	}

	return updatedPdam, nil
}

func (uc *pdamUseCase) DeletePdamByIDUseCase(userId string) error {
	err := uc.pdamRepository.DeletePdamByIdRepository(userId)
	if err != nil {
		return errors.New("PDAM not found")
	}
	return err
}

func (uc *pdamUseCase) BillInquiryPdamUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

	lastDigit := payload.CustomerId[len(payload.CustomerId)-1]
	if lastDigit == '9' {
		return nil, errors.New("invalid customer ID")
	}

	vaNumber := generateVANumber(16)
	payload.PartnerTxId = fmt.Sprintf("PDAM-%s", vaNumber)

	currentTime := time.Now()
	currentMonth := currentTime.Month().String()
	currentYear := strconv.Itoa(currentTime.Year())

	payload.Period = currentMonth + "-" + currentYear
	productype := strings.ToLower(payload.ProductId)

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	existingPdam, err := uc.transactionRepository.GetProductDetailsByPeriodAndCustomerID(model.GetProductDetail{
		ProductId:  productype,
		Period:     payload.Period,
		CustomerId: payload.CustomerId,
	})
	if err == nil {
		if existingPdam.Status == model.STATUS_SUCCESSFUL {
			return nil, errors.New("this month's bill has been paid")
		}
		if existingPdam != nil {
			return existingPdam, nil
		}
	}

	discount, _ := uc.discountRepository.GetDiscountByIdRepository(payload.DiscountId)

	pdam, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, err
	}

	min := 1
	max := 50
	amount := rand.Intn(max-min+1) + min

	price := calculatePDAMBill(amount)

	totalPrice := price + pdam.AdminFee - float64(discount.DiscountPrice)

	productDetail := &model.Pdam{
		Period:       payload.Period,
		CustomerID:   pdam.CustomerID,
		ProviderName: pdam.ProductID,
		Name:         user.Name,
		Type:         productype,
		DiscountId:   discount.ID,
		Address:      user.Address,
		Price:        float64(price),
	}
	transaction := &model.Transaction{
		ID:            pdam.PartnerTxID,
		UserID:        userId,
		Status:        model.STATUS_UNPAID,
		Description:   "Pembayaran Tagihan PDAM",
		ProductType:   productype,
		DiscountPrice: float64(discount.DiscountPrice),
		AdminFee:      model.ADMIN_FEE,
		Price:         float64(price),
		TotalPrice:    float64(totalPrice),
		ProductDetail: productDetail,
	}

	_, err = uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating PDAM in database: %w", err)
	}
	return transaction, nil
}

func (uc *pdamUseCase) PayBillPdamUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

	transaction, err := uc.transactionRepository.GetTransactionByIdRepository(payload.PartnerTxId)
	if err != nil {
		return nil, err
	} else if transaction.Status == model.STATUS_SUCCESSFUL {
		return nil, errors.New("this month's bill has been paid")
	}

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, err
	}

	if user.Amount < transaction.TotalPrice {
		transactionFail := &model.Transaction{
			Status:    model.STATUS_FAIL,
			UpdatedAt: time.Now(),
		}

		_, err := uc.transactionRepository.UpdateTransactionByIdRepository(payload.PartnerTxId, transactionFail)
		if err != nil {
			return nil, errors.New("your balance is not enough")
		}

		return nil, errors.New("your balance is not enough")
	}

	user.Amount -= transaction.TotalPrice

	_, err = uc.userRepository.UpdateUserAmountByIDRepository(userId, user)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(transaction.ProductDetail)
	if err != nil {
		return nil, fmt.Errorf("error serializing transaction response to JSON: %w", err)
	}

	var pdam model.Pdam
	err = json.Unmarshal([]byte(jsonData), &pdam)
	if err != nil {
		return nil, fmt.Errorf("error Unmarshal: %w", err)
	}

	updateTransaction := &model.Transaction{

		Status:    model.STATUS_SUCCESSFUL,
		UpdatedAt: time.Now(),
	}

	resp, err := uc.transactionRepository.UpdateTransactionByIdRepository(payload.PartnerTxId, updateTransaction)
	if err != nil {
		return nil, fmt.Errorf("error Updating Transactions in database: %w", err)
	}

	transactionresp := &model.Transaction{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		Status:        resp.Status,
		ProductType:   transaction.ProductType,
		Description:   transaction.Description,
		DiscountPrice: transaction.DiscountPrice,
		AdminFee:      transaction.AdminFee,
		Price:         transaction.Price,
		TotalPrice:    transaction.TotalPrice,
		ProductDetail: transaction.ProductDetail,
	}

	// mailsend := model.PayloadMail{
	// 	OrderId:       transaction.ID,
	// 	CustomerName:  user.Name,
	// 	Status:        resp.Status,
	// 	Period:        pdam.Period,
	// 	RecipentEmail: user.Email,
	// 	ProductType:   "PDAM",
	// 	TransactionAt: resp.UpdatedAt,
	// 	Description:   transaction.Description,
	// 	AdminFee:      transaction.AdminFee,
	// 	Price:         transaction.Price,
	// 	TotalPrice:    transaction.TotalPrice,
	// }
	// mail.SendingMail(mailsend)

	return transactionresp, nil
}

func (uc *pdamUseCase) BillPdamStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {

	pdam, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve PDAM: %v", err)
	}

	return pdam, nil
}

func generateVANumber(length int) string {
	charset := "0123456789"
	rand.Seed(time.Now().Unix())

	vaNumber := make([]byte, length)
	for i := 0; i < length; i++ {
		vaNumber[i] = charset[rand.Intn(len(charset))]
	}

	return string(vaNumber)
}

func calculatePDAMBill(usage int) float64 {
	const (
		tariffPerM3 = 7450.0
		fixedFee    = 7450.0
		ppn         = 1195.0
	)

	waterUsage := float64(usage)
	usageCost := waterUsage * tariffPerM3
	totalBill := usageCost + fixedFee + ppn

	return totalBill
}
