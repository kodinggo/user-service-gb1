package model

import "context"

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type IRoleRepository interface {
	GetRoleByUserID(ctx context.Context, userID int64) (*Role, error)
}
