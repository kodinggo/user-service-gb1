package grpc

import (
	"context"

	"github.com/kodinggo/user-service-gb1/internal/model"
	pbAuth "github.com/kodinggo/user-service-gb1/pb/auth"
)

type JWTValidatorServer struct {
	pbAuth.JWTValidatorServer
	jwtSecret   string
	authUsecase model.IAuthUsecase
}

func NewJWTValidatorServer(jwtSecret string, authUsecase model.IAuthUsecase) *JWTValidatorServer {
	return &JWTValidatorServer{jwtSecret: jwtSecret, authUsecase: authUsecase}
}

func (s *JWTValidatorServer) ValidateToken(ctx context.Context, req *pbAuth.ValidateTokenRequest) (*pbAuth.ValidateTokenResponse, error) {
	user, err := s.authUsecase.ValidateToken(ctx, req.Token)
	if err != nil {
		return &pbAuth.ValidateTokenResponse{Valid: false}, err
	}
	if user == nil {
		return &pbAuth.ValidateTokenResponse{Valid: false}, nil
	}

	return &pbAuth.ValidateTokenResponse{
		Valid: true,
		User:  user.ToProto(),
	}, nil
}
