package insurance

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

type InsuranceUseCase interface {
	CreateInsuranceUseCase(payload *model.Insurance) (*model.Insurance, error)
	GetAllInsuranceUseCase(page, limit int) ([]*model.Insurance, error)
	GetInsuranceByIdUseCase(insuranceId string) (*model.Insurance, error)
	UpdateInsuranceByIdUseCase(insuranceId string, payload *model.Insurance) (*model.Insurance, error)
	DeleteInsuranceByIDUseCase(userId string) error
	BillInquiryInsuranceUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	PayBillInsuranceUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error)
	BillInsuranceStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error)

	BpjsInquiryIakUseCase(payload *model.IakInquiryBody) (*model.BpjsIAKResponse, error)
	BpjsPayIakUseCase(payload *model.IakPayBody, userId string) (*model.BpjsIAKResponse, error)
}

type insuranceUseCase struct {
	insuranceRepository   repository.InsuranceRepository
	userRepository        repository.UserRepository
	discountRepository    repository.DiscountRepository
	transactionRepository repository.TransactionRepository
	billerOyApi           repository.BillerOyApiRepository
	iakRepository         repository.IakApiRepository
}

func NewInsuranceUseCase(insuranceRepository repository.InsuranceRepository, userRepository repository.UserRepository, discountRepository repository.DiscountRepository, transactionRepository repository.TransactionRepository, billerOyApiRepository repository.BillerOyApiRepository, iakRepository repository.IakApiRepository) *insuranceUseCase {
	return &insuranceUseCase{insuranceRepository: insuranceRepository, userRepository: userRepository, discountRepository: discountRepository, transactionRepository: transactionRepository, billerOyApi: billerOyApiRepository, iakRepository: iakRepository}
}

func (uc *insuranceUseCase) CreateInsuranceUseCase(payload *model.Insurance) (*model.Insurance, error) {

	insurance, err := uc.insuranceRepository.CreateInsuranceRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("error creating insurance in database: %w", err)
	}
	return insurance, err

}

func (uc *insuranceUseCase) GetAllInsuranceUseCase(page, limit int) ([]*model.Insurance, error) {
	insurance, err := uc.insuranceRepository.GetAllInsuranceRepository(page, limit)

	if err != nil {
		return nil, err
	}

	return insurance, nil
}

func (uc *insuranceUseCase) GetInsuranceByIdUseCase(insuranceId string) (*model.Insurance, error) {

	insurance, err := uc.insuranceRepository.GetInsuranceyIdRepository(insuranceId)
	if err != nil {
		return nil, errors.New("insurance not found")
	}

	return insurance, nil

}

func (uc *insuranceUseCase) UpdateInsuranceByIdUseCase(insuranceId string, payload *model.Insurance) (*model.Insurance, error) {
	insurance, err := uc.insuranceRepository.GetInsuranceyIdRepository(insuranceId)
	if err != nil {
		return nil, fmt.Errorf("failed to update insurance: %v", err)
	}

	insurance.ProviderName = payload.ProviderName
	insurance.Type = payload.Type
	insurance.UpdatedAt = time.Now()

	updatedinsurance, err := uc.insuranceRepository.UpdateInsuranceByIdRepository(insuranceId, insurance)
	if err != nil {
		return nil, fmt.Errorf("failed to update insurance: %v", err)
	}

	return updatedinsurance, nil
}

func (uc *insuranceUseCase) DeleteInsuranceByIDUseCase(userId string) error {
	err := uc.insuranceRepository.DeleteInsuranceByIdRepository(userId)
	if err != nil {
		return errors.New("insurance not found")
	}
	return err
}

func (uc *insuranceUseCase) BillInquiryInsuranceUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

	lastDigit := payload.CustomerId[len(payload.CustomerId)-1]
	if lastDigit == '9' {
		return nil, errors.New("invalid customer ID")
	}

	vaNumber := generateVANumber(16)
	payload.PartnerTxId = fmt.Sprintf("INSURANCE-%s", vaNumber)

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

	discount, err := uc.discountRepository.GetDiscountByIdRepository(payload.DiscountId)
	if err != nil {
		return nil, errors.New("discount Not Found")
	}

	insurance, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, err
	}

	class := generateRandomClass()
	numberOfFamilyMembers := generateRandomNumberOfFamilyMembers()

	amount := calculateBPJSKesehatanInsurance(class, numberOfFamilyMembers)

	totalPrice := float64(amount) + insurance.AdminFee - float64(discount.DiscountPrice)

	productDetail := &model.Insurance{
		Period:         payload.Period,
		CustomerID:     payload.CustomerId,
		ProviderName:   insurance.ProductID,
		Name:           user.Name,
		NumberOffamily: numberOfFamilyMembers,
		Type:           insurance.ProductID,
		DiscountId:     discount.ID,
		Price:          float64(amount),
	}
	transaction := &model.Transaction{
		ID:            insurance.PartnerTxID,
		UserID:        userId,
		Status:        model.STATUS_UNPAID,
		ProductType:   productype,
		DiscountPrice: float64(discount.DiscountPrice),
		AdminFee:      insurance.AdminFee,
		Description:   fmt.Sprintf("Pembayaran Tagihan asuransi %s ", payload.Period),
		Price:         float64(amount),
		TotalPrice:    float64(totalPrice),
		ProductDetail: productDetail,
	}

	_, err = uc.transactionRepository.CreateTransactionByUserIdRepository(transaction)
	if err != nil {
		return nil, fmt.Errorf("error creating insurance in database: %w", err)
	}
	return transaction, nil
}

func (uc *insuranceUseCase) PayBillInsuranceUseCase(userId string, payload *model.OyBillerApi) (*model.Transaction, error) {

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

	var insurance model.Insurance
	err = json.Unmarshal([]byte(jsonData), &insurance)
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
	// 	OrderId:        transaction.ID,
	// 	CustomerName:   user.Name,
	// 	Status:         resp.Status,
	// 	Class:          insurance.Class,
	// 	ProductType:    "BPJS",
	// 	DiscountPrice:  transaction.DiscountPrice,
	// 	NumberOffamily: insurance.NumberOffamily,
	// 	Period:         insurance.Period,
	// 	RecipentEmail:  user.Email,
	// 	TransactionAt:  resp.UpdatedAt,
	// 	Description:    transaction.Description,
	// 	AdminFee:       transaction.AdminFee,
	// 	Price:          transaction.Price,
	// 	TotalPrice:     transaction.TotalPrice,
	// }
	// mail.SendingMail(mailsend)

	return transactionresp, nil
}

func (uc *insuranceUseCase) BillInsuranceStatusUseCase(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {

	insurance, err := uc.billerOyApi.BillInquryRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return insurance, nil
}

func (uc *insuranceUseCase) BpjsInquiryIakUseCase(payload *model.IakInquiryBody) (*model.BpjsIAKResponse, error) {
	bpjs, err := uc.iakRepository.BpjsInquiryRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return bpjs, nil
}
func (uc *insuranceUseCase) BpjsPayIakUseCase(payload *model.IakPayBody, userId string) (*model.BpjsIAKResponse, error) {
	bpjsCheck, err := uc.iakRepository.BpjsCheckRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	payload.TrID = bpjsCheck.Data.TrID

	fmt.Printf("bpjsCheck.Data.Price = %d \n", bpjsCheck.Data.Price)

	totalPrice := float64(bpjsCheck.Data.Price)

	fmt.Printf("totalPrice = %f \n", totalPrice)

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, err
	}

	if user.Amount < totalPrice {
		return nil, errors.New("your balance is not enough")
	}

	bpjs, err := uc.iakRepository.BpjsPayRepository(payload)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	transaction := &model.Transaction{
		ID:            bpjs.Data.RefID,
		UserID:        userId,
		Status:        model.STATUS_SUCCESSFUL,
		ProductType:   "BPJS",
		DiscountPrice: 0,
		AdminFee:      1000,
		Description:   fmt.Sprintf("Pembayaran Tagihan asuransi %s ", bpjs.Data.Period),
		Price:         float64(bpjs.Data.Nominal),
		TotalPrice:    totalPrice,
		ProductDetail: bpjs.Data,
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

	return bpjs, nil
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

func generateRandomClass() int {
	classes := []int{1, 2, 3}
	randomIndex := rand.Intn(len(classes))
	return classes[randomIndex]
}

func generateRandomNumberOfFamilyMembers() int {
	min := 1
	max := 3
	return rand.Intn(max-min+1) + min
}

func calculateBPJSKesehatanInsurance(class int, numberOfFamilyMembers int) float64 {
	const (
		kelas1IuranPerOrang = 150000.0
		kelas2IuranPerOrang = 100000.0
		kelas3IuranPerOrang = 35000.0
		kelas3PbpuBantuan   = 7000.0
	)

	var iuranPerOrang float64

	switch class {
	case 1:
		iuranPerOrang = kelas1IuranPerOrang
	case 2:
		iuranPerOrang = kelas2IuranPerOrang
	case 3:
		iuranPerOrang = kelas3IuranPerOrang
	}

	totalIuran := iuranPerOrang * float64(numberOfFamilyMembers)

	if class == 3 {
		totalIuran -= kelas3PbpuBantuan
	}

	return totalIuran
}
