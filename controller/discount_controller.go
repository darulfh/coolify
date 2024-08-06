package controller

import (
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/cloudinary"
	"github.com/darulfh/skuy_pay_be/usecase/discount"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/labstack/echo/v4"
)

type DiscountController interface {
	CreatediscountController(c echo.Context) error
	GetAlldiscountController(c echo.Context) error
	GetdiscountByIdController(c echo.Context) error
	UpdatediscountController(c echo.Context) error
	DeletediscountByIdController(c echo.Context) error
}

type discountController struct {
	discountUseCase discount.DiscountUseCase
}

func NewDiscountController(discountUseCase discount.DiscountUseCase) *discountController {
	return &discountController{
		discountUseCase: discountUseCase,
	}
}

func (ctrl *discountController) CreateDiscountController(c echo.Context) error {
	var payload model.Discount
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	imageFiles := form.File["image"]
	if len(imageFiles) > 0 {
		file, _ := imageFiles[0].Open()
		defer file.Close()

		imageURL, err := cloudinary.ImageUploadHelper(file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
		}
		payload.Image = imageURL
	}

	payload.DiscountCode = form.Value["discount_code"][0]
	payload.Description = form.Value["description"][0]
	discountPriceStr := form.Value["discount_price"][0]
	discountPrice, err := strconv.ParseFloat(discountPriceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid discount price",
		})
	}
	payload.DiscountPrice = discountPrice

	response, err := ctrl.discountUseCase.CreateDiscountUseCase(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Create discount ",
		},
		Data: response,
	})

}

func (ctrl *discountController) GetAllDiscountController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	response, err := ctrl.discountUseCase.GetAllDiscountUseCase(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get discounts ",
		},
		Data: response,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *discountController) GetDiscountByIdController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	discountID := c.Param("id")
	if discountID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid discount ID",
		})
	}

	response, err := ctrl.discountUseCase.GetDiscountByIdUseCase(discountID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get discount",
		},
		Data: response,
	})

}

func (ctrl *discountController) GetDiscountByCodeController(c echo.Context) error {
	var payload model.Discount

	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.discountUseCase.GetDiscountByCodeUseCase(payload.DiscountCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get discount",
		},
		Data: response,
	})

}

func (ctrl *discountController) UpdateDiscountController(c echo.Context) error {

	var payload model.Discount
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	discountID := c.Param("id")
	_, err := ctrl.discountUseCase.GetDiscountByIdUseCase(discountID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})

	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	imageFiles := form.File["image"]
	if len(imageFiles) > 0 {
		file, _ := imageFiles[0].Open()
		defer file.Close()

		imageURL, err := cloudinary.ImageUploadHelper(file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
		}
		payload.Image = imageURL
	}

	payload.DiscountCode = form.Value["discount_code"][0]
	payload.Description = form.Value["description"][0]
	discountPriceStr := form.Value["discount_price"][0]
	discountPrice, err := strconv.ParseFloat(discountPriceStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid discount price",
		})
	}
	payload.DiscountPrice = discountPrice

	response, err := ctrl.discountUseCase.UpdateDiscountByIdUseCase(discountID, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Create discount ",
		},
		Data: response,
	})

}

func (ctrl *discountController) DeleteDiscountByIdController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	discountID := c.Param("id")
	if discountID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid discount ID",
		})
	}

	err := ctrl.discountUseCase.DeleteDiscountByIDUseCase(discountID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Delete discount",
		},
	})

}
