package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (r *Repository) GetUser(ctx context.Context) ([]User, error) {
	rows, err := r.DB.Query(ctx, "SELECT id, name, position FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Position)
		if err != nil {
			return nil, err
		}
		user = append(user, u)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return user, nil
}

func (r *Repository) InsertUser(ctx context.Context, user *User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	query := `INSERT INTO employee (id, name, position) VALUES ($1, $2, $3) RETURNING id`

	err := r.DB.QueryRow(ctx, query, user.ID, user.Name, user.Position).Scan(&user.ID)
	if err != nil {
		fmt.Println("failed to insert employee: %w", err)
	}
	return nil
}

func(r *Repository) DeleteUser(ctx context.Context, id string) error{
	objID, err := uuid.Parse(id)
	if err != nil{
		return fmt.Errorf("error UUID %w", err)
	}

	query := `DELETE FROM employee WHERE id = $1`

	result, err := r.DB.Exec(ctx, query, objID)
	if err != nil {
		return fmt.Errorf("failed to delete %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	return nil
}