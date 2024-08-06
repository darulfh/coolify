package controller

import (
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/electricity"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/labstack/echo/v4"
)

type ElectricityController interface {
	CreateElectricityController(c echo.Context) error
	GetAllElectricityController(c echo.Context) error
	GetElectricityByIdController(c echo.Context) error
	UpdateElectricityController(c echo.Context) error
	DeleteElectricityByIdController(c echo.Context) error
	BillInquiryElectricityController(c echo.Context) error
	PayBillInquiryElectricityController(c echo.Context) error
	BuyBillInquiryElectricityController(c echo.Context) error

	ElectricityBillInquiryIakUseCase(c echo.Context) error
	ElectricityBillPayIakUseCase(c echo.Context) error

	ElectricityTokenInquiryIakUseCase(c echo.Context) error
	ElectricityTokenPayIakUseCase(c echo.Context) error
}

type electricityController struct {
	electricityUseCase electricity.ElectricityUseCase
}

func NewElectricityController(electricityUseCase electricity.ElectricityUseCase) *electricityController {
	return &electricityController{
		electricityUseCase: electricityUseCase,
	}
}

func (ctrl *electricityController) CreateElectricityController(c echo.Context) error {
	var payload model.Electricity
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
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
	response, err := ctrl.electricityUseCase.CreateElectricityUseCase(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Create electricity",
		},
		Data: response,
	})
}

func (ctrl *electricityController) GetAllElectricityController(c echo.Context) error {
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

	response, err := ctrl.electricityUseCase.GetAllElectricityUseCase(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get electricitys",
		},
		Data: response,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *electricityController) GetElectricityByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	electricityID := c.Param("id")
	if electricityID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid electricity ID",
		})
	}

	response, err := ctrl.electricityUseCase.GetElectricityByIdUseCase(electricityID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get electricity",
		},
		Data: response,
	})
}

func (ctrl *electricityController) UpdateElectricityController(c echo.Context) error {
	var payload model.Electricity
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	electricityID := c.Param("id")
	if electricityID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid electricity ID",
		})
	}
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	response, err := ctrl.electricityUseCase.UpdateElectricityByIdUseCase(electricityID, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Update electricity",
		},
		Data: response,
	})
}

func (ctrl *electricityController) DeleteElectricityByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	electricityID := c.Param("id")
	if electricityID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid electricity ID",
		})
	}
	err := ctrl.electricityUseCase.DeleteElectricityByIDUseCase(electricityID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Delete electricity",
		},
	})
}

func (ctrl *electricityController) BillInquiryPostPaidElectricityController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.OyBillerApi
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.PostBillInquiryElectricityUseCase(userId, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Succesfully",
		},
		Data: response,
	})
}

func (ctrl *electricityController) PayBillInquiryElectricityController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.OyBillerApi
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.PostPayBillElectricityUseCase(userId, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusAccepted,
			Message:    "Succesfully pay bill",
		},
		Data: response,
	})
}

// token
func (ctrl *electricityController) BillInquiryPrePaidElectricityController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.OyBillerApi
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.PreBillInquiryElectricityUseCase(userId, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Succesfully Get Bill",
		},
		Data: response,
	})
}

func (ctrl *electricityController) ElectricityBillInquiryIakUseCase(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.IakInquiryBody

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.ElectricityBillInquiryIakUseCase(&payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, response)
}

func (ctrl *electricityController) ElectricityBillPayIakUseCase(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.IakPayBody

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.ElectricityBillPayIakUseCase(&payload, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, response)
}

func (ctrl *electricityController) ElectricityTokenInquiryIakUseCase(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.PrePaidIakBody

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.ElectricityTokenInquiryIakUseCase(&payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, response)
}

func (ctrl *electricityController) ElectricityTokenPayIakUseCase(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload model.PrePaidIakBody

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := ctrl.electricityUseCase.ElectricityTokenPayIakUseCase(&payload, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, response)
}
