package wifi

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"
)

type WifiUsecase interface {
	CreateWifiUseCase(wifi *model.Wifi) (*model.Wifi, error)
	GetAllWifiUseCase(page, limit int) ([]*model.Wifi, error)
	GetWifiByIDUseCase(wifiID string) (*model.Wifi, error)
	GetWifiByCodeUseCase(wifiCode string) (*model.Wifi, error)
	UpdateWifiByIDUseCase(wifiID string, payload *model.Wifi) (*model.Wifi, error)
	DeleteWifiByIDUseCase(wifiID string) error
	BillInquiryWifiUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	PayBillWifiUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	BillWifiStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error)
}

type wifiUsecase struct {
	wifiRepository        repository.WifiRepository
	userRepository        repository.UserRepository
	discountRepository    repository.DiscountRepository
	transactionRepository repository.TransactionRepository
	billerOyApi           repository.BillerOyApiRepository
}

func NewWifiUseCase(wifiRepository repository.WifiRepository, userRepository repository.UserRepository, discountRepository repository.DiscountRepository, transactionRepository repository.TransactionRepository, billerOyApiRepository repository.BillerOyApiRepository) *wifiUsecase {
	return &wifiUsecase{
		wifiRepository:        wifiRepository,
		userRepository:        userRepository,
		discountRepository:    discountRepository,
		transactionRepository: transactionRepository,
		billerOyApi:           billerOyApiRepository,
	}
}

func (uc *wifiUsecase) CreateWifiUseCase(payload *model.Wifi) (*model.Wifi, error) {
	wifi, err := uc.wifiRepository.CreateWifiRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("error creating WiFi in database: %w", err)
	}
	return wifi, nil
}

func (uc *wifiUsecase) GetAllWifiUseCase(page, limit int) ([]*model.Wifi, error) {
	wifis, err := uc.wifiRepository.GetAllWifiRepository(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all wifi: %v", err)
	}
	return wifis, nil
}

func (uc *wifiUsecase) GetWifiByIDUseCase(wifiID string) (*model.Wifi, error) {
	wifi, err := uc.wifiRepository.GetWifiByIDRepository(wifiID)
	if err != nil {
		return nil, errors.New("wifi not found")
	}
	return wifi, nil
}

func (uc *wifiUsecase) GetWifiByCodeUseCase(wifiCode string) (*model.Wifi, error) {
	wifi, err := uc.wifiRepository.GetWifiByCodeRepository(wifiCode)
	if err != nil {
		return nil, errors.New("wifi not found")
	}
	return wifi, nil
}

func (uc *wifiUsecase) UpdateWifiByIDUseCase(wifiID string, payload *model.Wifi) (*model.Wifi, error) {
	wifi, err := uc.wifiRepository.GetWifiByIDRepository(wifiID)
	if err != nil {
		return nil, fmt.Errorf("failed to update wifi: %v", err)
	}

	wifi.ProductType = payload.ProductType
	wifi.Code = payload.Code
	wifi.Name = payload.Name
	wifi.UpdatedAt = time.Now()

	updatedWifi, err := uc.wifiRepository.UpdateWifiByIDRepository(wifiID, wifi)
	if err != nil {
		return nil, fmt.Errorf("failed to update wifi: %v", err)
	}

	return updatedWifi, nil
}

func (uc *wifiUsecase) DeleteWifiByIDUseCase(wifiID string) error {
	err := uc.wifiRepository.DeleteWifiByIDRepository(wifiID)
	if err != nil {
		return errors.New("wifi not found")
	}
	return nil
}

func (uc *wifiUsecase) BillInquiryWifiUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

	lastDigit := payload.CustomerId[len(payload.CustomerId)-1]
	if lastDigit == '9' {
		return nil, errors.New("invalid customer ID")
	}

	vaNumber := generateVANumber(16)
	payload.PartnerTxId = fmt.Sprintf("WIFI-%s", vaNumber)

	currentTime := time.Now()
	currentMonth := currentTime.Month().String()
	currentYear := strconv.Itoa(currentTime.Year())

	payload.Period = currentMonth + "-" + currentYear

	productype := strings.ToLower(payload.ProductId)
	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	existingWifi, err := uc.transactionRepository.GetProductDetailsByPeriodAndCustomerID(model.GetProductDetail{
		ProductId:  productype,
		Period:     payload.Period,
		CustomerId: payload.CustomerId,
	})
	if err == nil {
		if existingWifi.Status == model.STATUS_SUCCESSFUL {
			return nil, errors.New("this month's bill has been paid")
		}
		if existingWifi != nil {
			return existingWifi, nil
		}
	}

	discount, err := uc.discountRepository.GetDiscountByIdRepository(payload.DiscountId)
	if err != nil {
		return nil, errors.New("discount Not Found")
	}

	oy, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, err
	}

	bandwith := []int{10, 20, 30, 40, 50, 70, 80, 150}
	randomIndex := rand.Intn(len(bandwith))
	randomBandwith := bandwith[randomIndex]
	price := calculatePriceBandwith(randomBandwith)

	totalPrice := float64(price) + oy.AdminFee - float64(discount.DiscountPrice)

	productDetail := &model.Wifi{
		Name:         user.Name,
		CustomerID:   payload.CustomerId,
		Code:         oy.Code,
		ProviderName: oy.ProductID,
		ProductType:  productype,
		Period:       payload.Period,
		WifiBandwith: randomBandwith,
		DiscountId:   discount.ID,
		Price:        float64(price),
	}

	transaction := &model.Transaction{
		ID:            oy.PartnerTxID,
		UserID:        userId,
		Status:        model.STATUS_UNPAID,
		ProductType:   productype,
		DiscountPrice: float64(discount.DiscountPrice),
		AdminFee:      model.ADMIN_FEE,
		Description:   fmt.Sprintf("Pembayaran Tagihan WIFI %s ", payload.Period),
		Price:         float64(price),
		TotalPrice:    float64(totalPrice),
		ProductDetail: productDetail,
	}

	_, err = uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating WiFi in database: %w", err)
	}

	return transaction, nil
}

func (uc *wifiUsecase) PayBillWifiUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {
	transaction, err := uc.transactionRepository.GetTransactionByIdRepository(payload.PartnerTxId)
	if err != nil {
		return nil, err
	} else if transaction.Status == model.STATUS_SUCCESSFUL {
		return nil, errors.New("this WiFi bill has been paid")
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
	// 	RecipentEmail: user.Email,
	// 	TransactionAt: resp.UpdatedAt,
	// 	ProductType:   "WIFI",
	// 	Description:   transaction.Description,
	// 	AdminFee:      transaction.AdminFee,
	// 	Price:         transaction.Price,
	// 	TotalPrice:    transaction.TotalPrice,
	// }
	// mail.SendingMail(mailsend)

	return transactionresp, nil
}

func (uc *wifiUsecase) BillWifiStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {
	wifi, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve WiFi: %v", err)
	}
	return wifi, nil
}

func generateVANumber(length int) string {
	charset := "0123456789"
	rand.Seed(time.Now().UnixNano())

	vaNumber := make([]byte, length)
	for i := 0; i < length; i++ {
		vaNumber[i] = charset[rand.Intn(len(charset))]
	}

	return string(vaNumber)
}

func calculatePriceBandwith(bandwidth int) float64 {
	var totalPrice float64

	switch {
	case bandwidth <= 20:
		totalPrice = 275000.0
	case bandwidth <= 30:
		totalPrice = 315000.0
	case bandwidth <= 50:
		totalPrice = 445000.0
	default:
		totalPrice = 795000.0
	}

	return totalPrice
}
