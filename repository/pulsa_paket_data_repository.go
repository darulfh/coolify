package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/dto"
	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type PulsaPaketDataRepository interface {
	CreatePulsaPaketData(data model.PulsaPaketData) (model.PulsaPaketData, error)
	GetAllPulsaPaketData(data dto.PulsaDto, isUser *bool) ([]model.PulsaPaketData, error)
	GetPulsaPaketDataById(id string) (model.PulsaPaketData, error)
	UpdatePulsaById(id string, data model.PulsaPaketData) error
	DeletePulsaById(id string) error
}

type pulsaPaketDataRepository struct {
	db *gorm.DB
}

func NewPulsaPaketDataRepository(db *gorm.DB) *pulsaPaketDataRepository {
	return &pulsaPaketDataRepository{db}
}

func (r *pulsaPaketDataRepository) CreatePulsaPaketData(data model.PulsaPaketData) (model.PulsaPaketData, error) {
	var count int64

	r.db.Model(&model.PulsaPaketData{}).Where("code = ?", data.Code).Count(&count)

	if count > 0 {
		return model.PulsaPaketData{}, errors.New("code already exists")

	}
	if err := r.db.Create(&data).Error; err != nil {
		return model.PulsaPaketData{}, err
	}

	return data, nil
}
func (r *pulsaPaketDataRepository) GetAllPulsaPaketData(data dto.PulsaDto, isUser *bool) ([]model.PulsaPaketData, error) {
	var ppd []model.PulsaPaketData
	offset := (data.Page - 1) * data.Limit
	if isUser != nil {
		if err := r.db.Where("provider = ? AND is_active = ?", data.Provider, true).Offset(offset).Limit(100).Find(&ppd).Error; err != nil {
			return ppd, fmt.Errorf("error getting %s: %s", data.Type, err)
		}
	} else {
		if err := r.db.Where("type Like ? AND provider LIKE ?", "%"+data.Type+"%", "%"+data.Provider+"%").Offset(offset).Limit(data.Limit).Find(&ppd).Error; err != nil {
			return ppd, fmt.Errorf("error getting %s: %s", data.Type, err)
		}
	}

	return ppd, nil
}

func (r *pulsaPaketDataRepository) GetPulsaPaketDataById(id string) (model.PulsaPaketData, error) {
	var ppd model.PulsaPaketData

	result := r.db.First(&ppd, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ppd, fmt.Errorf("pulsa or paket data with ID %s not found", id)
		}
		return ppd, fmt.Errorf("error getting Pulsa or Paket Data with ID %s: %s", id, result.Error)
	}

	return ppd, nil
}

func (r *pulsaPaketDataRepository) UpdatePulsaById(id string, data model.PulsaPaketData) error {
	fmt.Printf("test::: %+v\n", data)

	result := r.db.Model(&model.PulsaPaketData{}).Where("id = ?", id).Updates(&data)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("pulsa not found")
	}

	return nil
}

func (r *pulsaPaketDataRepository) DeletePulsaById(id string) error {
	if err := r.db.Delete(&model.PulsaPaketData{}, "id = ?", id).Error; err != nil {
		return err

	}

	return nil
}
