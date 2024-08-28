package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kodinggo/user-service-gb1/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type JWTMiddleware struct {
	authUsecase model.IAuthUsecase
}

func NewJWTMiddleware(authUsecase model.IAuthUsecase) *JWTMiddleware {
	return &JWTMiddleware{authUsecase: authUsecase}
}

func (m *JWTMiddleware) ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token format"})
		}

		// Call gRPC to validate the token
		user, err := m.authUsecase.ValidateToken(context.Background(), token)
		if err != nil || user == nil {
			logrus.Error(err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		}

		// Store the user information in the context
		c.Set("user", user)

		return next(c)
	}
}
