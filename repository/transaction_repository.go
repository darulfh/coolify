package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransactionByUserIdRepository(transaction *model.Transaction) (*model.Transaction, error)
	GetAllTransactionsRepository(page, limit int) ([]*model.Transaction, error)
	GetTransactionByIdRepository(transactionID string) (*model.Transaction, error)
	GetTransactionByUserIdRepository(userID, productType string, page, limit int) ([]*model.Transaction, error)
	GetProductDetailsByPeriodAndCustomerID(payload model.GetProductDetail) (*model.Transaction, error)
	GetTransactionsProductTypeRepository(productType, status string, page, limit int) ([]*model.Transaction, error)
	UpdateTransactionByIdRepository(userId string, transaction *model.Transaction) (*model.Transaction, error)
	GetTransactionsByQueryRepository(query string, page, limit int) ([]*model.Transaction, error)
	GetTransactionsByStatusQueryRepository(query, status string, page, limit int) ([]*model.Transaction, error)
	GetTransactionsPriceCountRepository() ([]*model.Transaction, error)
	GetTransactionsByMonthRepository(month time.Month, year int) ([]*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransactionByUserIdRepository(transaction *model.Transaction) (*model.Transaction, error) {

	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetAllTransactionsRepository(page, limit int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	offset := (page - 1) * limit

	err := r.db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionByIdRepository(transactionID string) (*model.Transaction, error) {
	var transaction *model.Transaction
	err := r.db.Where("id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) GetTransactionByUserIdRepository(userID, productType string, page, limit int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	offset := (page - 1) * limit

	query := r.db.Where("user_id = ?", userID).Offset(offset).Limit(limit)

	if productType != "" {
		query = query.Where("product_type LIKE ?", "%"+productType+"%")
	}

	result := query.Order("created_at DESC").Find(&transactions)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting users: %s", result.Error)
	}

	return transactions, nil
}

func (r *transactionRepository) GetProductDetailsByPeriodAndCustomerID(payload model.GetProductDetail) (*model.Transaction, error) {
	query := r.db.Where("product_type = ? AND product_detail::jsonb ->>'period' = ? AND product_detail::jsonb  ->>'customer_id' = ?", payload.ProductId, payload.Period, payload.CustomerId)

	var transaction model.Transaction
	err := query.First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) GetTransactionsProductTypeRepository(productType, status string, page, limit int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	offset := (page - 1) * limit

	query := r.db.Offset(offset).Limit(limit)
	if status != "" {
		query = query.Where("status LIKE ?", "%"+status+"%")
	}
	if productType != "" {
		query = query.Where("product_type LIKE ?", "%"+productType+"%")
	}

	err := query.Order("created_at DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionsByQueryRepository(query string, page, limit int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	offset := (page - 1) * limit

	queryString := "%" + query + "%"
	dbQuery := r.db.Where("id LIKE ? OR status LIKE ? OR product_type LIKE ? OR CAST(total_price AS TEXT) LIKE ? OR product_detail::jsonb ->>'name' LIKE ? ", queryString, queryString, queryString, queryString, queryString).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&transactions)

	if dbQuery.Error != nil {
		return nil, dbQuery.Error
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionsByStatusQueryRepository(query, status string, page, limit int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	offset := (page - 1) * limit

	queryString := "%" + query + "%"

	dbQuery := r.db.Where("status = ? AND (id LIKE ? OR product_type LIKE ? OR CAST(total_price AS TEXT) LIKE ? OR product_detail::jsonb ->>'name' LIKE ?)", status, queryString, queryString, queryString, queryString).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&transactions)

	if dbQuery.Error != nil {
		return nil, dbQuery.Error
	}

	return transactions, nil
}

func (r *transactionRepository) UpdateTransactionByIdRepository(id string, transaction *model.Transaction) (*model.Transaction, error) {
	result := r.db.Model(&model.Transaction{}).Where("id = ?", id).Updates(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransactionsPriceCountRepository() ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	err := r.db.Where("status = ?", model.STATUS_SUCCESSFUL).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionsByMonthRepository(month time.Month, year int) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	err := r.db.Where("status = ? AND EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", model.STATUS_SUCCESSFUL, int(month), year).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
