package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"
	"github.com/darulfh/skuy_pay_be/usecase/transaction"
)

type TransactionController interface {
	GetAllTransactionsController(c echo.Context) error
	AdminGetTransactionByUserIdController(c echo.Context) error
	GetTransactionByUserIdController(c echo.Context) error
	GetTransactionByProductController(c echo.Context) error
	GetTransactionByQueryController(c echo.Context) error
	GetTransactionsPriceCountController(c echo.Context) error
}

type transactionController struct {
	transactionUsecase transaction.TransactionUsecase
}

func NewTransactionController(transactionUsecase transaction.TransactionUsecase) *transactionController {
	return &transactionController{
		transactionUsecase: transactionUsecase,
	}
}

func (ctrl *transactionController) GetAllTransactionsController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
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

	transactions, err := ctrl.transactionUsecase.GetAllTransactionsUseCase(page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *transactionController) GetTransactionByIdController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ALL_TYPE, c)

	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	id := c.Param("id")

	transactions, err := ctrl.transactionUsecase.GetTransactionByIdUseCase(id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
	})
}
func (ctrl *transactionController) AdminGetTransactionByUserIdController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	userID := c.Param("id")
	productType := c.QueryParam("product")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionByUserIdUseCase(userID, productType, page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *transactionController) GetTransactionByUserIdController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	productType := c.QueryParam("product")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionByUserIdUseCase(user, productType, page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *transactionController) GetTransactionByProductController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	product := c.QueryParam("product")
	status := c.QueryParam("status")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionProductTypeUseCase(product, status, page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *transactionController) GetTransactionByQueryController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	query := c.QueryParam("query")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionQueryUseCase(query, page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *transactionController) GetTransactionByStatusQueryController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	status := c.QueryParam("status")
	query := c.QueryParam("query")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionStatusQueryUseCase(query, status, page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Transaction",
		},
		Data: transactions,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *transactionController) GetTransactionsPriceCountController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionsPriceCountUseCase()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Total Transactions by Product Type",
		},
		Data: transactions,
	})
}

func (ctrl *transactionController) GetTransactionsPriceByMonthController(c echo.Context) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	transactions, err := ctrl.transactionUsecase.GetTransactionsPriceByMonthUseCase()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Total Transactions by month and year",
		},
		Data: transactions,
	})
}
