package handler

import (
	"fmt"
	"time"

	"github.com/kodinggo/user-service-gb1/internal/config"
	"github.com/kodinggo/user-service-gb1/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jwtCustomClaims struct {
	Id     int64  `json:"id"`
	Email  string `json:"email"`
	RoleId int64  `json:"role_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func signJwtToken(id int64, email string, roleId int64, roleName string) (string, error) {
	claims := &jwtCustomClaims{
		id,
		email,
		roleId,
		roleName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Cfg.JWT.TTL))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("secret pas login: ", config.GetJwtSecret())
	t, err := token.SignedString([]byte(config.Cfg.JWT.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func getUserSession(c echo.Context) *model.User {
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
