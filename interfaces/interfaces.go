package interfaces

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string    `validate:"required,email"`
	Username string
	Password string `validate:"required"`
	Balence  int
}

type Register struct {
	Username string
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
}

type UpdateUserEmail struct {
	Username string
	Password string
	Email    string
	NewEmail string
}

type UpdateUserBalance struct {
	Username string
	Password string
	Email    string
	TopUp    int
}

type Transaction struct {
	PayeeEmail string
	PayorEmail string
	Amount     int
}
