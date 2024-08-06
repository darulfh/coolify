package repository

import (
	"errors"
	"fmt"

	"github.com/darulfh/skuy_pay_be/model"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	CreateDiscountRepository(discount *model.Discount) (*model.Discount, error)
	GetDiscountByIdRepository(discountId string) (*model.Discount, error)
	GetDiscountByCodeRepository(discountCode string) (*model.Discount, error)
	GetAllDiscountRepository(page, limit int) ([]*model.Discount, error)
	UpdateDiscountByIdRepository(discountId string, discount *model.Discount) (*model.Discount, error)
	DeleteDiscountByIdRepository(discountId string) error
}

type discountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *discountRepository {
	return &discountRepository{db}
}

func (r *discountRepository) CreateDiscountRepository(discount *model.Discount) (*model.Discount, error) {
	result := r.db.First(&model.Discount{}, "discount_code = ?", discount.DiscountCode)
	if result.Error == nil {
		return nil, errors.New("discount with the same promo code already exists")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to check for existing discount: %s", result.Error)
	}

	result = r.db.Create(discount)
	if result.Error != nil {
		return nil, errors.New("failed to create discount")
	}

	return discount, nil
}

func (r *discountRepository) GetDiscountByIdRepository(discountId string) (*model.Discount, error) {

	var discount model.Discount

	result := r.db.First(&discount, "id = ?", discountId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &model.Discount{
				DiscountPrice: 0,
			}, nil
		}
		return nil, fmt.Errorf("error getting discount with ID %s: %s", discountId, result.Error)
	}

	return &discount, nil
}
func (r *discountRepository) GetDiscountByCodeRepository(discountCode string) (*model.Discount, error) {

	var discount model.Discount

	result := r.db.First(&discount, "discount_code = ?", discountCode)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("discount with Code %s not found", discountCode)
		}
		return nil, fmt.Errorf("error getting discount with Code %s: %s", discountCode, result.Error)
	}

	return &discount, nil
}

func (r *discountRepository) GetAllDiscountRepository(page, limit int) ([]*model.Discount, error) {
	var discount []*model.Discount
	offSet := (page - 1) * limit
	result := r.db.Offset(offSet).Limit(limit).Find(&discount)
	if result.Error != nil {
		return nil, errors.New("failed to get discount")
	}

	return discount, nil
}

func (r *discountRepository) UpdateDiscountByIdRepository(discountId string, discount *model.Discount) (*model.Discount, error) {

	result := r.db.Model(&model.Discount{}).Where("id = ?", discountId).Updates(discount)
	if result.Error != nil {
		return nil, errors.New("failed to Update discount")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("discount Not Found")
	}

	return discount, nil
}

func (r *discountRepository) DeleteDiscountByIdRepository(discountId string) error {

	result := r.db.Delete((&model.Discount{}), "id = ?", discountId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("discount Not Found")
	}
	return nil
}
