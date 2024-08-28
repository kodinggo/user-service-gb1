package handler

import (
	"net/http"

	"github.com/kodinggo/user-service-gb1/middleware-example/helper"
	"github.com/labstack/echo"
)

func (u *ProductHandler) GetProductByID(c echo.Context) error {
	userSession := helper.GetUserSession(c)
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
