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

func(r *Repository) GetLogin(name string) (*User, error){
	var user User
	err := r.DB.QueryRow(context.Background(),"SELECT id, name, password FROM employee WHERE name=$1", name).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user,nil
}