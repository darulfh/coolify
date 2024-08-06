package transaction_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository/mocks"
	"github.com/darulfh/skuy_pay_be/usecase/transaction"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionUseCaseTest struct {
	suite.Suite
	transactionUseCase  transaction.TransactionUsecase
	transactionRepoMock *mocks.TransactionRepository
}

func TestTransactionsTest(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseTest))
}

func (s *TransactionUseCaseTest) SetupTest() {
	s.transactionRepoMock = &mocks.TransactionRepository{}
	s.transactionUseCase = transaction.NewTransactionUsecase(s.transactionRepoMock)
}
func (s *TransactionUseCaseTest) TestGetAllTransactionsUseCase_Success() {
	mockTransactions := []*model.Transaction{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}

	s.transactionRepoMock.On("GetAllTransactionsRepository", 1, 10).Return(mockTransactions, nil)

	resp, err := s.transactionUseCase.GetAllTransactionsUseCase(1, 10)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "2", resp[1].ID)
}

func (s *TransactionUseCaseTest) TestGetAllTransactionsUseCase_Error() {
	expectedErr := fmt.Errorf("error retrieving transactions")

	s.transactionRepoMock.On("GetAllTransactionsRepository", 1, 10).Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetAllTransactionsUseCase(1, 10)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionByIdUseCase_Success() {
	expectedTransaction := &model.Transaction{
		ID: "1",
	}

	s.transactionRepoMock.On("GetTransactionByIdRepository", "1").Return(expectedTransaction, nil)

	resp, err := s.transactionUseCase.GetTransactionByIdUseCase("1")
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedTransaction, resp)
}

func (s *TransactionUseCaseTest) TestGetTransactionByIdUseCase_Error() {
	expectedErr := fmt.Errorf("error retrieving transaction")

	s.transactionRepoMock.On("GetTransactionByIdRepository", "1").Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionByIdUseCase("1")
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionByUserIdUseCase_Success() {
	mockTransactions := []*model.Transaction{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}

	s.transactionRepoMock.On("GetTransactionByUserIdRepository", "user123", "type1", 1, 10).Return(mockTransactions, nil)

	resp, err := s.transactionUseCase.GetTransactionByUserIdUseCase("user123", "type1", 1, 10)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "2", resp[1].ID)
}

func (s *TransactionUseCaseTest) TestGetTransactionByUserIdUseCase_Error() {
	expectedErr := fmt.Errorf("error retrieving transactions for user")

	s.transactionRepoMock.On("GetTransactionByUserIdRepository", "user123", "type1", 1, 10).Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionByUserIdUseCase("user123", "type1", 1, 10)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionProductTypeUseCase_Success() {
	product := "product1"
	status := "active"
	page := 1
	limit := 10

	mockTransactions := []*model.Transaction{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}

	s.transactionRepoMock.On("GetTransactionsProductTypeRepository", product, "active", page, limit).Return(mockTransactions, nil)

	resp, err := s.transactionUseCase.GetTransactionProductTypeUseCase(product, status, page, limit)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "2", resp[1].ID)
}

func (s *TransactionUseCaseTest) TestGetTransactionProductTypeUseCase_Error() {
	product := "product1"
	status := "active"
	page := 1
	limit := 10
	expectedErr := fmt.Errorf("error retrieving transactions by product type")

	s.transactionRepoMock.On("GetTransactionsProductTypeRepository", product, "active", page, limit).Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionProductTypeUseCase(product, status, page, limit)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionQueryUseCase_Success() {
	query := "search"
	page := 1
	limit := 10

	mockTransactions := []*model.Transaction{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}

	s.transactionRepoMock.On("GetTransactionsByQueryRepository", "search", page, limit).Return(mockTransactions, nil)

	resp, err := s.transactionUseCase.GetTransactionQueryUseCase(query, page, limit)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "2", resp[1].ID)
}

func (s *TransactionUseCaseTest) TestGetTransactionQueryUseCase_Error() {
	query := "search"
	page := 1
	limit := 10
	expectedErr := fmt.Errorf("error retrieving transactions by query")

	s.transactionRepoMock.On("GetTransactionsByQueryRepository", "search", page, limit).Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionQueryUseCase(query, page, limit)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionStatusQueryUseCase_Success() {
	query := "search"
	status := "active"
	page := 1
	limit := 10

	mockTransactions := []*model.Transaction{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}

	s.transactionRepoMock.On("GetTransactionsByStatusQueryRepository", "search", "active", page, limit).Return(mockTransactions, nil)

	resp, err := s.transactionUseCase.GetTransactionStatusQueryUseCase(query, status, page, limit)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "1", resp[0].ID)
	assert.Equal(s.T(), "2", resp[1].ID)
}

func (s *TransactionUseCaseTest) TestGetTransactionStatusQueryUseCase_Error() {
	query := "search"
	status := "active"
	page := 1
	limit := 10
	expectedErr := fmt.Errorf("error retrieving transactions by status and query")

	s.transactionRepoMock.On("GetTransactionsByStatusQueryRepository", "search", "active", page, limit).Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionStatusQueryUseCase(query, status, page, limit)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionsPriceCountUseCase_Success() {
	mockTransactions := []*model.Transaction{
		{
			ProductType: "A",
			Price:       10.0,
		},
		{
			ProductType: "B",
			Price:       20.0,
		},
		{
			ProductType: "A",
			Price:       15.0,
		},
	}

	s.transactionRepoMock.On("GetTransactionsPriceCountRepository").Return(mockTransactions, nil)

	resp, err := s.transactionUseCase.GetTransactionsPriceCountUseCase()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 2, len(resp))
	assert.Equal(s.T(), "A", resp[0].Product)
	assert.Equal(s.T(), 25.0, resp[0].Price)
	assert.Equal(s.T(), "B", resp[1].Product)
	assert.Equal(s.T(), 20.0, resp[1].Price)
}

func (s *TransactionUseCaseTest) TestGetTransactionsPriceCountUseCase_Error() {
	expectedErr := fmt.Errorf("error retrieving transactions for price count")

	s.transactionRepoMock.On("GetTransactionsPriceCountRepository").Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionsPriceCountUseCase()
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}

func (s *TransactionUseCaseTest) TestGetTransactionsPriceByMonthUseCase_Error() {
	expectedErr := fmt.Errorf("error retrieving transactions for price by month")
	currentYear := time.Now().Year()

	s.transactionRepoMock.On("GetTransactionsByMonthRepository", time.January, currentYear).Return(nil, expectedErr)

	resp, err := s.transactionUseCase.GetTransactionsPriceByMonthUseCase()
	assert.Error(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.EqualError(s.T(), err, expectedErr.Error())
}
