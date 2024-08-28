package model

import (
	"context"
	"errors"
	"time"

	pb "github.com/kodinggo/user-service-gb1/pb/auth"
)

var (
	// ErrInvalidPassword is an error for invalid password
	ErrInvalidPassword = errors.New("invalid password")
	// ErrEmailNotFound is an error for email not found
	ErrEmailNotFound = errors.New("email not found")

	ErrUnauthorized = errors.New("unauthorized")
)

// User is a struct that represents a user for database model
type User struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Address   string     `json:"address"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Role Role `pq:"-" json:"role"`
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Role: &pb.Role{
			Id:   u.Role.Id,
			Name: u.Role.Name,
		},
	}
}

// IsPasswordMatch is a method to check if the password is matched
func (u User) IsPasswordMatch(password string) bool {
	return u.Password == password
}

// IUserRepository is an interface that represents a user repository
type IUserRepository interface {
	Create(ctx context.Context, user User) error
	Login(ctx context.Context, email string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)
}

// IUserUsecase is an interface that represents a user usecase
type IUserUsecase interface {
	Create(ctx context.Context, user User) error
	Login(ctx context.Context, email string, password string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
