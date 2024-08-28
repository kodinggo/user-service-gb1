package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kodinggo/user-service-gb1/internal/model"
)

type authUsecase struct {
	jwtSecret      string
	userRepository model.IUserRepository
	roleRepository model.IRoleRepository
}

func NewAuthUsecase(jwtSecret string, userRepository model.IUserRepository, roleRepository model.IRoleRepository) model.IAuthUsecase {
	return &authUsecase{
		jwtSecret:      jwtSecret,
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (uc *authUsecase) ValidateToken(ctx context.Context, tokenString string) (*model.User, error) {
	fmt.Println("secret: ", uc.jwtSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(uc.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("user ID not found in token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("email not found in token claims")
	}

	roleId, ok := claims["role_id"].(float64)
	if !ok {
		return nil, errors.New("roleId not found in token claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("role not found in token claims")
	}

	return &model.User{
		Id:    int64(userId),
		Email: email,
		Role: model.Role{
			Id: int64(roleId), Name: role,
		},
	}, nil
}
