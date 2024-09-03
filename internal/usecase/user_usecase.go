package usecase

import (
	"context"
	"encoding/json"

	"github.com/kodinggo/user-service-gb1/internal/model"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	userRepo model.IUserRepository
	roleRepo model.IRoleRepository
	js       jetstream.JetStream
}

func NewUserUsecase(
	userRepo model.IUserRepository,
	roleRepo model.IRoleRepository,
	js jetstream.JetStream,
) model.IUserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		roleRepo: roleRepo,
		js:       js,
	}
}

func (u *UserUsecase) Create(ctx context.Context, user model.User) error {
	log := logrus.WithFields(logrus.Fields{
		"user": user,
	})

	err := u.userRepo.Create(ctx, user)
	if err != nil {
		log.Error(err)
		return err
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = u.js.Publish(ctx, "USER.created", data)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Login is a usecase to login a user
// It will return a user and an error
func (u *UserUsecase) Login(ctx context.Context, email string, password string) (*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"email":    email,
		"password": password,
	})

	user, err := u.userRepo.Login(ctx, email)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if user == nil {
		return nil, model.ErrUnauthorized
	}

	if !user.IsPasswordMatch(password) {
		log.Error(model.ErrInvalidPassword)
		return nil, model.ErrInvalidPassword
	}

	role, err := u.roleRepo.GetRoleByUserID(ctx, user.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if role == nil {
		return nil, model.ErrUnauthorized
	}

	user.Role = *role

	return user, nil
}

// FindByEmail is a usecase to find a user by email
func (u *UserUsecase) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	log := logrus.WithFields(logrus.Fields{
		"email": email,
	})

	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return user, nil
}
