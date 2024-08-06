package transaction

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	GetAllTransactionsUseCase(page, limit int) ([]*model.Transaction, error)
	GetTransactionByIdUseCase(id string) (*model.Transaction, error)
	GetTransactionByUserIdUseCase(userID, productType string, page, limit int) ([]*model.Transaction, error)
	GetTransactionProductTypeUseCase(product, status string, page, limit int) ([]*model.Transaction, error)
	GetTransactionQueryUseCase(query string, page, limit int) ([]*model.Transaction, error)
	GetTransactionStatusQueryUseCase(query, status string, page, limit int) ([]*model.Transaction, error)
	GetTransactionsPriceCountUseCase() ([]model.TransactionCountInfo, error)
	GetTransactionsPriceByMonthUseCase() ([]model.TransactionCountInfo, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) *transactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (uc *transactionUsecase) GetAllTransactionsUseCase(page, limit int) ([]*model.Transaction, error) {
	transactions, err := uc.transactionRepository.GetAllTransactionsRepository(page, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
func (uc *transactionUsecase) GetTransactionByIdUseCase(id string) (*model.Transaction, error) {
	transactions, err := uc.transactionRepository.GetTransactionByIdRepository(id)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (uc *transactionUsecase) GetTransactionByUserIdUseCase(userID, productType string, page, limit int) ([]*model.Transaction, error) {
	transactions, err := uc.transactionRepository.GetTransactionByUserIdRepository(userID, productType, page, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (uc *transactionUsecase) GetTransactionProductTypeUseCase(product, status string, page, limit int) ([]*model.Transaction, error) {
	lowercaseStatus := strings.ToLower(status)
	transactions, err := uc.transactionRepository.GetTransactionsProductTypeRepository(product, lowercaseStatus, page, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
func (uc *transactionUsecase) GetTransactionQueryUseCase(query string, page, limit int) ([]*model.Transaction, error) {
	lowercaseQuery := strings.ToLower(query)
	transactions, err := uc.transactionRepository.GetTransactionsByQueryRepository(lowercaseQuery, page, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (uc *transactionUsecase) GetTransactionStatusQueryUseCase(query, status string, page, limit int) ([]*model.Transaction, error) {
	lowercaseQuery := strings.ToLower(query)
	lowercaseStatus := strings.ToLower(status)
	transactions, err := uc.transactionRepository.GetTransactionsByStatusQueryRepository(lowercaseQuery, lowercaseStatus, page, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (uc *transactionUsecase) GetTransactionsPriceCountUseCase() ([]model.TransactionCountInfo, error) {
	transactions, err := uc.transactionRepository.GetTransactionsPriceCountRepository()
	if err != nil {
		return nil, err
	}

	totalPriceByProductType := make(map[string]float64)

	for _, transaction := range transactions {
		productType := transaction.ProductType
		price := transaction.Price

		totalPriceByProductType[productType] += price
	}

	data := make([]model.TransactionCountInfo, 0)

	for productType, price := range totalPriceByProductType {
		productInfo := model.TransactionCountInfo{
			ID:      fmt.Sprintf("%s-%s", productType, uuid.New().String()),
			Product: productType,
			Price:   price,
		}
		data = append(data, productInfo)
	}

	return data, nil
}

func (uc *transactionUsecase) GetTransactionsPriceByMonthUseCase() ([]model.TransactionCountInfo, error) {
	currentTime := time.Now()
	currentYear := currentTime.Year()

	transactionsByMonth := make(map[time.Month]float64)

	for month := time.January; month <= time.December; month++ {
		transactions, err := uc.transactionRepository.GetTransactionsByMonthRepository(month, currentYear)
		if err != nil {
			return nil, err
		}

		totalPrice := 0.0

		for _, transaction := range transactions {
			totalPrice += transaction.Price
		}

		transactionsByMonth[month] = totalPrice
	}

	data := make([]model.TransactionCountInfo, 0)

	for month, totalPrice := range transactionsByMonth {
		transactionData := model.TransactionCountInfo{
			ID:    fmt.Sprintf("%d%02d", currentYear, int(month)),
			Price: totalPrice,
			Month: month.String(),
			Year:  currentYear,
		}
		data = append(data, transactionData)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	return data, nil
}
