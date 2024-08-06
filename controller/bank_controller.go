package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/bank"
	"github.com/darulfh/skuy_pay_be/usecase/cloudinary"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/labstack/echo/v4"
)

type BankController interface {
	CreateBankController(c echo.Context) error
	GetAllBankController(c echo.Context) error
	GetBankByIdController(c echo.Context) error
	UpdateBankController(c echo.Context) error
	DeleteBankByIdController(c echo.Context) error
}

type bankController struct {
	bankUsecase bank.BankUseCase
}

func NewBankController(bankUsecase bank.BankUseCase) *bankController {
	return &bankController{
		bankUsecase: bankUsecase,
	}
}

func (ctrl *bankController) CreateBankController(c echo.Context) error {
	var payload model.Bank
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

	payload.Name = form.Value["name"][0]
	payload.BankCode = form.Value["bank_code"][0]

	bank, err := ctrl.bankUsecase.CreateBankUseCase(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "bank Created Successfully",
		},
		Data: bank,
	})
}

func (ctrl *bankController) GetAllBankController(c echo.Context) error {
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

	bank, err := ctrl.bankUsecase.GetAllBanksUseCase(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get banks ",
		},
		Data: bank,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *bankController) GetBankByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	bankID := c.Param("id")
	if bankID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid bank ID",
		})
	}

	bank, err := ctrl.bankUsecase.GetBankByIdUseCase(bankID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    fmt.Sprintf("Successfully get bank with ID: %s", bank.ID),
		},
		Data: bank,
	})

}

func (ctrl *bankController) UpdateBankController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	bankID := c.Param("id")
	if bankID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid bank ID",
		})
	}

	var payload model.Bank
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

	payload.Name = form.Value["name"][0]
	payload.BankCode = form.Value["bank_code"][0]

	bank, err := ctrl.bankUsecase.UpdateBankByIdUseCase(bankID, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "bank Updated Successfully",
		},
		Data: bank,
	})
}

func (ctrl *bankController) DeleteBankByIdController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}
	bankID := c.Param("id")
	if bankID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Bank ID",
		})
	}

	err := ctrl.bankUsecase.DeleteBankByIdUseCase(bankID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Delete Bank",
		},
	})

}
