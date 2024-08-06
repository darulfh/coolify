package controller

import (
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/model"
	insurance "github.com/darulfh/skuy_pay_be/usecase/Insurance"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"

	"github.com/labstack/echo/v4"
)

type InsuranceController interface {
	CreateInsuranceController(c echo.Context) error
	GetAllInsuranceController(c echo.Context) error
	GetInsuranceByIdController(c echo.Context) error
	UpdateInsuranceController(c echo.Context) error
	DeleteInsuranceByIdController(c echo.Context) error
	BillInquiryInsuranceController(c echo.Context) error
	PayBillInquiryInsuranceController(c echo.Context) error

	BpjsInquiryIakUseCase(c echo.Context) error
	BpjsPayIakUseCase(c echo.Context) error
}

type insuranceController struct {
	insuranceUseCase insurance.InsuranceUseCase
}

func NewInsuranceController(insuranceUseCase insurance.InsuranceUseCase) *insuranceController {
	return &insuranceController{
		insuranceUseCase: insuranceUseCase,
	}
}

func (ctrl *insuranceController) CreateInsuranceController(c echo.Context) error {
	var payload model.Insurance
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
	response, err := ctrl.insuranceUseCase.CreateInsuranceUseCase(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Create insurance",
		},
		Data: response,
	})
}

func (ctrl *insuranceController) GetAllInsuranceController(c echo.Context) error {
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

	response, err := ctrl.insuranceUseCase.GetAllInsuranceUseCase(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get insurances",
		},
		Data: response,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *insuranceController) GetInsuranceByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	insuranceID := c.Param("id")
	if insuranceID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid insurance ID",
		})
	}

	response, err := ctrl.insuranceUseCase.GetInsuranceByIdUseCase(insuranceID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get insurance",
		},
		Data: response,
	})
}

func (ctrl *insuranceController) UpdateInsuranceController(c echo.Context) error {
	var payload model.Insurance
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	insuranceID := c.Param("id")
	if insuranceID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid insurance ID",
		})
	}
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	response, err := ctrl.insuranceUseCase.UpdateInsuranceByIdUseCase(insuranceID, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Update insurance",
		},
		Data: response,
	})
}

func (ctrl *insuranceController) DeleteInsuranceByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	insuranceID := c.Param("id")
	if insuranceID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid insurance ID",
		})
	}
	err := ctrl.insuranceUseCase.DeleteInsuranceByIDUseCase(insuranceID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Delete insurance",
		},
	})
}

func (ctrl *insuranceController) BillInquiryInsuranceController(c echo.Context) error {
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

	response, err := ctrl.insuranceUseCase.BillInquiryInsuranceUseCase(userId, &payload)

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

func (ctrl *insuranceController) PayBillInquiryInsuranceController(c echo.Context) error {
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

	response, err := ctrl.insuranceUseCase.PayBillInsuranceUseCase(userId, &payload)
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

func (ctrl *insuranceController) BpjsInquiryIakUseCase(c echo.Context) error {
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

	response, err := ctrl.insuranceUseCase.BpjsInquiryIakUseCase(&payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, response)

}
func (ctrl *insuranceController) BpjsPayIakUseCase(c echo.Context) error {
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

	response, err := ctrl.insuranceUseCase.BpjsPayIakUseCase(&payload, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusAccepted, response)

}
