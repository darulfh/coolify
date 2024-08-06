package bank

import (
	"errors"
	"fmt"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"
)

type BankUseCase interface {
	CreateBankUseCase(bank *model.Bank) (*model.Bank, error)
	GetAllBanksUseCase(page, limit int) ([]*model.Bank, error)
	GetBankByIdUseCase(bankID string) (*model.Bank, error)
	UpdateBankByIdUseCase(bankId string, bank *model.Bank) (*model.Bank, error)
	DeleteBankByIdUseCase(bankID string) error
}

type bankUseCase struct {
	bankRepository repository.BankRepository
}

func NewbankUseCase(bankRepository repository.BankRepository) *bankUseCase {
	return &bankUseCase{
		bankRepository: bankRepository,
	}
}

func (uc *bankUseCase) CreateBankUseCase(payload *model.Bank) (*model.Bank, error) {

	bank, err := uc.bankRepository.CreateBankRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("error creating bank in database: %w", err)
	}
	return bank, err

}

func (uc *bankUseCase) GetAllBanksUseCase(page, limit int) ([]*model.Bank, error) {
	banks, err := uc.bankRepository.GetAllBanksRepository(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all banks: %v", err)
	}
	return banks, nil
}

func (uc *bankUseCase) GetBankByIdUseCase(bankID string) (*model.Bank, error) {

	bank, err := uc.bankRepository.GetBankByIdRepository(bankID)
	if err != nil {
		return nil, errors.New("bank not found")
	}

	return bank, nil

}

func (uc *bankUseCase) UpdateBankByIdUseCase(bankId string, payload *model.Bank) (*model.Bank, error) {
	bank, err := uc.bankRepository.GetBankByIdRepository(bankId)
	if err != nil {
		return nil, fmt.Errorf("failed to update bank: %v", err)
	}

	bank.Name = payload.Name
	bank.Image = payload.Image
	bank.BankCode = payload.BankCode
	bank.UpdatedAt = time.Now()

	updatedbank, err := uc.bankRepository.UpdateBankByIdRepository(bankId, bank)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return updatedbank, nil
}

func (uc *bankUseCase) DeleteBankByIdUseCase(bankID string) error {
	err := uc.bankRepository.DeleteBankByIdRepository(bankID)
	if err != nil {
		return errors.New("bank not found")
	}
	return err
}
