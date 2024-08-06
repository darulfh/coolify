package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/balance"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"
)

type BalanceController interface {
	CreateBalanceController(c echo.Context, payload model.GenerateVirtualAgregator) error
	GetPayBalanceStatusController(c echo.Context) error
}

type balanceController struct {
	BalanceUsecase balance.BalanceUsecase
}

func NewBalanceController(BalanceUseCase balance.BalanceUsecase) *balanceController {
	return &balanceController{BalanceUsecase: BalanceUseCase}
}

func (ctrl *balanceController) GenerateVaController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	var payload model.GenerateVirtualAgregator
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response, err := ctrl.BalanceUsecase.GenerateVaUseCase(userId, payload)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Create VA ",
		},
		Data: response,
	})
}

func (ctrl *balanceController) CreateBalanceController(c echo.Context) error {
	var payload model.PartnerCallbackVirtualAggregator

	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	response, err := ctrl.BalanceUsecase.CreateBalanceUseCase(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully processed the callback",
		},
		Data: response,
	})
}

func (ctrl *balanceController) GetPayBalanceStatusController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	vaId := c.Param("id")

	response, err := ctrl.BalanceUsecase.GetPayBalanceStatusUseCase(userId, vaId)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get VA",
		},
		Data: response,
	})
}
