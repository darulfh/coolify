package controller

import (
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"
	"github.com/darulfh/skuy_pay_be/usecase/pdam"

	"github.com/labstack/echo/v4"
)

type PdamController interface {
	CreatePdamController(c echo.Context) error
	GetAllPdamController(c echo.Context) error
	GetPdamByIdController(c echo.Context) error
	UpdatePdamController(c echo.Context) error
	DeletePdamByIdController(c echo.Context) error
	BillInquiryPdamController(c echo.Context) error
	PayBillInquiryPdamController(c echo.Context) error
}

type pdamController struct {
	pdamUseCase pdam.PdamUseCase
}

func NewPdamController(pdamUseCase pdam.PdamUseCase) *pdamController {
	return &pdamController{
		pdamUseCase: pdamUseCase,
	}
}

func (ctrl *pdamController) CreatePdamController(c echo.Context) error {
	var payload model.Pdam
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
	response, err := ctrl.pdamUseCase.CreatePdamUseCase(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Create PDAM",
		},
		Data: response,
	})
}

func (ctrl *pdamController) GetAllPdamController(c echo.Context) error {
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

	response, err := ctrl.pdamUseCase.GetAllPdamUseCase(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Pdams",
		},
		Data: response,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *pdamController) GetPdamByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	pdamID := c.Param("id")
	if pdamID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Pdam ID",
		})
	}

	response, err := ctrl.pdamUseCase.GetPdamByIdUseCase(pdamID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get PDAM",
		},
		Data: response,
	})
}

func (ctrl *pdamController) UpdatePdamController(c echo.Context) error {
	var payload model.Pdam
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	PdamID := c.Param("id")
	if PdamID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid PDAM ID",
		})
	}
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	response, err := ctrl.pdamUseCase.UpdatePdamByIdUseCase(PdamID, &payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Update Pdam",
		},
		Data: response,
	})
}

func (ctrl *pdamController) DeletePdamByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	PdamID := c.Param("id")
	if PdamID == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid PDAM ID",
		})
	}
	err := ctrl.pdamUseCase.DeletePdamByIDUseCase(PdamID)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Delete PDAM",
		},
	})
}

func (ctrl *pdamController) BillInquiryPdamController(c echo.Context) error {
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

	response, err := ctrl.pdamUseCase.BillInquiryPdamUseCase(userId, &payload)
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

func (ctrl *pdamController) PayBillInquiryPdamController(c echo.Context) error {
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

	response, err := ctrl.pdamUseCase.PayBillPdamUseCase(userId, &payload)
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
