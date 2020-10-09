package users

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string    `validate:"required"`
	Username string    `validate:"required"`
	Password string    `validate:"required"`
	Balence  int       `validate:omitempty`
}

type Register struct {
	Username string
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
}

type UpdateUserEmail struct {
	Username string `validate:"required"`
	Password string
	Email    string
	NewEmail string
}
