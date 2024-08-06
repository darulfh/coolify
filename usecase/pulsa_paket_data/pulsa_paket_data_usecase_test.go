package pulsa

// import (
// 	"errors"
// 	"testing"

// 	"github.com/darulfh/skuy_pay_be/dto"
// 	"github.com/darulfh/skuy_pay_be/model"
// 	"github.com/darulfh/skuy_pay_be/repository/mocks"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// var mockID = uuid.New().String()

// func TestCreatePulsaPaketDataSuccess(t *testing.T) {
// 	mockPPD := model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	}

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("CreatePulsaPaketData", mockPPD).Return(mockPPD, nil)

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo, mockDiscountRepo)

// 	result, err := service.CreatePulsaPaketData(mockPPD)

// 	if err != nil {
// 		t.Errorf("Got Error %v", err)
// 	}
// 	assert.Equal(t, result.Name, mockPPD.Name)
// 	assert.Equal(t, result.Code, mockPPD.Code)
// 	assert.Equal(t, result.Price, mockPPD.Price)
// 	assert.Equal(t, result.Provider, mockPPD.Provider)
// }

// func TestCreatePulsaPaketDataError(t *testing.T) {
// 	mockPPD := model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	}

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("CreatePulsaPaketData", mockPPD).Return(mockPPD, errors.New("code already exists"))

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	_, err := service.CreatePulsaPaketData(mockPPD)

// 	if err != nil {
// 		assert.Error(t, err, "code already exists")
// 	}

// }

// func TestGetAllPulsaPaketDataSuccess(t *testing.T) {
// 	var provider string
// 	mockPPD := make([]model.PulsaPaketData, 0)
// 	mockPPD = append(mockPPD, model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	})

// 	mockPayload := dto.PulsaDto{
// 		Type:        "pulsa",
// 		Provider:    "Telkomsel",
// 		PhoneNumber: "081323",
// 	}

// 	isUser := false

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("GetAllPulsaPaketData", mockPayload, &isUser).Return(mockPPD, nil)

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	if provider != "" {
// 		mockPayload.Provider = provider
// 	}

// 	result, err := service.GetAllPulsaPaketData(mockPayload, &isUser)

// 	if err != nil {
// 		t.Errorf("Got Error %v", err)
// 	}
// 	assert.Equal(t, result[0].Name, mockPPD[0].Name)
// 	assert.Equal(t, result[0].Code, mockPPD[0].Code)
// 	assert.Equal(t, result[0].Price, mockPPD[0].Price)
// 	assert.Equal(t, result[0].Provider, mockPPD[0].Provider)
// }

// func TestGetAllPulsaPaketDataError2(t *testing.T) {
// 	mockPPD := make([]model.PulsaPaketData, 0)
// 	mockPPD = append(mockPPD, model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	})

// 	mockPayload := make([]dto.PulsaDto, 0)
// 	mockPayload = append(mockPayload,
// 		dto.PulsaDto{
// 			Type:        "pulsa",
// 			Provider:    "",
// 			PhoneNumber: "",
// 		})
// 	mockPayload = append(mockPayload, dto.PulsaDto{
// 		Type:        "pulsa",
// 		Provider:    "Telkomsel",
// 		PhoneNumber: "081323",
// 	})

// 	for _, v := range mockPayload {
// 		if v.Provider == "" {
// 			isUser := true

// 			mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 			mockUserRepo := mocks.NewUserRepository(t)
// 			mockTransactionRepo := mocks.NewTransactionRepository(t)
// 			mockDiscountRepo := mocks.NewDiscountRepository(t)
// 			service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 			_, err := service.GetAllPulsaPaketData(v, &isUser)

// 			if err != nil {
// 				assert.Error(t, err, "failed get all ppd")
// 			}

// 		} else {
// 			isUser := true

// 			mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 			mockUserRepo := mocks.NewUserRepository(t)
// 			mockTransactionRepo := mocks.NewTransactionRepository(t)
// 			mockDiscountRepo := mocks.NewDiscountRepository(t)
// 			mockPPDRepository.On("GetAllPulsaPaketData", v, &isUser).Return(mockPPD, errors.New("failed get all ppd"))

// 			service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 			_, err := service.GetAllPulsaPaketData(v, &isUser)

// 			if err != nil {
// 				assert.Error(t, err, "failed get all ppd")
// 			}
// 		}

// 	}

// }

// func TestGetProviderByPhone(t *testing.T) {
// 	// Test cases
// 	testCases := []struct {
// 		phone    string
// 		expected string
// 	}{
// 		{"0852xxxxxxxx", "AS"},
// 		{"0853xxxxxxxx", "AS"},
// 		{"0823xxxxxxxx", "AS"},
// 		{"0851xxxxxxxx", "AS"},
// 		{"0811xxxxxxxx", "Halo"},
// 		{"0812xxxxxxxx", "Telkomsel"},
// 		{"0813xxxxxxxx", "Telkomsel"},
// 		{"0821xxxxxxxx", "Telkomsel"},
// 		{"0822xxxxxxxx", "Telkomsel"},
// 		{"0814xxxxxxxx", "Indosat"},
// 		{"0815xxxxxxxx", "Indosat"},
// 		{"0816xxxxxxxx", "Indosat"},
// 		{"0855xxxxxxxx", "Indosat"},
// 		{"0856xxxxxxxx", "Indosat"},
// 		{"0857xxxxxxxx", "Indosat"},
// 		{"0858xxxxxxxx", "Indosat"},
// 		{"0817xxxxxxxx", "XL"},
// 		{"0818xxxxxxxx", "XL"},
// 		{"0859xxxxxxxx", "XL"},
// 		{"0877xxxxxxxx", "XL"},
// 		{"0878xxxxxxxx", "XL"},
// 		{"0838xxxxxxxx", "Axis"},
// 		{"0831xxxxxxxx", "Axis"},
// 		{"0832xxxxxxxx", "Axis"},
// 		{"0833xxxxxxxx", "Axis"},
// 		{"0895xxxxxxxx", "Three"},
// 		{"0896xxxxxxxx", "Three"},
// 		{"0897xxxxxxxx", "Three"},
// 		{"0898xxxxxxxx", "Three"},
// 		{"0899xxxxxxxx", "Three"},
// 		{"0881xxxxxxxx", "Smatfren"},
// 		{"0882xxxxxxxx", "Smatfren"},
// 		{"0883xxxxxxxx", "Smatfren"},
// 		{"0884xxxxxxxx", "Smatfren"},
// 		{"0885xxxxxxxx", "Smatfren"},
// 		{"0886xxxxxxxx", "Smatfren"},
// 		{"0887xxxxxxxx", "Smatfren"},
// 		{"0888xxxxxxxx", "Smatfren"},
// 		{"0889xxxxxxxx", "Smatfren"},
// 		{"", ""},    // Empty phone case
// 		{"123", ""}, // Invalid phone case
// 		{"023473ß", ""},
// 	}

// 	// Iterate over test cases
// 	for _, tc := range testCases {
// 		result := getProviderByPhone(tc.phone)
// 		if result != tc.expected {
// 			t.Errorf("Phone: %s, Expected: %s, Got: %s", tc.phone, tc.expected, result)
// 		}
// 	}
// }

// func TestGetPulsaPaketDataByIdSuccess(t *testing.T) {
// 	mockPPD := model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	}

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("GetPulsaPaketDataById", mockID).Return(mockPPD, nil)

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	result, err := service.GetPulsaPaketDataById(mockID)

// 	if err != nil {
// 		t.Errorf("Got Error %v", err)
// 	}
// 	assert.Equal(t, result.Name, mockPPD.Name)
// 	assert.Equal(t, result.Code, mockPPD.Code)
// 	assert.Equal(t, result.Price, mockPPD.Price)
// 	assert.Equal(t, result.Provider, mockPPD.Provider)
// }

// func TestGetPulsaPaketDataByIdError(t *testing.T) {

// 	mockPPD := model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	}

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("GetPulsaPaketDataById", mockID).Return(mockPPD, errors.New("not found"))

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	_, err := service.GetPulsaPaketDataById(mockID)

// 	if err != nil {
// 		assert.Error(t, err, "not found")
// 	}

// }

// func TestUpdatePulsaByIdSuccess(t *testing.T) {
// 	mockPPD := model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	}

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("UpdatePulsaById", mockID, mockPPD).Return(nil)
// 	mockPPDRepository.On("GetPulsaPaketDataById", mockID).Return(mockPPD, nil)

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	result, err := service.UpdatePulsaById(mockID, mockPPD)
// 	if err != nil {
// 		t.Errorf("Got Error %v", err)
// 	}

// 	assert.Equal(t, result.Name, mockPPD.Name)
// 	assert.Equal(t, result.Code, mockPPD.Code)
// 	assert.Equal(t, result.Price, mockPPD.Price)
// 	assert.Equal(t, result.Provider, mockPPD.Provider)
// }

// func TestUpdatePulsaByIdError(t *testing.T) {

// 	mockPPD := model.PulsaPaketData{
// 		Name:     "Pulsa 10000",
// 		Price:    11000,
// 		Code:     "PSTS10",
// 		Type:     model.PULSA_TYPE,
// 		Provider: "Telkomsel",
// 	}

// 	testCase := []struct {
// 		newCase  string
// 		expected string
// 	}{
// 		{"UpdatePulsaById", "not found"},
// 		{"GetPulsaPaketDataById", "not found"},
// 	}

// 	for _, v := range testCase {
// 		if v.newCase == "UpdatePulsaById" {
// 			mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 			mockUserRepo := mocks.NewUserRepository(t)
// 			mockTransactionRepo := mocks.NewTransactionRepository(t)
// 			mockDiscountRepo := mocks.NewDiscountRepository(t)
// 			mockPPDRepository.On("UpdatePulsaById", mockID, mockPPD).Return(errors.New("not found"))

// 			service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 			_, err := service.UpdatePulsaById(mockID, mockPPD)

// 			if err != nil {
// 				assert.Error(t, err, "not found")

// 			}
// 		} else if v.newCase == "GetPulsaPaketDataById" {
// 			mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 			mockUserRepo := mocks.NewUserRepository(t)
// 			mockTransactionRepo := mocks.NewTransactionRepository(t)
// 			mockDiscountRepo := mocks.NewDiscountRepository(t)

// 			mockPPDRepository.On("UpdatePulsaById", mockID, mockPPD).Return(nil)
// 			mockPPDRepository.On("GetPulsaPaketDataById", mockID).Return(mockPPD, errors.New("failed"))

// 			service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)
// 			_, err := service.UpdatePulsaById(mockID, mockPPD)

// 			if err != nil {
// 				assert.Error(t, err, "not found")
// 			}

// 		}
// 	}

// }

// func TestDeletePulsaByIdSuccess(t *testing.T) {

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("DeletePulsaById", mockID).Return(nil)

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	err := service.DeletePulsaById(mockID)

// 	if err != nil {
// 		t.Errorf("Got Error %v", err)
// 	}

// }

// func TestDeletePulsaByIdError(t *testing.T) {

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("DeletePulsaById", mockID).Return(errors.New("not found"))

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	err := service.DeletePulsaById(mockID)

// 	if err != nil {
// 		assert.Error(t, err, "not found")
// 	}

// }

// func TestCreateTransactionPPDError(t *testing.T) {
// 	userID := uuid.New().String()
// 	payload := dto.TransactionPPDDto{
// 		ProductID: "sadsad",
// 	}

// 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// 	mockUserRepo := mocks.NewUserRepository(t)
// 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// 	mockPPDRepository.On("GetPulsaPaketDataById", payload.ProductID).Return(model.PulsaPaketData{}, errors.New("not found"))

// 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// 	_, err := service.CreateTransactionPPD(userID, payload)

// 	if err != nil {
// 		assert.Error(t, err, "not found")
// 	}
// }

// // func TestCreateTransactionPPDSuccess(t *testing.T) {
// // 	userID := uuid.New().String()
// // 	ppdID := uuid.New().String()
// // 	// trID := uuid.New().String()
// // 	discountID := uuid.New().String()

// // 	mockPayload := dto.TransactionPPDDto{
// // 		Type:        model.PULSA_TYPE,
// // 		ProductID:   ppdID,
// // 		DiscountID:  discountID,
// // 		PhoneNumber: "081323",
// // 	}

// // 	mockPPD := model.PulsaPaketData{
// // 		Name:     "Pulsa 10000",
// // 		Price:    11000,
// // 		Code:     "PSTS10",
// // 		Type:     model.PULSA_TYPE,
// // 		Provider: "Telkomsel",
// // 	}

// // 	mockTransaction := &model.Transaction{
// // 		ID:          uuid.New().String(),
// // 		UserID:      userID,
// // 		Status:      model.STATUS_SUCCESSFUL,
// // 		ProductType: model.PULSA_TYPE,
// // 		ProductDetail: model.TransactionPPD{
// // 			Phone:    "081323",
// // 			Name:     "Pulsa 10000",
// // 			Provider: "Telkomsel",
// // 			Code:     "PSTS10",
// // 			// DiscountID: discountID,
// // 		},
// // 		Price:         11000,
// // 		Description:   "Pembelian Pulsa Paket Data  pulsa ",
// // 		AdminFee:      2500,
// // 		DiscountPrice: 0,
// // 		TotalPrice:    13500,
// // 	}

// // 	mockUser := &model.User{
// // 		UserType: model.USER_TYPE,
// // 		Amount:   10000000,
// // 	}

// // 	mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// // 	mockUserRepo := mocks.NewUserRepository(t)
// // 	mockTransactionRepo := mocks.NewTransactionRepository(t)
// // 	mockDiscountRepo := mocks.NewDiscountRepository(t)

// // 	mockPPDRepository.On("GetPulsaPaketDataById", mockPayload.ProductID).Return(mockPPD, nil)
// // 	mockDiscountRepo.On("GetDiscountByIdRepository", discountID).Return(&model.Discount{}, nil)
// // 	mockUserRepo.On("GetUserByIDRepository", userID).Return(mockUser, nil)
// // 	mockTransactionRepo.On("CreateTransactionByUserIdRepository", mockTransaction).Return(mockTransaction, nil)
// // 	mockUser.Amount -= mockTransaction.TotalPrice
// // 	mockUserRepo.On("UpdateUserAmountByIDRepository", mockUser.ID, mockUser).Return(mockUser, nil)

// // 	service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// // 	result, err := service.CreateTransactionPPD(userID, mockPayload)
// // 	if err != nil {
// // 		t.Errorf("Got Error %v", err)
// // 	}

// // 	assert.Equal(t, result.ProductType, mockTransaction.ProductType)
// // }

// // func TestUpdatePulsaByIdError(t *testing.T) {

// // 	mockPPD := model.PulsaPaketData{
// // 		Name:     "Pulsa 10000",
// // 		Price:    11000,
// // 		Code:     "PSTS10",
// // 		Type:     model.PULSA_TYPE,
// // 		Provider: "Telkomsel",
// // 	}

// // 	testCase := []struct {
// // 		newCase  string
// // 		expected string
// // 	}{
// // 		{"UpdatePulsaById", "not found"},
// // 		{"GetPulsaPaketDataById", "not found"},
// // 	}

// // 	for _, v := range testCase {
// // 		if v.newCase == "UpdatePulsaById" {
// // 			mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// // 			mockUserRepo := mocks.NewUserRepository(t)
// // 			mockTransactionRepo := mocks.NewTransactionRepository(t)
// // 			mockDiscountRepo := mocks.NewDiscountRepository(t)
// // 			mockPPDRepository.On("UpdatePulsaById", mockID, mockPPD).Return(errors.New("not found"))

// // 			service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)

// // 			_, err := service.UpdatePulsaById(mockID, mockPPD)

// // 			if err != nil {
// // 				assert.Error(t, err, "not found")

// // 			}
// // 		} else if v.newCase == "GetPulsaPaketDataById" {
// // 			mockPPDRepository := mocks.NewPulsaPaketDataRepository(t)
// // 			mockUserRepo := mocks.NewUserRepository(t)
// // 			mockTransactionRepo := mocks.NewTransactionRepository(t)
// // 			mockDiscountRepo := mocks.NewDiscountRepository(t)

// // 			mockPPDRepository.On("UpdatePulsaById", mockID, mockPPD).Return(nil)
// // 			mockPPDRepository.On("GetPulsaPaketDataById", mockID).Return(mockPPD, errors.New("failed"))

// // 			service := NewPulsaPaketDataUsecase(mockPPDRepository, mockUserRepo, mockTransactionRepo, mockDiscountRepo)
// // 			_, err := service.UpdatePulsaById(mockID, mockPPD)

// // 			if err != nil {
// // 				assert.Error(t, err, "not found")
// // 			}

// // 		}
// // 	}

// // }
