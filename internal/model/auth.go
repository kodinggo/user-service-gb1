package model

import (
	"context"
)

type IAuthUsecase interface {
	ValidateToken(ctx context.Context, tokenString string) (*User, error)
}
