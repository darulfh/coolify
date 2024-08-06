package insurance

// import (
// 	"errors"
// 	"math/rand"
// 	"strconv"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/darulfh/skuy_pay_be/model"
// 	"github.com/darulfh/skuy_pay_be/repository/mocks"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// )

// type InsuranceUsecaseTest struct {
// 	suite.Suite
// 	insuranceUsecase InsuranceUseCase
// 	insuraceRepo     *mocks.InsuranceRepository
// 	userRepo         *mocks.UserRepository
// 	discountRepo     *mocks.DiscountRepository
// 	transactionRepo  *mocks.TransactionRepository
// 	billerOyApiRepo  *mocks.BillerOyApiRepository
// }

// func TestInsuranceUsecase(t *testing.T) {
// 	suite.Run(t, new(InsuranceUsecaseTest))
// }

// func (m *InsuranceUsecaseTest) SetupTest() {
// 	m.insuraceRepo = &mocks.InsuranceRepository{}
// 	m.userRepo = &mocks.UserRepository{}
// 	m.discountRepo = &mocks.DiscountRepository{}
// 	m.transactionRepo = &mocks.TransactionRepository{}
// 	m.billerOyApiRepo = &mocks.BillerOyApiRepository{}
// 	m.insuranceUsecase = NewInsuranceUseCase(m.insuraceRepo, m.userRepo, m.discountRepo, m.transactionRepo, m.billerOyApiRepo)
// }

// func (m *InsuranceUsecaseTest) TestCreateInsuranceUseCaseSuccess() {
// 	mockInsurance := &model.Insurance{
// 		CustomerID: "123213213",
// 	}

// 	m.insuraceRepo.On("CreateInsuranceRepository", mockInsurance).Return(mockInsurance, nil)

// 	resp, err := m.insuranceUsecase.CreateInsuranceUseCase(mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// 	assert.Equal(m.T(), resp.CustomerID, mockInsurance.CustomerID)

// }

// func (m *InsuranceUsecaseTest) TestCreateInsuranceUseCaseError() {
// 	mockInsurance := &model.Insurance{
// 		CustomerID: "123213213",
// 	}

// 	m.insuraceRepo.On("CreateInsuranceRepository", mockInsurance).Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.CreateInsuranceUseCase(mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// }

// func (m *InsuranceUsecaseTest) TestGetAllInsuranceUseCaseSuccess() {
// 	mockInsurance := []*model.Insurance{
// 		{
// 			CustomerID: "123213213",
// 		},
// 	}

// 	m.insuraceRepo.On("GetAllInsuranceRepository", 1, 1).Return(mockInsurance, nil)

// 	resp, err := m.insuranceUsecase.GetAllInsuranceUseCase(1, 1)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// 	assert.Equal(m.T(), resp[0].CustomerID, mockInsurance[0].CustomerID)

// }

// func (m *InsuranceUsecaseTest) TestGetAllInsuranceUseCaseError() {

// 	m.insuraceRepo.On("GetAllInsuranceRepository", 1, 1).Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.GetAllInsuranceUseCase(1, 1)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// }

// func (m *InsuranceUsecaseTest) TestGetInsuranceByIdUseCaseSuccess() {
// 	mockInsurance := &model.Insurance{
// 		CustomerID: "123213213",
// 	}

// 	m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(mockInsurance, nil)

// 	resp, err := m.insuranceUsecase.GetInsuranceByIdUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// 	assert.Equal(m.T(), resp.CustomerID, mockInsurance.CustomerID)

// }

// func (m *InsuranceUsecaseTest) TestGetInsuranceByIdUseCaseError() {

// 	m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.GetInsuranceByIdUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestUpdateInsuranceByIdUseCaseSuccess() {
// 	mockInsurance := &model.Insurance{
// 		CustomerID:   "123213213",
// 		ProviderName: "BPJS",
// 		Type:         "Type",
// 	}

// 	m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(mockInsurance, nil)
// 	m.insuraceRepo.On("UpdateInsuranceByIdRepository", "id", mockInsurance).Return(mockInsurance, nil)

// 	resp, err := m.insuranceUsecase.UpdateInsuranceByIdUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// 	assert.Equal(m.T(), resp.CustomerID, mockInsurance.CustomerID)

// }

// func (m *InsuranceUsecaseTest) TestUpdateInsuranceByIdUseCaseErrorGetID() {
// 	mockInsurance := &model.Insurance{
// 		CustomerID:   "123213213",
// 		ProviderName: "BPJS",
// 		Type:         "Type",
// 	}
// 	m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.UpdateInsuranceByIdUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestUpdateInsuranceByIdUseCaseErrorUpdate() {
// 	mockInsurance := &model.Insurance{
// 		CustomerID:   "123213213",
// 		ProviderName: "BPJS",
// 		Type:         "Type",
// 	}
// 	m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(mockInsurance, nil)
// 	m.insuraceRepo.On("UpdateInsuranceByIdRepository", "id", mockInsurance).Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.UpdateInsuranceByIdUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestDeleteInsuranceByIDUseCaseSuccess() {

// 	m.insuraceRepo.On("DeleteInsuranceByIdRepository", "id").Return(nil)

// 	err := m.insuranceUsecase.DeleteInsuranceByIDUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")

// 	}

// }

// func (m *InsuranceUsecaseTest) TestDeleteInsuranceByIDUseCaseError() {

// 	m.insuraceRepo.On("DeleteInsuranceByIdRepository", "id").Return(errors.New("repository error"))

// 	err := m.insuranceUsecase.DeleteInsuranceByIDUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// func TestGenerateVANumber(t *testing.T) {
// 	length := 10

// 	vaNumber := generateVANumber(length)

// 	if len(vaNumber) != length {
// 		t.Errorf("expected VA number length %d, but got %d", length, len(vaNumber))
// 	}
// 	// Add more assertions for the VA number if needed.
// }

// func TestGenerateRandomClass(t *testing.T) {
// 	class := generateRandomClass()

// 	if class < 1 || class > 3 {
// 		t.Errorf("unexpected class value: %d", class)
// 	}
// 	// Add more assertions for the class if needed.
// }

// func TestGenerateRandomNumberOfFamilyMembers(t *testing.T) {
// 	min := 1
// 	max := 3

// 	numFamilyMembers := generateRandomNumberOfFamilyMembers()

// 	if numFamilyMembers < min || numFamilyMembers > max {
// 		t.Errorf("unexpected number of family members: %d", numFamilyMembers)
// 	}
// 	// Add more assertions for the number of family members if needed.
// }

// func TestCalculateBPJSKesehatanInsurance(t *testing.T) {
// 	mockTest := []struct {
// 		classss int
// 		iuran   float64
// 	}{
// 		{1, 150000.0},
// 		{2, 100000.0},
// 		{3, 35000.0},
// 	}
// 	numFamilyMembers := 2

// 	kelas3PbpuBantuan := 7000.0

// 	for _, v := range mockTest {
// 		expectedTotalIuran := v.iuran * float64(numFamilyMembers)
// 		if v.classss == 3 {
// 			expectedTotalIuran -= kelas3PbpuBantuan

// 		}

// 		totalIuran := calculateBPJSKesehatanInsurance(v.classss, numFamilyMembers)
// 		if totalIuran != expectedTotalIuran {
// 			t.Errorf("expected total iuran %.2f, but got %.2f", expectedTotalIuran, totalIuran)
// 		}

// 	}

// }

// // func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseSuccess() {
// // 	mockInsurance := &model.Insurance{
// // 		CustomerID:   "123213213",
// // 		ProviderName: "BPJS",
// // 		Type:         "Type",
// // 	}

// // 	m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(mockInsurance, nil)
// // 	m.insuraceRepo.On("UpdateInsuranceByIdRepository", "id", mockInsurance).Return(mockInsurance, nil)

// // 	resp, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// // 	if err != nil {
// // 		assert.Error(m.T(), err, "repository error")

// // 	}

// // 	assert.Equal(m.T(), resp.CustomerID, mockInsurance.CustomerID)

// // }

// func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseErrorDigit() {
// 	mockInsurance := &model.OyBillerApi{
// 		CustomerId: "123213219",
// 	}
// 	// m.insuraceRepo.On("GetInsuranceyIdRepository", "id").Return(mockInsurance, nil)
// 	// m.insuraceRepo.On("UpdateInsuranceByIdRepository", "id", mockInsurance).Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseErrorGetUserId() {
// 	mockInsurance := &model.OyBillerApi{
// 		CustomerId: "123213217",
// 	}
// 	m.userRepo.On("GetUserByIDRepository", "id").Return(nil, errors.New("unauthorized"))
// 	// m.insuraceRepo.On("UpdateInsuranceByIdRepository", "id", mockInsurance).Return(nil, errors.New("repository error"))

// 	_, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "unauthorized")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseErrorGetProductByPeriod() {
// 	mockInsurance := &model.OyBillerApi{
// 		CustomerId: "123213217",
// 		ProductId:  "sadasdsa",
// 	}

// 	mockUser := &model.User{
// 		Name: "User",
// 	}

// 	currentTime := time.Now()
// 	currentMonth := currentTime.Month().String()
// 	currentYear := strconv.Itoa(currentTime.Year())

// 	mockInsurance.Period = currentMonth + "-" + currentYear
// 	productype := strings.ToLower(mockInsurance.ProductId)

// 	mockPayload := model.GetProductDetail{
// 		ProductId:  productype,
// 		Period:     mockInsurance.Period,
// 		CustomerId: mockInsurance.CustomerId,
// 	}

// 	mockTransaction := []*model.Transaction{
// 		{Status: model.STATUS_SUCCESSFUL},
// 		{Status: model.STATUS_UNPAID},
// 	}

// 	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
// 	for _, v := range mockTransaction {
// 		if v.Status == model.STATUS_SUCCESSFUL {
// 			m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(v, nil)

// 			_, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// 			if err != nil {
// 				assert.Error(m.T(), err, "this month's bill has been paid")
// 			}

// 		}

// 	}

// }

// func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseSuccessGetProductByPeriod() {
// 	mockInsurance := &model.OyBillerApi{
// 		CustomerId: "123213217",
// 		ProductId:  "sadasdsa",
// 	}

// 	mockUser := &model.User{
// 		Name: "User",
// 	}

// 	currentTime := time.Now()
// 	currentMonth := currentTime.Month().String()
// 	currentYear := strconv.Itoa(currentTime.Year())

// 	mockInsurance.Period = currentMonth + "-" + currentYear
// 	productype := strings.ToLower(mockInsurance.ProductId)

// 	mockPayload := model.GetProductDetail{
// 		ProductId:  productype,
// 		Period:     mockInsurance.Period,
// 		CustomerId: mockInsurance.CustomerId,
// 	}

// 	mockTransaction := &model.Transaction{
// 		Status: model.STATUS_UNPAID,
// 	}

// 	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
// 	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(mockTransaction, nil)

// 	resp, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "this month's bill has been paid")
// 	}

// 	assert.Equal(m.T(), resp.Status, mockTransaction.Status)

// }

// func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseErrorDiscount() {
// 	mockInsurance := &model.OyBillerApi{
// 		CustomerId: "123213217",
// 		ProductId:  "sadasdsa",
// 		DiscountId: "id",
// 	}

// 	mockUser := &model.User{
// 		Name: "User",
// 	}

// 	currentTime := time.Now()
// 	currentMonth := currentTime.Month().String()
// 	currentYear := strconv.Itoa(currentTime.Year())

// 	mockInsurance.Period = currentMonth + "-" + currentYear
// 	productype := strings.ToLower(mockInsurance.ProductId)

// 	mockPayload := model.GetProductDetail{
// 		ProductId:  productype,
// 		Period:     mockInsurance.Period,
// 		CustomerId: mockInsurance.CustomerId,
// 	}

// 	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
// 	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(nil, errors.New("not found"))
// 	m.discountRepo.On("GetDiscountByIdRepository", "id").Return(nil, errors.New("discount Not Found"))

// 	_, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "discount Not Found")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseErrorBiller() {
// 	mockInsurance := &model.OyBillerApi{
// 		CustomerId: "123213217",
// 		ProductId:  "sadasdsa",
// 		DiscountId: "id",
// 	}

// 	mockUser := &model.User{
// 		Name: "User",
// 	}

// 	currentTime := time.Now()
// 	currentMonth := currentTime.Month().String()
// 	currentYear := strconv.Itoa(currentTime.Year())

// 	mockInsurance.Period = currentMonth + "-" + currentYear
// 	productype := strings.ToLower(mockInsurance.ProductId)

// 	mockPayload := model.GetProductDetail{
// 		ProductId:  productype,
// 		Period:     mockInsurance.Period,
// 		CustomerId: mockInsurance.CustomerId,
// 	}

// 	mockDiscount := &model.Discount{}

// 	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
// 	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(nil, errors.New("not found"))
// 	m.discountRepo.On("GetDiscountByIdRepository", "id").Return(mockDiscount, nil)
// 	m.billerOyApiRepo.On("BillInquryRepository", mockInsurance).Return(nil, errors.New("repo error"))

// 	_, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repo error")
// 	}

// }

// // func (m *InsuranceUsecaseTest) TestBillInquiryInsuranceUseCaseErrorTransaction() {
// // 	mockInsurance := &model.OyBillerApi{
// // 		CustomerId: "123213217",
// // 		ProductId:  "sadasdsa",
// // 		DiscountId: "id",
// // 	}

// // 	mockUser := &model.User{
// // 		Name: "User",
// // 	}

// // 	currentTime := time.Now()
// // 	currentMonth := currentTime.Month().String()
// // 	currentYear := strconv.Itoa(currentTime.Year())

// // 	mockInsurance.Period = currentMonth + "-" + currentYear
// // 	productype := strings.ToLower(mockInsurance.ProductId)

// // 	mockPayload := model.GetProductDetail{
// // 		ProductId:  productype,
// // 		Period:     mockInsurance.Period,
// // 		CustomerId: mockInsurance.CustomerId,
// // 	}

// // 	mockDiscount := &model.Discount{}
// // 	numberOfFamilyMembers := generateRandomNumberOfFamilyMembers()

// // 	insurace := &model.OyBillerApiResponse{}
// // 	class := generateRandomClass()

// // 	amount := calculateBPJSKesehatanInsurance(class, numberOfFamilyMembers)

// // 	productDetail := &model.Insurance{
// // 		Period:         mockInsurance.Period,
// // 		CustomerID:     mockInsurance.CustomerId,
// // 		ProviderName:   insurace.ProductID,
// // 		Name:           mockUser.Name,
// // 		NumberOffamily: numberOfFamilyMembers,
// // 		Type:           insurace.ProductID,
// // 		DiscountId:     mockDiscount.ID,
// // 		Price:          float64(amount),
// // 	}

// // 	totalPrice := float64(amount) + insurace.AdminFee - float64(mockDiscount.DiscountPrice)

// // 	transaction := &model.Transaction{
// // 		ID:            insurace.PartnerTxID,
// // 		UserID:        "id",
// // 		Status:        model.STATUS_UNPAID,
// // 		ProductType:   productype,
// // 		DiscountPrice: float64(mockDiscount.DiscountPrice),
// // 		AdminFee:      insurace.AdminFee,
// // 		Description:   fmt.Sprintf("Pembayaran Tagihan asuransi %s ", mockInsurance.Period),
// // 		Price:         float64(amount),
// // 		TotalPrice:    float64(totalPrice),
// // 		ProductDetail: productDetail,
// // 	}

// // 	m.userRepo.On("GetUserByIDRepository", "id").Return(mockUser, nil)
// // 	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mockPayload).Return(nil, errors.New("not found"))
// // 	m.discountRepo.On("GetDiscountByIdRepository", "id").Return(mockDiscount, nil)
// // 	m.billerOyApiRepo.On("BillInquryRepository", mockInsurance).Return(insurace, nil)

// // 	// m.transactionRepo.On("CreateTransactionByUserIdRepository", transaction).Return(nil, errors.New("repo error"))

// // 	_, err := m.insuranceUsecase.BillInquiryInsuranceUseCase("id", mockInsurance)
// // 	if err != nil {
// // 		assert.Error(m.T(), err, "repo error")
// // 	}

// // }

// func (m *InsuranceUsecaseTest) TestBillInsuranceStatusUseCaseSuccess() {
// 	mockPayload := &model.OyBillerApi{}
// 	mockResponse := &model.OyBillerApiResponse{}

// 	m.billerOyApiRepo.On("BillInquryRepository", mockPayload).Return(mockResponse, nil)

// 	resp, err := m.insuranceUsecase.BillInsuranceStatusUseCase(mockPayload)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// 	assert.Equal(m.T(), resp.OyBillerStatus.Message, mockResponse.OyBillerStatus.Message)

// }

// func (m *InsuranceUsecaseTest) TestBillInsuranceStatusUseCaseError() {
// 	mockPayload := &model.OyBillerApi{}

// 	m.billerOyApiRepo.On("BillInquryRepository", mockPayload).Return(nil, errors.New("repo error"))

// 	_, err := m.insuranceUsecase.BillInsuranceStatusUseCase(mockPayload)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestPayBillInsuranceUseCase_ErrorTransaction() {
// 	userID := "userID"
// 	mockPayload := &model.OyBillerApi{
// 		PartnerTxId: "sadsadsadasd",
// 	}

// 	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(nil, errors.New("repo error"))

// 	_, err := m.insuranceUsecase.PayBillInsuranceUseCase(userID, mockPayload)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestPayBillInsuranceUseCase_ErrorTransactionStatus() {
// 	userID := "userID"
// 	mockPayload := &model.OyBillerApi{
// 		PartnerTxId: "sadsadsadasd",
// 	}

// 	mockTransaction := &model.Transaction{
// 		Status: model.STATUS_SUCCESSFUL,
// 	}

// 	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)

// 	_, err := m.insuranceUsecase.PayBillInsuranceUseCase(userID, mockPayload)
// 	if err != nil {
// 		assert.Error(m.T(), err, "this month's bill has been paid")
// 	}

// }

// func (m *InsuranceUsecaseTest) TestPayBillInsuranceUseCase_ErrorUserID() {
// 	userID := "userID"
// 	mockPayload := &model.OyBillerApi{
// 		PartnerTxId: "sadsadsadasd",
// 	}

// 	mockTransaction := &model.Transaction{
// 		Status:     model.STATUS_UNPAID,
// 		TotalPrice: 100000,
// 	}
// 	mockUser := &model.User{
// 		Amount: 1000000,
// 	}

// 	mockUser.Amount -= mockTransaction.TotalPrice

// 	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)
// 	m.userRepo.On("GetUserByIDRepository", userID).Return(mockUser, errors.New("repo error"))

// 	_, err := m.insuranceUsecase.PayBillInsuranceUseCase(userID, mockPayload)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repo error")
// 	}

// }

// // func (m *InsuranceUsecaseTest) TestPayBillInsuranceUseCase_ErrorUserID2() {
// // 	userID := "userID"
// // 	mockPayload := &model.OyBillerApi{
// // 		PartnerTxId: "sadsadsadasd",
// // 	}

// // 	mockTransaction := &model.Transaction{
// // 		Status:     model.STATUS_UNPAID,
// // 		TotalPrice: 100000,
// // 	}
// // 	mockUser := &model.User{
// // 		Amount: 10000,
// // 	}

// // 	transactionFail := &model.Transaction{
// // 		Status:    model.STATUS_FAIL,
// // 		UpdatedAt: time.Now(),
// // 	}

// // 	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)
// // 	m.userRepo.On("GetUserByIDRepository", userID).Return(mockUser, nil)

// // 	m.transactionRepo.On("UpdateTransactionByIdRepository", mockPayload.PartnerTxId, transactionFail).Return(transactionFail, errors.New("your balance is not enough"))

// // 	_, err := m.insuranceUsecase.PayBillInsuranceUseCase(userID, mockPayload)
// // 	if err != nil {
// // 		assert.Error(m.T(), err, "your balance is not enough")
// // 	}

// // }

// func (m *InsuranceUsecaseTest) TestPayBillInsuranceUseCase_ErrorUpdateUserID() {
// 	userID := "userID"
// 	mockPayload := &model.OyBillerApi{
// 		PartnerTxId: "sadsadsadasd",
// 	}

// 	mockTransaction := &model.Transaction{
// 		Status:     model.STATUS_UNPAID,
// 		TotalPrice: 100000,
// 	}
// 	mockUser := &model.User{
// 		Amount: 1000000,
// 	}

// 	// transactionFail := &model.Transaction{
// 	// 	Status:    model.STATUS_FAIL,
// 	// 	UpdatedAt: time.Now(),
// 	// }

// 	mockUser.Amount -= mockTransaction.TotalPrice

// 	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)
// 	m.userRepo.On("GetUserByIDRepository", userID).Return(mockUser, nil)
// 	m.userRepo.On("UpdateUserAmountByIDRepository", userID, mockUser).Return(nil, errors.New("your balance is not enough"))

// 	_, err := m.insuranceUsecase.PayBillInsuranceUseCase(userID, mockPayload)
// 	if err != nil {
// 		assert.Error(m.T(), err, "your balance is not enough")
// 	}

// }

// // func (m *InsuranceUsecaseTest) TestPayBillInsuranceUseCase_ErrorJson() {
// // 	userID := "userID"
// // 	mockPayload := &model.OyBillerApi{
// // 		PartnerTxId: "sadsadsadasd",
// // 	}
// // 	productDetail := &model.Insurance{
// // 		CustomerID: "asdsa",
// // 	}
// // 	mockTransaction := &model.Transaction{
// // 		Status:        model.STATUS_UNPAID,
// // 		TotalPrice:    100000,
// // 		ProductDetail: productDetail,
// // 	}
// // 	mockUser := &model.User{
// // 		Amount: 1000000,
// // 	}

// // 	mockUser.Amount -= mockTransaction.TotalPrice

// // 	m.transactionRepo.On("GetTransactionByIdRepository", mockPayload.PartnerTxId).Return(mockTransaction, nil)
// // 	m.userRepo.On("GetUserByIDRepository", userID).Return(mockUser, nil)
// // 	m.userRepo.On("UpdateUserAmountByIDRepository", userID, mockUser).Return(mockUser, nil)
// // 	updateTransaction := &model.Transaction{

// // 		Status:    model.STATUS_SUCCESSFUL,
// // 		UpdatedAt: time.Now(),
// // 	}
// // 	m.transactionRepo.On("UpdateTransactionByIdRepository", mockPayload.PartnerTxId, updateTransaction).Return(updateTransaction, errors.New("error Updating Transactions"))

// // 	_, err := m.insuranceUsecase.PayBillInsuranceUseCase(userID, mockPayload)
// // 	if err != nil {
// // 		assert.Error(m.T(), err, "error Updating Transactions")
// // 	}

// // }
