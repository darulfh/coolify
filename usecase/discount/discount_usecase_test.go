package discount_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/repository/mocks"
	"github.com/darulfh/skuy_pay_be/usecase/discount"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DiscountUseCaseTest struct {
	suite.Suite
	discountUseCase discount.DiscountUseCase

	discountRepositoryMock *mocks.DiscountRepository
}

func TestDiscountUseCaseTest(t *testing.T) {
	suite.Run(t, new(DiscountUseCaseTest))
}

func (s *DiscountUseCaseTest) SetupTest() {
	s.discountRepositoryMock = &mocks.DiscountRepository{}
	s.discountUseCase = discount.NewDiscountUseCase(s.discountRepositoryMock)
}

func (s *DiscountUseCaseTest) TestCreateDiscountUseCase_Success() {
	payload := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	expectedDiscount := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	s.discountRepositoryMock.On("CreateDiscountRepository", payload).Return(expectedDiscount, nil)

	result, err := s.discountUseCase.CreateDiscountUseCase(payload)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedDiscount, result)
	s.discountRepositoryMock.AssertCalled(s.T(), "CreateDiscountRepository", payload)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestCreateDiscountUseCase_Error() {
	payload := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	expectedErr := errors.New("failed to create discount")
	s.discountRepositoryMock.On("CreateDiscountRepository", payload).Return(nil, expectedErr)

	result, err := s.discountUseCase.CreateDiscountUseCase(payload)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.discountRepositoryMock.AssertCalled(s.T(), "CreateDiscountRepository", payload)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestGetAllDiscountUseCase_Success() {
	page := 1
	limit := 10
	expectedDiscounts := []*model.Discount{
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "1",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			DiscountCode: "LEBARAN",
			Description:  "Hari Raya",
			Image:        "url",
		},
		{
			UUIDPrimaryKey: model.UUIDPrimaryKey{
				ID:        "2",
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			DiscountCode: "WINTER50",
			Description:  "Winter Sale",
			Image:        "winter_discount.png",
		},
	}
	s.discountRepositoryMock.On("GetAllDiscountRepository", page, limit).Return(expectedDiscounts, nil)

	result, err := s.discountUseCase.GetAllDiscountUseCase(page, limit)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedDiscounts, result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetAllDiscountRepository", page, limit)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestGetAllDiscountUseCase_Error() {
	page := 1
	limit := 10
	expectedErr := errors.New("failed to get all discounts")
	s.discountRepositoryMock.On("GetAllDiscountRepository", page, limit).Return(nil, expectedErr)

	result, err := s.discountUseCase.GetAllDiscountUseCase(page, limit)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetAllDiscountRepository", page, limit)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestGetDiscountByIdUseCase_Success() {
	discountID := "1"
	expectedDiscount := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	s.discountRepositoryMock.On("GetDiscountByIdRepository", discountID).Return(expectedDiscount, nil)

	result, err := s.discountUseCase.GetDiscountByIdUseCase(discountID)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedDiscount, result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestGetDiscountByIdUseCase_Error() {
	discountID := "1"
	expectedErr := fmt.Errorf("discount with ID: %s not found", discountID)
	s.discountRepositoryMock.On("GetDiscountByIdRepository", discountID).Return(nil, expectedErr)

	result, err := s.discountUseCase.GetDiscountByIdUseCase(discountID)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestGetDiscountByCodeUseCase_Success() {
	discountCode := "LEBARAN"
	expectedDiscount := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	s.discountRepositoryMock.On("GetDiscountByCodeRepository", discountCode).Return(expectedDiscount, nil)

	result, err := s.discountUseCase.GetDiscountByCodeUseCase(discountCode)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedDiscount, result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByCodeRepository", discountCode)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestGetDiscountByCodeUseCase_Error() {
	discountCode := "LEBARAN"
	expectedErr := fmt.Errorf("discount with Code: %s not found", discountCode)
	s.discountRepositoryMock.On("GetDiscountByCodeRepository", discountCode).Return(nil, expectedErr)

	result, err := s.discountUseCase.GetDiscountByCodeUseCase(discountCode)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByCodeRepository", discountCode)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestUpdateDiscountByIdUseCase_Success() {
	discountID := "1"
	payload := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	expectedUpdatedDiscount := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Updated Sale",
		Image:        "updated_discount.png",
	}
	s.discountRepositoryMock.On("GetDiscountByIdRepository", discountID).Return(payload, nil)
	s.discountRepositoryMock.On("UpdateDiscountByIdRepository", discountID, payload).Return(expectedUpdatedDiscount, nil)

	result, err := s.discountUseCase.UpdateDiscountByIdUseCase(discountID, payload)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedUpdatedDiscount, result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertCalled(s.T(), "UpdateDiscountByIdRepository", discountID, payload)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestUpdateDiscountByIdUseCase_Error_GetDiscount() {
	discountID := "1"
	payload := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	expectedErr := fmt.Errorf("failed to update Discount: discount not found")
	s.discountRepositoryMock.On("GetDiscountByIdRepository", discountID).Return(nil, expectedErr)

	result, err := s.discountUseCase.UpdateDiscountByIdUseCase(discountID, payload)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertNotCalled(s.T(), "UpdateDiscountByIdRepository", discountID, payload)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestUpdateDiscountByIdUseCase_Error_UpdateDiscount() {
	discountID := "1"
	payload := &model.Discount{
		UUIDPrimaryKey: model.UUIDPrimaryKey{
			ID:        "1",
			DeletedAt: gorm.DeletedAt{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DiscountCode: "LEBARAN",
		Description:  "Hari Raya",
		Image:        "url",
	}
	expectedErr := fmt.Errorf("failed to update Discount: failed to update discount")
	s.discountRepositoryMock.On("GetDiscountByIdRepository", discountID).Return(payload, nil)
	s.discountRepositoryMock.On("UpdateDiscountByIdRepository", discountID, payload).Return(nil, expectedErr)

	result, err := s.discountUseCase.UpdateDiscountByIdUseCase(discountID, payload)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.discountRepositoryMock.AssertCalled(s.T(), "GetDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertCalled(s.T(), "UpdateDiscountByIdRepository", discountID, payload)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestDeleteDiscountByIDUseCase_Success() {
	discountID := "1"
	s.discountRepositoryMock.On("DeleteDiscountByIdRepository", discountID).Return(nil)

	err := s.discountUseCase.DeleteDiscountByIDUseCase(discountID)

	assert.NoError(s.T(), err)
	s.discountRepositoryMock.AssertCalled(s.T(), "DeleteDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertExpectations(s.T())
}

func (s *DiscountUseCaseTest) TestDeleteDiscountByIDUseCase_Error() {
	discountID := "1"
	expectedErr := errors.New("discount not found")
	s.discountRepositoryMock.On("DeleteDiscountByIdRepository", discountID).Return(expectedErr)

	err := s.discountUseCase.DeleteDiscountByIDUseCase(discountID)

	assert.Error(s.T(), err)
	s.discountRepositoryMock.AssertCalled(s.T(), "DeleteDiscountByIdRepository", discountID)
	s.discountRepositoryMock.AssertExpectations(s.T())
}
