package login

import (
	// "database/sql"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Roleid   string    `json:"roleid"`
	Rolename string    `json:"rolename"`
}

type UserRole struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Roleid   string    `json:"roleid"`
	Rolename []string    `json:"rolename"`
}
