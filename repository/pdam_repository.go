package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type PdamRepository interface {
	CreatePdamRepository(pdam *model.Pdam) (*model.Pdam, error)
	GetPdamByIdRepository(pdamId string) (*model.Pdam, error)
	GetAllPdamRepository(page, limit int) ([]*model.Pdam, error)
	UpdatePdamByIdRepository(pdamId string, pdam *model.Pdam) (*model.Pdam, error)
	DeletePdamByIdRepository(pdamId string) error
	GetPdamByPeriodRepository(customerID, period string) (*model.Pdam, error)
}

type pdamRepository struct {
	db *gorm.DB
}

func NewPdamRepository(db *gorm.DB) *pdamRepository {
	return &pdamRepository{db}
}

func (r *pdamRepository) CreatePdamRepository(pdam *model.Pdam) (*model.Pdam, error) {
	result := r.db.Create(pdam)
	if result.Error != nil {
		return nil, errors.New("failed to Create Pdam")
	}

	return pdam, nil
}

func (r *pdamRepository) GetPdamByIdRepository(pdamId string) (*model.Pdam, error) {

	var pdam model.Pdam

	result := r.db.First(&pdam, "id = ?", pdamId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("PDAM with ID %s not found", pdamId)
		}
		return nil, fmt.Errorf("error getting PDAM with ID %s: %s", pdamId, result.Error)
	}

	return &pdam, nil
}

func (r *pdamRepository) GetAllPdamRepository(page, limit int) ([]*model.Pdam, error) {
	var pdam []*model.Pdam
	offSet := (page - 1) * limit
	result := r.db.Offset(offSet).Limit(limit).Find(&pdam)
	if result.Error != nil {
		return nil, errors.New("failed to get PDAM")
	}

	return pdam, nil
}

func (r *pdamRepository) GetPdamByPeriodRepository(customerID, period string) (*model.Pdam, error) {
	var pdam model.Pdam

	result := r.db.First(&pdam, "customer_id = ? AND period = ?", customerID, period)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting PDAM for customer ID %s and period %s: %w", customerID, period, result.Error)
	}

	return &pdam, nil
}

func (r *pdamRepository) UpdatePdamByIdRepository(pdamId string, pdam *model.Pdam) (*model.Pdam, error) {

	result := r.db.Model(&model.Pdam{}).Where("id = ?", pdamId).Updates(pdam)
	if result.Error != nil {
		return nil, errors.New("failed to Update PDAM")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("pdam Not Found")
	}

	return pdam, nil
}

func (r *pdamRepository) DeletePdamByIdRepository(pdamId string) error {

	result := r.db.Delete((&model.Pdam{}), "id = ?", pdamId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("pdam Not Found")
	}
	return nil
}
