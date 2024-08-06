package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type ElectricityRepository interface {
	CreateElectricityRepository(electricity *model.Electricity) (*model.Electricity, error)
	GetElectricityByIdRepository(electricityId string) (*model.Electricity, error)
	GetAllElectricityRepository(page, limit int) ([]*model.Electricity, error)
	GetElectricityByPeriodRepository(customerID, period string) (*model.Electricity, error)
	UpdateElectricityByIdRepository(electricityId string, electricity *model.Electricity) (*model.Electricity, error)
	DeleteElectricityByIdRepository(electricityId string) error
}

type electricityRepository struct {
	db *gorm.DB
}

func NewElectricityRepository(db *gorm.DB) *electricityRepository {
	return &electricityRepository{db}
}

func (r *electricityRepository) CreateElectricityRepository(electricity *model.Electricity) (*model.Electricity, error) {
	result := r.db.Create(electricity)
	if result.Error != nil {
		return nil, errors.New("failed to Create electricity")
	}

	return electricity, nil
}

func (r *electricityRepository) GetElectricityByIdRepository(electricityId string) (*model.Electricity, error) {

	var electricity model.Electricity

	result := r.db.First(&electricity, "id = ?", electricityId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("electricity with ID %s not found", electricityId)
		}
		return nil, fmt.Errorf("error getting electricity with ID %s: %s", electricityId, result.Error)
	}

	return &electricity, nil
}

func (r *electricityRepository) GetAllElectricityRepository(page, limit int) ([]*model.Electricity, error) {
	var electricity []*model.Electricity
	offSet := (page - 1) * limit
	result := r.db.Offset(offSet).Limit(limit).Find(&electricity)
	if result.Error != nil {
		return nil, errors.New("failed to get electricity")
	}

	return electricity, nil
}

func (r *electricityRepository) GetElectricityByPeriodRepository(customerID, period string) (*model.Electricity, error) {
	var electricity model.Electricity

	result := r.db.First(&electricity, "customer_id = ? AND period = ?", customerID, period)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting electricity for customer ID %s and period %s: %w", customerID, period, result.Error)
	}

	return &electricity, nil
}

func (r *electricityRepository) UpdateElectricityByIdRepository(electricityId string, electricity *model.Electricity) (*model.Electricity, error) {

	result := r.db.Model(&model.Electricity{}).Where("id = ?", electricityId).Updates(electricity)
	if result.Error != nil {
		return nil, errors.New("failed to Update electricity")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("electricity Not Found")
	}

	return electricity, nil
}

func (r *electricityRepository) DeleteElectricityByIdRepository(electricityId string) error {

	result := r.db.Delete((&model.Electricity{}), "id = ?", electricityId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("electricity Not Found")
	}
	return nil
}
