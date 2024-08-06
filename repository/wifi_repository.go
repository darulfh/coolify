package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type WifiRepository interface {
	CreateWifiRepository(wifi *model.Wifi) (*model.Wifi, error)
	GetAllWifiRepository(page, limit int) ([]*model.Wifi, error)
	GetWifiByIDRepository(id string) (*model.Wifi, error)
	GetWifiByCodeRepository(code string) (*model.Wifi, error)
	UpdateWifiByIDRepository(id string, wifi *model.Wifi) (*model.Wifi, error)
	DeleteWifiByIDRepository(id string) error
}

type wifiRepository struct {
	db *gorm.DB
}

func NewWifiRepository(db *gorm.DB) *wifiRepository {
	return &wifiRepository{db}
}

func (r *wifiRepository) CreateWifiRepository(wifi *model.Wifi) (*model.Wifi, error) {
	result := r.db.Create(wifi)
	if result.Error != nil {
		return nil, errors.New("failed to create WiFi")
	}

	return wifi, nil
}

func (r *wifiRepository) GetAllWifiRepository(page, limit int) ([]*model.Wifi, error) {
	var wifis []*model.Wifi
	offset := (page - 1) * limit
	result := r.db.Offset(offset).Limit(limit).Find(&wifis)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting wifis: %s", result.Error)
	}
	return wifis, nil
}

func (r *wifiRepository) GetWifiByIDRepository(id string) (*model.Wifi, error) {
	var wifi model.Wifi
	result := r.db.First(&wifi, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("wifi with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting wifi with ID %s: %s", id, result.Error)
	}
	return &wifi, nil
}

func (r *wifiRepository) GetWifiByCodeRepository(code string) (*model.Wifi, error) {
	var wifi model.Wifi
	result := r.db.First(&wifi, "code = ?", code)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("wifi with code %s not found", code)
		}
		return nil, fmt.Errorf("error getting wifi with code %s: %s", code, result.Error)
	}
	return &wifi, nil
}

func (r *wifiRepository) UpdateWifiByIDRepository(id string, wifi *model.Wifi) (*model.Wifi, error) {
	result := r.db.Model(&model.Wifi{}).Where("id = ?", id).Updates(wifi)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("wifi not found")
	}
	return wifi, nil
}

func (r *wifiRepository) DeleteWifiByIDRepository(id string) error {
	result := r.db.Delete(&model.Wifi{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("wifi not found")
	}
	return nil
}
