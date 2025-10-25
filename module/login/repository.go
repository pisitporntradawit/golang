package login

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository{
	return &Repository{
		DB:db,
	}
}

func(r *Repository) GetLogin(ctx context.Context, username string) (*User, error){
	var user User
	err := r.DB.QueryRow(ctx,"SELECT e.id, e.username, e.password, p.position FROM employee e LEFT JOIN profile p ON e.id = p.employeeid WHERE username=$1" , username).Scan(&user.ID, &user.Username, &user.Password, &user.Position)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user,nil
}