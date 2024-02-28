package models

import (
	"time"

	"github.com/grim13/go-rest-api/db"
	"github.com/grim13/go-rest-api/utils"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `binding:"required" json:"email"`
	Password  string    `binding:"required" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) Save() error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	query := `INSERT INTO users (email, password, created_at, updated_at) values (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
