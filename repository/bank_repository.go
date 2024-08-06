package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type BankRepository interface {
	CreateBankRepository(bank *model.Bank) (*model.Bank, error)
	GetBankByIdRepository(bankId string) (*model.Bank, error)
	GetAllBanksRepository(page, limit int) ([]*model.Bank, error)
	UpdateBankByIdRepository(bankId string, bank *model.Bank) (*model.Bank, error)
	DeleteBankByIdRepository(bankId string) error
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) *bankRepository {
	return &bankRepository{db}
}

func (r *bankRepository) CreateBankRepository(bank *model.Bank) (*model.Bank, error) {
	result := r.db.Create(bank)
	if result.Error != nil {
		return nil, errors.New("failed to create bank")
	}

	return bank, nil
}

func (r *bankRepository) GetBankByIdRepository(bankId string) (*model.Bank, error) {
	var bank model.Bank
	result := r.db.First(&bank, "id = ?", bankId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("bank with ID %s not found", bankId)
		}
		return nil, fmt.Errorf("error getting Bank with ID %s: %s", bankId, result.Error)
	}

	return &bank, nil
}

func (r *bankRepository) GetAllBanksRepository(page, limit int) ([]*model.Bank, error) {
	var banks []*model.Bank
	offset := (page - 1) * limit
	result := r.db.Offset(offset).Limit(limit).Find(&banks)
	if result.Error != nil {
		return nil, errors.New("failed to get banks")
	}

	return banks, nil
}

func (r *bankRepository) UpdateBankByIdRepository(bankId string, bank *model.Bank) (*model.Bank, error) {
	result := r.db.Model(&model.Bank{}).Where("id = ?", bankId).Updates(bank)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("bank not found")
	}

	return bank, nil
}

func (r *bankRepository) DeleteBankByIdRepository(bankId string) error {
	result := r.db.Delete(&model.Bank{}, "id = ?", bankId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("bank not found")
	}
	return nil
}
