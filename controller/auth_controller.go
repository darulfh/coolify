package controller

import (
	"net/http"

	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/auth"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	LoginController(c echo.Context) error
	RegisterController(c echo.Context) error
}

type authController struct {
	authUsecase auth.AuthUsecase
}

func NewAuthController(authUsecase auth.AuthUsecase) *authController {
	return &authController{
		authUsecase: authUsecase,
	}
}

func (u *authController) LoginController(c echo.Context) error {
	var payload model.User

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := u.authUsecase.LoginUseCase(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Login successfully",
		},
		Data: user,
	})

}

func (u *authController) RegisterController(c echo.Context) error {
	var payload model.User

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := u.authUsecase.RegisterUseCase(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusCreated,
			Message:    "User Created successfully",
		},
		Data: user,
	})

}

func (u *authController) RegisterAdminController(c echo.Context) error {
	var payload model.User

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := u.authUsecase.RegisterAdminUseCase(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusCreated,
			Message:    "Admin Created successfully",
		},
		Data: user,
	})

}
