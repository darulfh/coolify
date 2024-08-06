package electricity

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
	"github.com/google/uuid"
)

type ElectricityUseCase interface {
	CreateElectricityUseCase(payload *model.Electricity) (*model.Electricity, error)
	GetAllElectricityUseCase(page, limit int) ([]*model.Electricity, error)
	GetElectricityByIdUseCase(electricityId string) (*model.Electricity, error)
	UpdateElectricityByIdUseCase(electricityId string, payload *model.Electricity) (*model.Electricity, error)
	DeleteElectricityByIDUseCase(userId string) error
	PostBillInquiryElectricityUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	PostPayBillElectricityUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	PreBillInquiryElectricityUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	BillElectricityStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error)

	ElectricityBillInquiryIakUseCase(payload *model.IakInquiryBody) (*model.IakPostPaidResponse, error)
	ElectricityBillPayIakUseCase(payload *model.IakPayBody, userId string) (*model.IakPostPaidResponse, error)

	ElectricityTokenInquiryIakUseCase(payload *model.PrePaidIakBody) (*model.IakElectricityTokenInquiry, error)
	ElectricityTokenPayIakUseCase(payload *model.PrePaidIakBody, userId string) (*model.PrePaidIakResponse, error)
}

type electricityUseCase struct {
	electricityRepository repository.ElectricityRepository
	userRepository        repository.UserRepository
	discountRepository    repository.DiscountRepository
	transactionRepository repository.TransactionRepository
	billerOyApi           repository.BillerOyApiRepository
	iakRepository         repository.IakApiRepository
}

func NewElectricityUseCase(electricityRepository repository.ElectricityRepository, userRepository repository.UserRepository, discountRepository repository.DiscountRepository, transactionRepository repository.TransactionRepository, billerOyApiRepository repository.BillerOyApiRepository, iakRepository repository.IakApiRepository) *electricityUseCase {
	return &electricityUseCase{electricityRepository: electricityRepository, userRepository: userRepository, discountRepository: discountRepository, transactionRepository: transactionRepository, billerOyApi: billerOyApiRepository, iakRepository: iakRepository}
}

func (uc *electricityUseCase) CreateElectricityUseCase(payload *model.Electricity) (*model.Electricity, error) {

	electricity, err := uc.electricityRepository.CreateElectricityRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("error creating electricity in database: %w", err)
	}
	return electricity, err

}

func (uc *electricityUseCase) GetAllElectricityUseCase(page, limit int) ([]*model.Electricity, error) {
	electricity, err := uc.electricityRepository.GetAllElectricityRepository(page, limit)

	if err != nil {
		return nil, err
	}

	return electricity, nil
}

func (uc *electricityUseCase) GetElectricityByIdUseCase(electricityId string) (*model.Electricity, error) {

	electricity, err := uc.electricityRepository.GetElectricityByIdRepository(electricityId)
	if err != nil {
		return nil, errors.New("electricity not found")
	}

	return electricity, nil

}

func (uc *electricityUseCase) UpdateElectricityByIdUseCase(electricityId string, payload *model.Electricity) (*model.Electricity, error) {
	electricity, err := uc.electricityRepository.GetElectricityByIdRepository(electricityId)
	if err != nil {
		return nil, fmt.Errorf("failed to update electricity: %v", err)
	}

	electricity.ProviderName = payload.ProviderName
	electricity.Type = payload.Type
	electricity.UpdatedAt = time.Now()

	updatedelectricity, err := uc.electricityRepository.UpdateElectricityByIdRepository(electricityId, electricity)
	if err != nil {
		return nil, fmt.Errorf("failed to update electricity: %v", err)
	}

	return updatedelectricity, nil
}

func (uc *electricityUseCase) DeleteElectricityByIDUseCase(userId string) error {
	err := uc.electricityRepository.DeleteElectricityByIdRepository(userId)
	if err != nil {
		return errors.New("electricity not found")
	}
	return err
}

// Tagihan
func (uc *electricityUseCase) PostBillInquiryElectricityUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {
	lastDigit := payload.CustomerId[len(payload.CustomerId)-1]
	if lastDigit == '9' {
		return nil, errors.New("invalid customer ID")
	}

	vaNumber := generateVANumber(16)
	payload.PartnerTxId = fmt.Sprintf("POSTPAID-%s", vaNumber)

	currentTime := time.Now()
	currentMonth := currentTime.Month().String()
	currentYear := strconv.Itoa(currentTime.Year())

	payload.Period = currentMonth + "-" + currentYear
	productype := strings.ToLower(payload.ProductId)

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	existingElectricity, err := uc.transactionRepository.GetProductDetailsByPeriodAndCustomerID(model.GetProductDetail{
		ProductId:  productype,
		Period:     payload.Period,
		CustomerId: payload.CustomerId,
	})
	if err == nil {
		if existingElectricity.Status == model.STATUS_SUCCESSFUL {
			return nil, errors.New("this month's bill has been paid")
		}
		if existingElectricity != nil {
			return existingElectricity, nil
		}
	}

	discount, err := uc.discountRepository.GetDiscountByIdRepository(payload.DiscountId)
	if err != nil {
		return nil, errors.New("discount Not Found")
	}

	electricity, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, err
	}

	powers := []int{450, 900, 1300, 2200, 3500, 4400, 5500, 7700}
	randomIndex := rand.Intn(len(powers))
	randomPower := powers[randomIndex]
	min := 100
	max := 100000
	amount := rand.Intn(max-min+1) + min
	price := calculatePricePower(randomPower, amount)

	totalPrice := float64(price) + electricity.AdminFee - float64(discount.DiscountPrice)

	productDetail := &model.Electricity{
		Period:          payload.Period,
		Name:            user.Name,
		CustomerId:      payload.CustomerId,
		ProviderName:    electricity.ProductID,
		ElectricalPower: randomPower,
		Type:            electricity.ProductID,
		DiscountId:      discount.ID,
		Price:           float64(price),
	}
	transaction := &model.Transaction{
		ID:            electricity.PartnerTxID,
		UserID:        userId,
		Status:        model.STATUS_UNPAID,
		ProductType:   productype,
		Description:   fmt.Sprintf("Pembayaran Tagihan Listrik %s ", payload.Period),
		DiscountPrice: float64(discount.DiscountPrice),
		AdminFee:      electricity.AdminFee,
		Price:         float64(price),
		TotalPrice:    float64(totalPrice),
		ProductDetail: productDetail,
	}

	createdTransaction, err := uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating electricity in database: %w", err)
	}

	return createdTransaction, nil
}

func (uc *electricityUseCase) PostPayBillElectricityUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

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

	var electricity model.Electricity
	err = json.Unmarshal([]byte(jsonData), &electricity)
	if err != nil {
		return nil, fmt.Errorf("error Unmarshal: %w", err)
	}

	var token string

	if electricity.Type == "plnpre" {
		token = generateVANumber(20)
	}

	productDetail := &model.Electricity{
		Name:            user.Name,
		ProviderName:    electricity.ProviderName,
		Type:            electricity.Type,
		Period:          electricity.Period,
		Token:           token,
		ElectricalPower: electricity.ElectricalPower,
		DiscountId:      electricity.DiscountId,
		Price:           float64(payload.Amount),
	}
	updateTransaction := &model.Transaction{

		Status:        model.STATUS_SUCCESSFUL,
		ProductDetail: productDetail,
		UpdatedAt:     time.Now(),
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
	// 	OrderId:         transaction.ID,
	// 	CustomerName:    user.Name,
	// 	Status:          resp.Status,
	// 	RecipentEmail:   user.Email,
	// 	ElectricalPower: electricity.ElectricalPower,
	// 	Token:           token,
	// 	ProductType:     "ELECTRICITY",
	// 	TransactionAt:   resp.UpdatedAt,
	// 	Description:     transaction.Description,
	// 	AdminFee:        transaction.AdminFee,
	// 	Price:           transaction.Price,
	// 	TotalPrice:      transaction.TotalPrice,
	// }
	// mail.SendingMail(mailsend)

	return transactionresp, nil
}

// TOKEN
func (uc *electricityUseCase) PreBillInquiryElectricityUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

	lastDigit := payload.CustomerId[len(payload.CustomerId)-1]
	if lastDigit == '9' {
		return nil, errors.New("invalid customer ID")
	}

	vaNumber := generateVANumber(16)
	payload.PartnerTxId = fmt.Sprintf("PREPAID-%s", vaNumber)

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	discount, err := uc.discountRepository.GetDiscountByIdRepository(payload.DiscountId)
	if err != nil {
		return nil, errors.New("discount Not Found")
	}

	electricity, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, err
	}

	powers := []int{450, 900, 1300, 2200, 3500, 4400, 5500, 7700}
	randomIndex := rand.Intn(len(powers))
	randomPower := powers[randomIndex]

	totalPrice := float64(payload.Amount) + electricity.AdminFee - float64(discount.DiscountPrice)

	productDetail := &model.Electricity{
		Name:            user.Name,
		ProviderName:    electricity.ProductID,
		Type:            electricity.ProductID,
		ElectricalPower: randomPower,
		DiscountId:      discount.ID,
		Price:           float64(payload.Amount),
	}
	transaction := &model.Transaction{
		ID:            electricity.PartnerTxID,
		UserID:        user.ID,
		Status:        model.STATUS_PROCESSING,
		ProductType:   electricity.ProductID,
		Description:   fmt.Sprintf("Pembelian Token Listrik %.2f ", payload.Amount),
		DiscountPrice: float64(discount.DiscountPrice),
		AdminFee:      electricity.AdminFee,
		Price:         float64(payload.Amount),
		TotalPrice:    float64(totalPrice),
		ProductDetail: productDetail,
	}

	resp, err := uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating electricity in database: %w", err)
	}

	return resp, nil
}

func (uc *electricityUseCase) BillElectricityStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {

	electricity, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve electricity: %v", err)
	}

	return electricity, nil
}

func (uc *electricityUseCase) ElectricityBillInquiryIakUseCase(payload *model.IakInquiryBody) (*model.IakPostPaidResponse, error) {
	electricity, err := uc.iakRepository.ElectricityBillInquiryRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return electricity, nil

}
func (uc *electricityUseCase) ElectricityBillPayIakUseCase(payload *model.IakPayBody, userId string) (*model.IakPostPaidResponse, error) {
	electricityCheck, err := uc.iakRepository.ElectricityBillCheckRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	payload.TrID = electricityCheck.Data.TrID

	totalPrice := float64(electricityCheck.Data.Price)

	fmt.Printf("totalPrice = %f \n", totalPrice)

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, err
	}

	if user.Amount < totalPrice {
		return nil, errors.New("your balance is not enough")
	}

	electricity, err := uc.iakRepository.ElectricityBillPayRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	transaction := &model.Transaction{
		ID:            electricity.Data.RefID,
		UserID:        userId,
		Status:        model.STATUS_SUCCESSFUL,
		ProductType:   "ELECTRICITY_BILL",
		DiscountPrice: 0,
		AdminFee:      float64(electricity.Data.Admin),
		Description:   fmt.Sprintf("Pembayaran Tagihan asuransi %s ", electricity.Data.Period),
		Price:         float64(electricity.Data.Nominal),
		TotalPrice:    totalPrice,
		ProductDetail: electricity.Data,
	}

	_, err = uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating insurance in database: %w", err)
	}

	fmt.Printf("user.Amount1 = %f \n", user.Amount)
	fmt.Printf("transaction.TotalPrice = %f \n", transaction.TotalPrice)

	user.Amount -= transaction.TotalPrice

	fmt.Printf("user.Amount2 = %f \n", user.Amount)

	_, err = uc.userRepository.UpdateUserAmountByIDRepository(userId, user)
	if err != nil {
		return nil, err
	}

	return electricity, nil
}

func (uc *electricityUseCase) ElectricityTokenInquiryIakUseCase(payload *model.PrePaidIakBody) (*model.IakElectricityTokenInquiry, error) {
	electricity, err := uc.iakRepository.ElectricityTokenInquiryRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return electricity, nil

}
func (uc *electricityUseCase) ElectricityTokenPayIakUseCase(payload *model.PrePaidIakBody, userId string) (*model.PrePaidIakResponse, error) {
	payload.RefID = uuid.New().String()
	electricity, err := uc.iakRepository.ElectricityTokenInquiryRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	totalPrice := totalPriceToken(payload.ProductCode)

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, err
	}

	if user.Amount < totalPrice {
		return nil, errors.New("your balance is not enough")
	}

	pay, err := uc.iakRepository.IakTopUpPayRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	transaction := &model.Transaction{
		ID:            payload.RefID,
		UserID:        userId,
		Status:        model.STATUS_SUCCESSFUL,
		ProductType:   "ELECTRICITY_TOKEN",
		DiscountPrice: 0,
		AdminFee:      float64(500),
		Description:   fmt.Sprintf("Pembayaran Token Listrik code %s ", payload.ProductCode),
		Price:         float64(500),
		TotalPrice:    totalPrice,
		ProductDetail: electricity.Data,
	}

	_, err = uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating insurance in database: %w", err)
	}

	fmt.Printf("user.Amount1 = %f \n", user.Amount)
	fmt.Printf("transaction.TotalPrice = %f \n", transaction.TotalPrice)

	user.Amount -= transaction.TotalPrice

	fmt.Printf("user.Amount2 = %f \n", user.Amount)

	_, err = uc.userRepository.UpdateUserAmountByIDRepository(userId, user)
	if err != nil {
		return nil, err
	}

	return pay, nil

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

func calculatePricePower(power, amount int) float64 {
	totalPrice := 1.0
	switch {
	case power <= 1300:
		totalPrice = 1.0
	case power <= 3500:
		totalPrice = 1.5
	case power <= 5500:
		totalPrice = 2.0
	default:
		totalPrice = 2.5
	}

	return float64(amount) * totalPrice
}

func totalPriceToken(code string) float64 {
	switch code {
	case "hpln20000":
		return 21000
	case "hpln50000":
		return 51000
	case "hpln100000":
		return 101000
	case "hpln200000":
		return 201000
	case "hpln500000":
		return 501000
	default:
		return 0
	}
}
