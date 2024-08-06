package bank_test

import (
	"errors"
	"testing"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository/mocks"
	"github.com/darulfh/skuy_pay_be/usecase/bank"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BankUseCaseTest struct {
	suite.Suite
	bankUseCase bank.BankUseCase

	bankRepoMock *mocks.BankRepository
}

func TestBankUseCaseTest(t *testing.T) {
	suite.Run(t, new(BankUseCaseTest))
}

func (s *BankUseCaseTest) SetupTest() {
	s.bankRepoMock = &mocks.BankRepository{}
	s.bankUseCase = bank.NewbankUseCase(s.bankRepoMock)
}

func (s *BankUseCaseTest) TestGetAllBanks() {
	mockBanks := []*model.Bank{
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "1",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     "BRI",
			Image:    "url",
			BankCode: "002",
		},
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "2",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:     "BCA",
			Image:    "url",
			BankCode: "012",
		},
	}

	s.bankRepoMock.On("GetAllBanksRepository", 1, 10).Return(mockBanks, nil)

	resp, err := s.bankUseCase.GetAllBanksUseCase(1, 10)
	assert.NoError(s.T(), err)
	// assert.Equal(s.T(), "Success Get all banks", resp.)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "BRI", resp[0].Name)
	assert.Equal(s.T(), "url", resp[0].Image)
	assert.Equal(s.T(), "002", resp[0].BankCode)
	assert.Equal(s.T(), mockBanks[0].CreatedAt.Unix(), resp[0].CreatedAt.Unix())
	assert.Equal(s.T(), mockBanks[0].UpdatedAt.Unix(), resp[0].UpdatedAt.Unix())
}

func (s *BankUseCaseTest) TestGetAllBanks_Error() {
	s.bankRepoMock.On("GetAllBanksRepository", 1, 10).Return(nil, errors.New("repository error"))

	resp, err := s.bankUseCase.GetAllBanksUseCase(1, 10)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), "failed to get all banks: repository error", err.Error())
}

func (s *BankUseCaseTest) TestCreateBankUseCase_Success() {
	payload := &model.Bank{
		Name:     "BRI",
		Image:    "url",
		BankCode: "002",
	}

	expectedBank := &model.Bank{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "BRI",
		Image:    "url",
		BankCode: "002",
	}

	s.bankRepoMock.On("CreateBankRepository", payload).Return(expectedBank, nil)

	createdBank, err := s.bankUseCase.CreateBankUseCase(payload)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), createdBank)
	assert.Equal(s.T(), expectedBank.ID, createdBank.ID)
	assert.Equal(s.T(), expectedBank.Name, createdBank.Name)
	assert.Equal(s.T(), expectedBank.Image, createdBank.Image)
	assert.Equal(s.T(), expectedBank.BankCode, createdBank.BankCode)
	assert.Equal(s.T(), expectedBank.CreatedAt.Unix(), createdBank.CreatedAt.Unix())
	assert.Equal(s.T(), expectedBank.UpdatedAt.Unix(), createdBank.UpdatedAt.Unix())
}

func (s *BankUseCaseTest) TestCreateBankUseCase_Error() {
	payload := &model.Bank{
		Name:     "BRI",
		Image:    "url",
		BankCode: "002",
	}

	s.bankRepoMock.On("CreateBankRepository", payload).Return(nil, errors.New("repository error"))

	createdBank, err := s.bankUseCase.CreateBankUseCase(payload)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), createdBank)
	assert.EqualError(s.T(), err, "error creating bank in database: repository error")
}

func (s *BankUseCaseTest) TestGetBankByIdUseCase_Success() {
	bankID := "1"

	expectedBank := &model.Bank{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        bankID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "Test Bank",
		Image:    "test-url",
		BankCode: "001",
	}

	s.bankRepoMock.On("GetBankByIdRepository", bankID).Return(expectedBank, nil)

	bank, err := s.bankUseCase.GetBankByIdUseCase(bankID)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), bank)
	assert.Equal(s.T(), expectedBank.ID, bank.ID)
	assert.Equal(s.T(), expectedBank.Name, bank.Name)
	assert.Equal(s.T(), expectedBank.Image, bank.Image)
	assert.Equal(s.T(), expectedBank.BankCode, bank.BankCode)
	assert.Equal(s.T(), expectedBank.CreatedAt.Unix(), bank.CreatedAt.Unix())
	assert.Equal(s.T(), expectedBank.UpdatedAt.Unix(), bank.UpdatedAt.Unix())
}

func (s *BankUseCaseTest) TestGetBankByIdUseCase_Error() {
	bankID := "1"

	s.bankRepoMock.On("GetBankByIdRepository", bankID).Return(nil, errors.New("repository error"))

	bank, err := s.bankUseCase.GetBankByIdUseCase(bankID)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), bank)
	assert.EqualError(s.T(), err, "bank not found")
}

func (s *BankUseCaseTest) TestUpdateBankByIdUseCase_Success() {
	bankID := "1"
	payload := &model.Bank{
		Name:     "Updated Bank",
		Image:    "updated-url",
		BankCode: "003",
	}

	existingBank := &model.Bank{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        bankID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "Original Bank",
		Image:    "original-url",
		BankCode: "002",
	}

	expectedBank := &model.Bank{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        bankID,
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "Updated Bank",
		Image:    "updated-url",
		BankCode: "003",
	}

	s.bankRepoMock.On("GetBankByIdRepository", bankID).Return(existingBank, nil)
	s.bankRepoMock.On("UpdateBankByIdRepository", bankID, existingBank).Return(expectedBank, nil)

	updatedBank, err := s.bankUseCase.UpdateBankByIdUseCase(bankID, payload)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), updatedBank)
	assert.Equal(s.T(), expectedBank.ID, updatedBank.ID)
	assert.Equal(s.T(), expectedBank.Name, updatedBank.Name)
	assert.Equal(s.T(), expectedBank.Image, updatedBank.Image)
	assert.Equal(s.T(), expectedBank.BankCode, updatedBank.BankCode)
	assert.Equal(s.T(), expectedBank.CreatedAt.Unix(), updatedBank.CreatedAt.Unix())
	assert.Equal(s.T(), expectedBank.UpdatedAt.Unix(), updatedBank.UpdatedAt.Unix())
}

func (s *BankUseCaseTest) TestUpdateBankByIdUseCase_Error() {
	bankID := "1"
	payload := &model.Bank{
		Name:     "Updated Bank",
		Image:    "updated-url",
		BankCode: "003",
	}

	s.bankRepoMock.On("GetBankByIdRepository", bankID).Return(nil, errors.New("repository error"))

	updatedBank, err := s.bankUseCase.UpdateBankByIdUseCase(bankID, payload)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), updatedBank)
	assert.EqualError(s.T(), err, "failed to update bank: repository error")
}

func (s *BankUseCaseTest) TestDeleteBankByIdUseCase_Success() {
	bankID := "1"

	s.bankRepoMock.On("DeleteBankByIdRepository", bankID).Return(nil)

	err := s.bankUseCase.DeleteBankByIdUseCase(bankID)

	assert.NoError(s.T(), err)
}

func (s *BankUseCaseTest) TestDeleteBankByIdUseCase_Error() {
	bankID := "1"

	s.bankRepoMock.On("DeleteBankByIdRepository", bankID).Return(errors.New("repository error"))

	err := s.bankUseCase.DeleteBankByIdUseCase(bankID)

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "bank not found")
}
