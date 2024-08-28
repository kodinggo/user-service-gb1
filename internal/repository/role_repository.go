package repository

import (
	"context"
	"database/sql"

	"github.com/kodinggo/user-service-gb1/internal/model"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) model.IRoleRepository {
	return &RoleRepository{db: db}
}

func (u *RoleRepository) GetRoleByUserID(ctx context.Context, userID int64) (*model.Role, error) {
	var role model.Role
	err := u.db.QueryRowContext(ctx, `SELECT r.id, r.name FROM roles r
										INNER JOIN user_roles ur ON ur.role_id = r.id 
										WHERE ur.user_id = ?`,
		userID).Scan(&role.Id, &role.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &role, err
	}

	return &role, nil
}
