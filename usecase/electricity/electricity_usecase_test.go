package electricity_test

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/darulfh/skuy_pay_be/model"
// 	"github.com/darulfh/skuy_pay_be/repository/mocks"
// 	"github.com/darulfh/skuy_pay_be/usecase/electricity"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// 	"gorm.io/gorm"
// )

// type ElectricityUsecaseTest struct {
// 	suite.Suite
// 	electricityUsecase electricity.ElectricityUseCase
// 	electricityRepo    *mocks.ElectricityRepository
// 	userRepo           *mocks.UserRepository
// 	discountRepo       *mocks.DiscountRepository
// 	transactionRepo    *mocks.TransactionRepository
// 	billerOyApiRepo    *mocks.BillerOyApiRepository
// }

// func TestElectricityUsecase(t *testing.T) {
// 	suite.Run(t, new(ElectricityUsecaseTest))
// }

// func (m *ElectricityUsecaseTest) SetupTest() {
// 	m.electricityRepo = &mocks.ElectricityRepository{}
// 	m.userRepo = &mocks.UserRepository{}
// 	m.discountRepo = &mocks.DiscountRepository{}
// 	m.transactionRepo = &mocks.TransactionRepository{}
// 	m.billerOyApiRepo = &mocks.BillerOyApiRepository{}
// 	m.electricityUsecase = electricity.NewElectricityUseCase(m.electricityRepo, m.userRepo, m.discountRepo, m.transactionRepo, m.billerOyApiRepo)
// }

// func (m *ElectricityUsecaseTest) TestCreateElectricityUseCaseSuccess() {
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}

// 	m.electricityRepo.On("CreateElectricityRepository", mockElectricity).Return(mockElectricity, nil)

// 	resp, err := m.electricityUsecase.CreateElectricityUseCase(mockElectricity)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// 	assert.Equal(m.T(), resp.ID, mockElectricity.ID)
// }

// func (m *ElectricityUsecaseTest) TestCreateElectricityUseCaseError() {
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}

// 	m.electricityRepo.On("CreateElectricityRepository", mockElectricity).Return(nil, errors.New("repository error"))

// 	_, err := m.electricityUsecase.CreateElectricityUseCase(mockElectricity)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestGetAllElectricityUseCaseSuccess() {
// 	mockElectricity := []*model.Electricity{
// 		{
// 			UUIDPrimaryKey: model.UUIDPrimaryKey{
// 				ID:        "1",
// 				DeletedAt: gorm.DeletedAt{},
// 				CreatedAt: time.Now(),
// 				UpdatedAt: time.Now(),
// 			},
// 		},
// 	}

// 	m.electricityRepo.On("GetAllElectricityRepository", 1, 1).Return(mockElectricity, nil)

// 	resp, err := m.electricityUsecase.GetAllElectricityUseCase(1, 1)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// 	assert.Equal(m.T(), resp[0].ID, mockElectricity[0].ID)
// }

// func (m *ElectricityUsecaseTest) TestGetAllElectricityUseCaseError() {
// 	m.electricityRepo.On("GetAllElectricityRepository", 1, 1).Return(nil, errors.New("repository error"))

// 	_, err := m.electricityUsecase.GetAllElectricityUseCase(1, 1)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestGetElectricityByIdUseCaseSuccess() {
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}

// 	m.electricityRepo.On("GetElectricityByIdRepository", "id").Return(mockElectricity, nil)

// 	resp, err := m.electricityUsecase.GetElectricityByIdUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// 	assert.Equal(m.T(), resp.ID, mockElectricity.ID)
// }

// func (m *ElectricityUsecaseTest) TestGetElectricityByIdUseCaseError() {
// 	m.electricityRepo.On("GetElectricityByIdRepository", "id").Return(nil, errors.New("repository error"))

// 	_, err := m.electricityUsecase.GetElectricityByIdUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestUpdateElectricityByIdUseCaseSuccess() {
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		ProviderName: "prepaid",
// 		Type:         "Type",
// 	}

// 	m.electricityRepo.On("GetElectricityByIdRepository", "id").Return(mockElectricity, nil)
// 	m.electricityRepo.On("UpdateElectricityByIdRepository", "id", mockElectricity).Return(mockElectricity, nil)

// 	resp, err := m.electricityUsecase.UpdateElectricityByIdUseCase("id", mockElectricity)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}

// 	assert.Equal(m.T(), resp.ID, mockElectricity.ID)
// }

// func (m *ElectricityUsecaseTest) TestUpdateElectricityByIdUseCaseErrorGetID() {
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		ProviderName: "prepaid",
// 		Type:         "Type",
// 	}
// 	m.electricityRepo.On("GetElectricityByIdRepository", "id").Return(nil, errors.New("repository error"))

// 	_, err := m.electricityUsecase.UpdateElectricityByIdUseCase("id", mockElectricity)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestUpdateElectricityByIdUseCaseErrorUpdate() {
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		ProviderName: "prepaid",
// 		Type:         "Type",
// 	}
// 	m.electricityRepo.On("GetElectricityByIdRepository", "id").Return(mockElectricity, nil)
// 	m.electricityRepo.On("UpdateElectricityByIdRepository", "id", mockElectricity).Return(nil, errors.New("repository error"))

// 	_, err := m.electricityUsecase.UpdateElectricityByIdUseCase("id", mockElectricity)
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestDeleteElectricityByIDUseCaseSuccess() {
// 	m.electricityRepo.On("DeleteElectricityByIdRepository", "id").Return(nil)

// 	err := m.electricityUsecase.DeleteElectricityByIDUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestDeleteElectricityByIDUseCaseError() {
// 	m.electricityRepo.On("DeleteElectricityByIdRepository", "id").Return(errors.New("repository error"))

// 	err := m.electricityUsecase.DeleteElectricityByIDUseCase("id")
// 	if err != nil {
// 		assert.Error(m.T(), err, "repository error")
// 	}
// }

// func (m *ElectricityUsecaseTest) TestPostBillInquiryElectricityUseCase() {
// 	userId := "user123"
// 	payload := &model.OyBillerApi{
// 		CustomerId: "12345678",
// 		ProductId:  "electricity",
// 		DiscountId: "1",
// 	}

// 	mockUser := &model.User{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        userId,
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		Name:    "arby",
// 		Email:   "arby@mail.com",
// 		Address: "Jl",
// 		Phone:   "08123456789",
// 	}
// 	mockExistingElectricity := &model.Transaction{
// 		ID:     "1",
// 		Status: model.STATUS_UNPAID,
// 	}
// 	mockDiscount := &model.Discount{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		DiscountCode: "LEBARAN",
// 		Description:  "Hari Raya",
// 		Image:        "url",
// 	}
// 	mockElectricity := &model.Electricity{
// 		UUIDPrimaryKey: model.UUIDPrimaryKey{
// 			ID:        "1",
// 			DeletedAt: gorm.DeletedAt{},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}
// 	mockTransaction := &model.Transaction{
// 		ID:            "transaction123",
// 		UserID:        userId,
// 		Status:        model.STATUS_UNPAID,
// 		ProductType:   payload.ProductId,
// 		Description:   "Pembayaran Tagihan Listrik January-2023",
// 		DiscountPrice: float64(mockDiscount.DiscountPrice),
// 		AdminFee:      2000,
// 		Price:         5000,
// 		TotalPrice:    7000,
// 	}

// 	m.userRepo.On("GetUserByIDRepository", userId).Return(mockUser, nil)
// 	m.transactionRepo.On("GetProductDetailsByPeriodAndCustomerID", mock.AnythingOfType("model.GetProductDetail")).Return(mockExistingElectricity, nil)
// 	m.discountRepo.On("GetDiscountByIdRepository", payload.DiscountId).Return(mockDiscount, nil)
// 	m.billerOyApiRepo.On("BillInquryRepository", payload).Return(mockElectricity, nil)
// 	m.transactionRepo.On("CreateTransactionByUserIdRepository", mockTransaction).Return(mockTransaction, nil)

// 	_, err := m.electricityUsecase.PostBillInquiryElectricityUseCase(userId, payload)

// 	// Assertions
// 	assert.NoError(m.T(), err)
// 	assert.Equal(m.T(), mockTransaction, mockTransaction)

// 	m.userRepo.AssertCalled(m.T(), "GetUserByIDRepository", userId)
// 	m.transactionRepo.AssertCalled(m.T(), "GetProductDetailsByPeriodAndCustomerID", mock.AnythingOfType("model.GetProductDetail"))
// }
