package discount

import (
	"errors"
	"fmt"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository"
)

type DiscountUseCase interface {
	CreateDiscountUseCase(payload *model.Discount) (*model.Discount, error)
	GetAllDiscountUseCase(page, limit int) ([]*model.Discount, error)
	GetDiscountByIdUseCase(DiscountId string) (*model.Discount, error)
	GetDiscountByCodeUseCase(DiscountCode string) (*model.Discount, error)
	UpdateDiscountByIdUseCase(DiscountId string, payload *model.Discount) (*model.Discount, error)
	DeleteDiscountByIDUseCase(userId string) error
}

type discountUseCase struct {
	discountRepository repository.DiscountRepository
}

func NewDiscountUseCase(discountRepository repository.DiscountRepository) *discountUseCase {
	return &discountUseCase{discountRepository: discountRepository}
}

func (uc *discountUseCase) CreateDiscountUseCase(payload *model.Discount) (*model.Discount, error) {

	discount, err := uc.discountRepository.CreateDiscountRepository(payload)
	if err != nil {
		return nil, fmt.Errorf("error creating Discount in database: %w", err)
	}
	return discount, err

}

func (uc *discountUseCase) GetAllDiscountUseCase(page, limit int) ([]*model.Discount, error) {
	Discount, err := uc.discountRepository.GetAllDiscountRepository(page, limit)

	if err != nil {
		return nil, err
	}

	return Discount, nil
}

func (uc *discountUseCase) GetDiscountByIdUseCase(DiscountId string) (*model.Discount, error) {

	Discount, err := uc.discountRepository.GetDiscountByIdRepository(DiscountId)
	if err != nil {
		return nil, fmt.Errorf("discount with ID: %s not found", DiscountId)
	}

	return Discount, nil

}

func (uc *discountUseCase) GetDiscountByCodeUseCase(DiscountCode string) (*model.Discount, error) {

	Discount, err := uc.discountRepository.GetDiscountByCodeRepository(DiscountCode)
	if err != nil {
		return nil, fmt.Errorf("discount with Code: %s not found", DiscountCode)
	}

	return Discount, nil

}

func (uc *discountUseCase) UpdateDiscountByIdUseCase(DiscountId string, payload *model.Discount) (*model.Discount, error) {
	discount, err := uc.discountRepository.GetDiscountByIdRepository(DiscountId)
	if err != nil {
		return nil, fmt.Errorf("failed to update Discount: %v", err)
	}

	discount.DiscountCode = payload.DiscountCode
	discount.Description = payload.Description
	discount.Image = payload.Image
	discount.UpdatedAt = time.Now()

	updateDiscount, err := uc.discountRepository.UpdateDiscountByIdRepository(DiscountId, discount)
	if err != nil {
		return nil, fmt.Errorf("failed to update Discount: %v", err)
	}

	return updateDiscount, nil
}

func (uc *discountUseCase) DeleteDiscountByIDUseCase(userId string) error {
	err := uc.discountRepository.DeleteDiscountByIdRepository(userId)
	if err != nil {
		return errors.New("discount not found")
	}
	return err
}
