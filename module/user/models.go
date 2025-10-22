package user

import (
	// "database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type GetUserModel struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Position string    `json:"position"`
}

type Profile struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Position   string    `json:"position"`
	EmployeeID uuid.UUID `json:"employeeid"`
}

type UserProfile struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Position   string    `json:"position"`
	EmployeeID uuid.UUID `json:"employeeid"`
}
