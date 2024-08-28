package handler

import (
	"errors"
	"net/http"

	"github.com/kodinggo/user-service-gb1/internal/model"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase model.IUserUsecase
}

func NewUserHandler(e *echo.Group, userUsecase model.IUserUsecase, JWTMiddleware echo.MiddlewareFunc) {
	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	e.POST("/users", userHandler.Create)
	e.POST("/users/login", userHandler.Login)

	protected := e.Group("/users/profile")
	protected.Use(JWTMiddleware)

	protected.GET("", userHandler.GetProfile)
}

func (u *UserHandler) Create(c echo.Context) error {
	var user model.User // TODO: better create model.CreateUserRequest
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Message: err.Error(),
		})
	}

	err = u.userUsecase.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (u *UserHandler) Login(c echo.Context) error {
	var reqUser model.User
	err := c.Bind(&reqUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if reqUser.Email == "" || reqUser.Password == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: "email and password are required",
		})
	}

	user, err := u.userUsecase.Login(c.Request().Context(), reqUser.Email, reqUser.Password)
	if err != nil {
		if errors.Is(err, model.ErrUnauthorized) {
			return c.JSON(http.StatusUnauthorized, response{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
		}

		if errors.Is(err, model.ErrUnauthorized) {
			return c.JSON(http.StatusNotFound, response{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// "admin" is hardcoded as the role
	// in a real application, the role should be fetched from the database
	token, err := signJwtToken(user.Id, user.Email, user.Role.Id, user.Role.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]string{"token": token},
	})
}

func (u *UserHandler) GetProfile(c echo.Context) error {
	userSession := getUserSession(c)
	if userSession == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		})
	}

	user, err := u.userUsecase.FindByEmail(c.Request().Context(), userSession.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    user,
	})
}
