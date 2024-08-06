package controller

import (
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/dto"
	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"
	pulsa "github.com/darulfh/skuy_pay_be/usecase/pulsa_paket_data"

	"github.com/labstack/echo/v4"
)

type PulsaPaketDataController interface {
	CreatePulsaPaketData(c echo.Context) error
	GetPPDByAdmin(c echo.Context) error
	GetPPDByUser(c echo.Context) error
	GetPPDByID(c echo.Context) error
	UpdatePPDById(c echo.Context) error
	DeletePPDById(c echo.Context) error
	CreateTransactionPPDUser(c echo.Context) error
}

type pulsaPaketDataController struct {
	PPDUsecase pulsa.PulsaPaketDataUsecase
}

func NewPulsaPaketDataController(ppdUsecase pulsa.PulsaPaketDataUsecase) *pulsaPaketDataController {
	return &pulsaPaketDataController{PPDUsecase: ppdUsecase}
}

func (ctrl *pulsaPaketDataController) CreatePulsaPaketData(c echo.Context) error {
	var payload model.PulsaPaketData

	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)

	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	ppd, err := ctrl.PPDUsecase.CreatePulsaPaketData(payload)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusCreated,
			Message:    "pulsa or paket data created successfully",
		},
		Data: ppd,
	})
}

func (ctrl *pulsaPaketDataController) GetPPDByAdmin(c echo.Context) error {
	var payload dto.PulsaDto
	payload.Type = c.QueryParam("type")
	payload.Provider = c.QueryParam("provider")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	payload.Page = page
	payload.Limit = limit

	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	ppd, err := ctrl.PPDUsecase.GetAllPulsaPaketData(payload, nil)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "success get all pulsa or paket data",
		},
		Data: ppd,
		Pagination: &model.Pagination{
			Limit: limit,
			Page:  page,
		},
	})
}

func (ctrl *pulsaPaketDataController) GetPPDByUser(c echo.Context) error {
	var payload dto.PulsaDto

	payload.Type = c.QueryParam("type")
	payload.Provider = c.QueryParam("provider")
	payload.PhoneNumber = c.QueryParam("phone_number")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	payload.Page = page
	payload.Limit = limit

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	isUser := true

	ppd, err := ctrl.PPDUsecase.GetAllPulsaPaketData(payload, &isUser)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "success get all pulsa or paket data",
		},
		Data: ppd,
	})
}

func (ctrl *pulsaPaketDataController) GetPPDByID(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "id ppd is empty",
		})
	}

	ppd, err := ctrl.PPDUsecase.GetPulsaPaketDataById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "success get PPD by id",
		},
		Data: ppd,
	})

}

func (ctrl *pulsaPaketDataController) UpdatePPDById(c echo.Context) error {
	var payload model.PulsaPaketData

	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)

	id := c.Param("id")

	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	ppd, err := ctrl.PPDUsecase.UpdatePulsaById(id, payload)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "pulsa or paket data updated successfully",
		},
		Data: ppd,
	})
}

func (ctrl *pulsaPaketDataController) DeletePPDById(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)

	id := c.Param("id")

	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	if err := ctrl.PPDUsecase.DeletePulsaById(id); err != nil {

		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "pulsa or paket data deleted successfully",
		},
	})
}

func (ctrl *pulsaPaketDataController) CreateTransactionPPDUser(c echo.Context) error {
	user := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload dto.TransactionPPDDto

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	result, err := ctrl.PPDUsecase.CreateTransactionPPD(user, payload)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusCreated,
			Message:    "success Transfer Amount",
		},
		Data: result,
	})
}
