package repository

import (
	"context"
	"database/sql"

	"github.com/kodinggo/user-service-gb1/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) model.IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user model.User) error {
	_, err := u.db.ExecContext(ctx, "INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Login(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email=?", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user, err
	}

	return &user, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email=?", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user, err
	}

	return &user, nil
}

func (u *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := u.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE id=?", id).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &user, err
	}

	return &user, nil
}
