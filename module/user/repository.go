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

func (r *Repository) GetUser(ctx context.Context) ([]GetUserModel, error) {
	rows, err := r.DB.Query(ctx, "SELECT e.id, e.username, p.name, p.position FROM employee e LEFT JOIN profile p ON e.id = p.employeeid")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user []GetUserModel
	for rows.Next() {
		var u GetUserModel
		err := rows.Scan(&u.ID, &u.Username, &u.Name, &u.Position)
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

func (r *Repository) GetUserByID(ctx context.Context, id string) (*GetUserModel, error) {
	objID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Error ObjID", err)
	}
	queryGetID := "SELECT e.id, e.username, p.name, p.position FROM employee e LEFT JOIN profile p ON e.id = p.employeeid WHERE e.id = $1"
	var u GetUserModel
	err = r.DB.QueryRow(ctx, queryGetID, objID).Scan(&u.ID, &u.Username, &u.Name, &u.Position)
	if err != nil {
		return nil, fmt.Errorf("ID Not Found")
	}
	return &u, nil
}

func (r *Repository) InsertUser(ctx context.Context, user *User, profile *Profile) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
		profile.ID = uuid.New()
	}
	queryEmployee := `INSERT INTO employee (id, username, password) VALUES ($1, $2, $3) RETURNING id`
	queryProfile := `INSERT INTO profile (id, name, position, employeeid) VALUES ($1, $2, $3, $4)`

	err := r.DB.QueryRow(ctx, queryEmployee, user.ID, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		fmt.Println("failed to insert employee: %w", err)
	}
	_, err = r.DB.Exec(ctx, queryProfile, profile.ID, profile.Name, profile.Position, user.ID)
	if err != nil {
		fmt.Println("failed to insert profile: %w", err)
	}
	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	objID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("error UUID %w", err)
	}

	queryProfile := `DELETE FROM profile WHERE employeeid = $1`
	_, err = r.DB.Exec(ctx, queryProfile, objID)
	if err != nil {
		return fmt.Errorf("failed to delete %w", err)
	}

	queryEmployee := `DELETE FROM employee WHERE id = $1`
	result, err := r.DB.Exec(ctx, queryEmployee, objID)
	if err != nil {
		return fmt.Errorf("failed to delete %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	return nil
}
