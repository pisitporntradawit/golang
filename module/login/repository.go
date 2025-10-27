package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetLogin(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.DB.QueryRow(ctx, "SELECT e.id, e.username, e.password, r.roleid, ro.rolename FROM employee e LEFT JOIN employeeroles r ON e.id = r.employeeid JOIN role ro ON ro.id = r.roleid WHERE username=$1", username).Scan(&user.ID, &user.Username, &user.Password, &user.Roleid, &user.Rolename)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

func (r *Repository) AuthRole(ctx context.Context, id string) (*UserRole, error) {
	var userrole UserRole
	var roles []string
	rows, err := r.DB.Query(ctx, "SELECT e.id, e.username, ro.rolename FROM employee e LEFT JOIN employeeroles r ON e.id = r.employeeid JOIN role ro ON ro.id = r.roleid WHERE e.id=$1", id)
	if err != nil {
		return nil, errors.New("User not found")
	}
	defer rows.Close()

	for rows.Next() {
		var roleName string
		if err := rows.Scan(&userrole.ID, &userrole.Username, &roleName); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		roles = append(roles, roleName)
	}
	fmt.Println(roles)
	// กรณีไม่มี role
	if len(roles) == 0 {
		return nil, errors.New("user has no roles or not found")
	}

	userrole.Rolename = roles

	return &userrole, nil
}
