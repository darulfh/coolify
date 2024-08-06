package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type InsuranceRepository interface {
	CreateInsuranceRepository(insurance *model.Insurance) (*model.Insurance, error)
	GetInsuranceyIdRepository(insuranceId string) (*model.Insurance, error)
	GetInsuranceByPeriodRepository(customerID, period string) (*model.Insurance, error)
	GetAllInsuranceRepository(page, limit int) ([]*model.Insurance, error)
	UpdateInsuranceByIdRepository(insuranceId string, insurance *model.Insurance) (*model.Insurance, error)
	DeleteInsuranceByIdRepository(insuranceId string) error
}

type insuranceRepository struct {
	db *gorm.DB
}

func NewInsuranceRepository(db *gorm.DB) *insuranceRepository {
	return &insuranceRepository{db}
}

func (r *insuranceRepository) CreateInsuranceRepository(insurance *model.Insurance) (*model.Insurance, error) {
	result := r.db.Create(insurance)
	if result.Error != nil {
		return nil, errors.New("failed to Create insurance")
	}

	return insurance, nil
}

func (r *insuranceRepository) GetInsuranceyIdRepository(insuranceId string) (*model.Insurance, error) {

	var insurance model.Insurance

	result := r.db.First(&insurance, "id = ?", insuranceId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("insurance with ID %s not found", insuranceId)
		}
		return nil, fmt.Errorf("error getting insurance with ID %s: %s", insuranceId, result.Error)
	}

	return &insurance, nil
}

func (r *insuranceRepository) GetAllInsuranceRepository(page, limit int) ([]*model.Insurance, error) {
	var insurance []*model.Insurance
	offSet := (page - 1) * limit
	result := r.db.Offset(offSet).Limit(limit).Find(&insurance)
	if result.Error != nil {
		return nil, errors.New("failed to get insurance")
	}

	return insurance, nil
}

func (r *insuranceRepository) GetInsuranceByPeriodRepository(customerID, period string) (*model.Insurance, error) {
	var insurance model.Insurance

	result := r.db.First(&insurance, "customer_id = ? AND period = ?", customerID, period)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting insurance for customer ID %s and period %s: %w", customerID, period, result.Error)
	}

	return &insurance, nil
}

func (r *insuranceRepository) UpdateInsuranceByIdRepository(insuranceId string, insurance *model.Insurance) (*model.Insurance, error) {

	result := r.db.Model(&model.Insurance{}).Where("id = ?", insuranceId).Updates(insurance)
	if result.Error != nil {
		return nil, errors.New("failed to Update insurance")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("insurance Not Found")
	}

	return insurance, nil
}

func (r *insuranceRepository) DeleteInsuranceByIdRepository(insuranceId string) error {

	result := r.db.Delete((&model.Insurance{}), "id = ?", insuranceId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insurance Not Found")
	}
	return nil
}
