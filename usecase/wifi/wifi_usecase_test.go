package wifi

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WifiUsecaseTest struct {
	suite.Suite
	wifiUsecase     WifiUsecase
	wifiRepo        *mocks.WifiRepository
	userRepo        *mocks.UserRepository
	discountRepo    *mocks.DiscountRepository
	transactionRepo *mocks.TransactionRepository
	billerOyApiRepo *mocks.BillerOyApiRepository
}

func TestWifiUsecase(t *testing.T) {
	suite.Run(t, new(WifiUsecaseTest))
}

func (m *WifiUsecaseTest) SetupTest() {
	m.wifiRepo = &mocks.WifiRepository{}
	m.userRepo = &mocks.UserRepository{}
	m.discountRepo = &mocks.DiscountRepository{}
	m.transactionRepo = &mocks.TransactionRepository{}
	m.billerOyApiRepo = &mocks.BillerOyApiRepository{}
	m.wifiUsecase = NewWifiUseCase(m.wifiRepo, m.userRepo, m.discountRepo, m.transactionRepo, m.billerOyApiRepo)
}

func (m *WifiUsecaseTest) TestCreateWifiUseCaseSuccess() {
	mockWifi := &model.Wifi{
		CustomerID: "123213213",
	}

	m.wifiRepo.On("CreateWifiRepository", mockWifi).Return(mockWifi, nil)

	resp, err := m.wifiUsecase.CreateWifiUseCase(mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}

	assert.Equal(m.T(), resp.CustomerID, mockWifi.CustomerID)
}

func (m *WifiUsecaseTest) TestCreateWifiUseCaseError() {
	mockWifi := &model.Wifi{
		CustomerID: "123213213",
	}

	m.wifiRepo.On("CreateWifiRepository", mockWifi).Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.CreateWifiUseCase(mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestGetAllWifiUseCaseSuccess() {
	mockWifi := []*model.Wifi{
		{
			CustomerID: "123213213",
		},
	}

	m.wifiRepo.On("GetAllWifiRepository", 1, 1).Return(mockWifi, nil)

	resp, err := m.wifiUsecase.GetAllWifiUseCase(1, 1)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}

	assert.Equal(m.T(), resp[0].CustomerID, mockWifi[0].CustomerID)
}

func (m *WifiUsecaseTest) TestGetAllWifiUseCaseError() {
	m.wifiRepo.On("GetAllWifiRepository", 1, 1).Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.GetAllWifiUseCase(1, 1)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestGetWifiByIdUseCaseSuccess() {
	mockWifi := &model.Wifi{
		CustomerID: "123213213",
	}

	m.wifiRepo.On("GetWifiByIDRepository", "id").Return(mockWifi, nil)

	resp, err := m.wifiUsecase.GetWifiByIDUseCase("id")
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}

	assert.Equal(m.T(), resp.CustomerID, mockWifi.CustomerID)
}

func (m *WifiUsecaseTest) TestGetWifiByIdUseCaseError() {
	m.wifiRepo.On("GetWifiByIDRepository", "id").Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.GetWifiByIDUseCase("id")
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestGetWifiByCodeUseCaseSuccess() {
	mockWifi := &model.Wifi{
		CustomerID: "123213213",
	}

	m.wifiRepo.On("GetWifiByCodeRepository", "code").Return(mockWifi, nil)

	resp, err := m.wifiUsecase.GetWifiByCodeUseCase("code")
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}

	assert.Equal(m.T(), resp.CustomerID, mockWifi.CustomerID)
}

func (m *WifiUsecaseTest) TestGetWifiByCodeUseCaseError() {
	m.wifiRepo.On("GetWifiByCodeRepository", "code").Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.GetWifiByCodeUseCase("code")
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestUpdateWifiByIDUsecaseSuccess() {
	mockWifi := &model.Wifi{
		CustomerID:   "123213213",
		ProviderName: "telkom",
		ProductType:  "Type",
	}

	m.wifiRepo.On("GetWifiByIDRepository", "id").Return(mockWifi, nil)
	m.wifiRepo.On("UpdateWifiByIDRepository", "id", mockWifi).Return(mockWifi, nil)

	resp, err := m.wifiUsecase.UpdateWifiByIDUseCase("id", mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}

	assert.Equal(m.T(), resp.CustomerID, mockWifi.CustomerID)
}

func (m *WifiUsecaseTest) TestUpdateWifiByIDUsecaseErrorGetID() {
	mockWifi := &model.Wifi{
		CustomerID:   "123213213",
		ProviderName: "telkom",
		ProductType:  "Type",
	}
	m.wifiRepo.On("GetWifiByIDRepository", "id").Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.UpdateWifiByIDUseCase("id", mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestDeleteWifiByIDUseCaseSuccess() {
	m.wifiRepo.On("DeleteWifiByIDRepository", "id").Return(nil)

	err := m.wifiUsecase.DeleteWifiByIDUseCase("id")
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestDeleteWifiByIDUseCaseError() {
	m.wifiRepo.On("DeleteWifiByIDRepository", "id").Return(errors.New("repository error"))

	err := m.wifiUsecase.DeleteWifiByIDUseCase("id")
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestBillWifiStatusUseCaseSuccess() {
	mockPayload := &model.OyBillerApi{}
	mockResponse := &model.OyBillerApiResponse{}

	m.billerOyApiRepo.On("BillInquryRepository", mockPayload).Return(mockResponse, nil)

	resp, err := m.wifiUsecase.BillWifiStatusUseCase(mockPayload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}

	assert.Equal(m.T(), resp.OyBillerStatus.Message, mockResponse.OyBillerStatus.Message)
}

func (m *WifiUsecaseTest) TestBillWifiStatusUseCaseError() {
	mockPayload := &model.OyBillerApi{}

	m.billerOyApiRepo.On("BillInquryRepository", mockPayload).Return(nil, errors.New("repo error"))

	_, err := m.wifiUsecase.BillWifiStatusUseCase(mockPayload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestGenerateVANumber(t *testing.T) {
	length := 10

	vaNumber := generateVANumber(length)

	if len(vaNumber) != length {
		t.Errorf("expected VA number length %d, but got %d", length, len(vaNumber))
	}
	// Add more assertions for the VA number if needed.
}

func TestCalculatePriceBandwidth(t *testing.T) {
	testCases := []struct {
		bandwidth     int
		expectedPrice float64
	}{
		{bandwidth: 10, expectedPrice: 275000.0},
		{bandwidth: 25, expectedPrice: 315000.0},
		{bandwidth: 40, expectedPrice: 445000.0},
		{bandwidth: 60, expectedPrice: 795000.0},
		{bandwidth: 5, expectedPrice: 275000.0},
		{bandwidth: 15, expectedPrice: 275000.0},
		{bandwidth: 35, expectedPrice: 445000.0},
		{bandwidth: 100, expectedPrice: 795000.0},
	}

	for _, tc := range testCases {
		actualPrice := calculatePriceBandwith(tc.bandwidth)
		if actualPrice != tc.expectedPrice {
			t.Errorf("For bandwidth %d, expected price %.2f but got %.2f", tc.bandwidth, tc.expectedPrice, actualPrice)
		}
	}
}

func (m *WifiUsecaseTest) TestBillInquiryWifiUseCaseErrorDigit() {
	mockWifi := &model.OyBillerApi{
		CustomerId: "123213219",
	}
	// m.wifiRepo.On("GetWifiByIdRepository", "id").Return(mockWifi, nil)
	// m.wifiRepo.On("UpdateWifiByIdRepository", "id", mockWifi).Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.BillInquiryWifiUseCase("id", mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestBillInquiryWifiUseCaseErrorGetUserId() {
	mockWifi := &model.OyBillerApi{
		CustomerId: "123213217",
	}
	m.userRepo.On("GetUserByIDRepository", "id").Return(nil, errors.New("unauthorized"))
	// m.wifiRepo.On("UpdateWifiByIdRepository", "id", mockWifi).Return(nil, errors.New("repository error"))

	_, err := m.wifiUsecase.BillInquiryWifiUseCase("id", mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "unauthorized")
	}
}

func (m *WifiUsecaseTest) TestBillInquiryWifiUseCaseErrorGetProductByPeriod() {
	mockWifi := &model.OyBillerApi{
		CustomerId: "123213217",
		ProductId:  "sadasdsa",
	}

	mockUser := &model.User{
		Name: "User",
	}

	currentTime := time.Now()
	currentMonth := currentTime.Month().String()
	currentYear := strconv.Itoa(currentTime.Year())

	mockWifi.Period = currentMonth + "-" + currentYear
	productType := strings.ToLower(mockWifi.ProductId)

	mockPayload := model.GetProductDetail{
		ProductId:  productType,
		Period:     mockWifi.Period,
		CustomerId: mockWifi.CustomerId,
	}

	mockTransaction := []*model.Transaction{
		{Status: model.STATUS_SUCCESSFUL},
		{Status: model.STATUS_UNPAID},
	}

	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
	for _, v := range mockTransaction {
		if v.Status == model.STATUS_SUCCESSFUL {
			m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(v, nil)

			_, err := m.wifiUsecase.BillInquiryWifiUseCase("id", mockWifi)
			if err != nil {
				assert.Error(m.T(), err, "this month's bill has been paid")
			}
		}
	}
}

func (m *WifiUsecaseTest) TestBillInquiryWifiUseCaseSuccessGetProductByPeriod() {
	mockWifi := &model.OyBillerApi{
		CustomerId: "123213217",
		ProductId:  "sadasdsa",
	}

	mockUser := &model.User{
		Name: "User",
	}

	currentTime := time.Now()
	currentMonth := currentTime.Month().String()
	currentYear := strconv.Itoa(currentTime.Year())

	mockWifi.Period = currentMonth + "-" + currentYear
	productType := strings.ToLower(mockWifi.ProductId)

	mockPayload := model.GetProductDetail{
		ProductId:  productType,
		Period:     mockWifi.Period,
		CustomerId: mockWifi.CustomerId,
	}

	mockTransaction := &model.Transaction{
		Status: model.STATUS_UNPAID,
	}

	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(mockTransaction, nil)

	resp, err := m.wifiUsecase.BillInquiryWifiUseCase("id", mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "this month's bill has been paid")
	}

	assert.Equal(m.T(), resp.Status, mockTransaction.Status)
}

func (m *WifiUsecaseTest) TestBillInquiryWifiUseCaseErrorDiscount() {
	mockWifi := &model.OyBillerApi{
		CustomerId: "123213217",
		ProductId:  "sadasdsa",
		DiscountId: "id",
	}

	mockUser := &model.User{
		Name: "User",
	}

	currentTime := time.Now()
	currentMonth := currentTime.Month().String()
	currentYear := strconv.Itoa(currentTime.Year())

	mockWifi.Period = currentMonth + "-" + currentYear
	productType := strings.ToLower(mockWifi.ProductId)

	mockPayload := model.GetProductDetail{
		ProductId:  productType,
		Period:     mockWifi.Period,
		CustomerId: mockWifi.CustomerId,
	}

	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(nil, errors.New("not found"))
	m.discountRepo.On("GetDiscountByIdRepository", "id").Return(nil, errors.New("discount Not Found"))

	_, err := m.wifiUsecase.BillInquiryWifiUseCase("id", mockWifi)
	if err != nil {
		assert.Error(m.T(), err, "discount Not Found")
	}
}

func (m *WifiUsecaseTest) TestPayBillWifiUseCaseErrorTransaction() {
	userID := "userID"
	mockPayload := &model.OyBillerApi{
		PartnerTxId: "sadsadsadasd",
	}

	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(nil, errors.New("repo error"))

	_, err := m.wifiUsecase.PayBillWifiUseCase(userID, mockPayload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestPayBillWifiUseCaseErrorTransactionStatus() {
	userID := "userID"
	mockPayload := &model.OyBillerApi{
		PartnerTxId: "sadsadsadasd",
	}

	mockTransaction := &model.Transaction{
		Status: model.STATUS_SUCCESSFUL,
	}

	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)

	_, err := m.wifiUsecase.PayBillWifiUseCase(userID, mockPayload)
	if err != nil {
		assert.Error(m.T(), err, "this month's bill has been paid")
	}
}

func (m *WifiUsecaseTest) TestPayBillWifiUseCaseErrorUserID() {
	userID := "userID"
	mockPayload := &model.OyBillerApi{
		PartnerTxId: "sadsadsadasd",
	}

	mockTransaction := &model.Transaction{
		Status:     model.STATUS_UNPAID,
		TotalPrice: 100000,
	}
	mockUser := &model.User{
		Amount: 1000000,
	}

	mockUser.Amount -= mockTransaction.TotalPrice

	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)
	m.userRepo.On("GetUserByIDRepository", userID).Return(mockUser, errors.New("repo error"))

	_, err := m.wifiUsecase.PayBillWifiUseCase(userID, mockPayload)
	if err != nil {
		assert.Error(m.T(), err, "repository error")
	}
}

func (m *WifiUsecaseTest) TestPayBillWifiUseCaseErrorUpdateUserID() {
	userID := "userID"
	mockPayload := &model.OyBillerApi{
		PartnerTxId: "sadsadsadasd",
	}

	mockTransaction := &model.Transaction{
		Status:     model.STATUS_UNPAID,
		TotalPrice: 100000,
	}
	mockUser := &model.User{
		Amount: 1000000,
	}
	mockUser.Amount -= mockTransaction.TotalPrice

	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)
	m.userRepo.On("GetUserByIDRepository", userID).Return(mockUser, nil)
	m.userRepo.On("UpdateUserAmountByIDRepository", userID, mockUser).Return(nil, errors.New("your balance is not enough"))

	_, err := m.wifiUsecase.PayBillWifiUseCase(userID, mockPayload)
	if err != nil {
		assert.Error(m.T(), err, "your balance is not enough")
	}
}
