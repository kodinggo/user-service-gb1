package helper

import (
	"github.com/kodinggo/user-service-gb1/internal/model"
	"github.com/labstack/echo/v4"
)

func GetUserSession(c echo.Context) *model.User {
	u := c.Get("user")
	if u == nil {
		return nil
	}

	user, ok := u.(*model.User)
	if !ok {
		return nil
	}

	return user
}
