package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/darulfh/skuy_pay_be/dto"
	"github.com/darulfh/skuy_pay_be/model"
	"github.com/darulfh/skuy_pay_be/usecase/cloudinary"
	"github.com/darulfh/skuy_pay_be/usecase/middlewares"
	"github.com/darulfh/skuy_pay_be/usecase/users"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetAllUsersController(c echo.Context)
	GetUserByIdController(c echo.Context) error
	AdminGetUserByIdController(c echo.Context) error
	UpdateUserImageByIDController(c echo.Context) error
	UpdateUserByIDController(c echo.Context) error
	DeleteUserByIDController(c echo.Context) error
	UpdatePasswordController(c echo.Context) error
	CreatePINController(c echo.Context) error
	UpdatePINController(c echo.Context) error
	CheckPINController(c echo.Context) error
	TransferAmountController(c echo.Context) error
	GetUserByEmailController(c echo.Context) error
	CheckPinForgotPasswordController(c echo.Context) error
	UpdateForgotPasswordController(c echo.Context) error
	GetUsersByQueryController(c echo.Context) error
}

type userController struct {
	UserUsecase users.UserUsecase
}

func NewUserController(userUsecase users.UserUsecase) *userController {
	return &userController{
		UserUsecase: userUsecase,
	}
}

func (ctrl *userController) GetAllUsersController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Token unauthorized",
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

	name := c.QueryParam("name")
	users, err := ctrl.UserUsecase.GetAllUsersUseCase(page, limit, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get users: %v", err),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get users",
		},
		Data: users,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

func (ctrl *userController) GetUserByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	user, err := ctrl.UserUsecase.GetUserByIDUseCase(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    fmt.Sprintf("Successfully get user with ID: %s", user.ID),
		},
		Data: user,
	})

}

func (ctrl *userController) AdminGetUserByIdController(c echo.Context) error {

	id := c.Param("id")

	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}

	user, err := ctrl.UserUsecase.GetUserByIDUseCase(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    fmt.Sprintf("Successfully get user with ID: %s", user.ID),
		},
		Data: user,
	})

}

func (ctrl *userController) UpdateUserImageByIDController(c echo.Context) error {

	var payload model.User

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
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

	user, err := ctrl.UserUsecase.UpdateUserImageByIDUseCase(userId, &payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Image Updated successfully",
		},
		Data: user,
	})
}

func (ctrl *userController) UpdateUserByIDController(c echo.Context) error {
	var payload model.User

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
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

	user, err := ctrl.UserUsecase.UpdateUserByIDUseCase(userId, &payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Image Updated successfully",
		},
		Data: user,
	})
}

func (ctrl *userController) DeleteUserByIDController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})

	}
	err := ctrl.UserUsecase.DeleteUserByIDUseCase(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Deleted successfully",
		},
	})

}

func (ctrl *userController) UpdatePasswordController(c echo.Context) error {
	var payload dto.Password

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
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

	err := ctrl.UserUsecase.ChangePasswordUseCase(userId, payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Password updated successfully",
		},
	})
}

func (ctrl *userController) CreatePINController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}
	var payload dto.PIN

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	err := ctrl.UserUsecase.CreatePINUseCase(userId, payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "PIN Created successfully",
		},
	})
}

func (ctrl *userController) UpdatePINController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload dto.PIN

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	err := ctrl.UserUsecase.ChangePINUseCase(userId, payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "PIN updated successfully",
		},
	})
}

func (ctrl *userController) CheckPINController(c echo.Context) error {

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload dto.PIN
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
	err := ctrl.UserUsecase.CheckPINUseCase(userId, payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Correct",
		},
	})
}

func (ctrl *userController) TransferAmountController(c echo.Context) error {
	user := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if user == "" {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token unauthorized",
		})
	}

	var payload dto.TransactionTransferDto

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
	result, err := ctrl.UserUsecase.TransferAmountUseCase(user, payload)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusCreated,
			Message:    "success create transaction transfer",
		},
		Data: result,
	})
}

func (ctrl *userController) GetUserByEmailController(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "email empty",
		})
	}

	result, err := ctrl.UserUsecase.GetUserByEmailUsecase(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get User by email",
		},
		Data: result,
	})
}
func (ctrl *userController) CheckPinForgotPasswordController(c echo.Context) error {
	var payload dto.PIN
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
	userID := c.Param("id")

	err := ctrl.UserUsecase.CheckPINUseCase(userID, payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Correct",
		},
	})
}
func (ctrl *userController) UpdateForgotPasswordController(c echo.Context) error {
	userID := c.Param("id")

	var payload dto.Password

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := ctrl.UserUsecase.UpdateForgotPasswordUsecase(userID, payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Password updated successfully",
		},
		Data: user,
	})
}

func (ctrl *userController) GetUsersByQueryController(c echo.Context) error {

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

	users, err := ctrl.UserUsecase.GetUserByQueryUseCase(query, page, limit)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully Get Users",
		},
		Data: users,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}
