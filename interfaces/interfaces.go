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
	Password string
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
	Email    string `validate:"required,email"`
}

type UpdateUserEmail struct {
	Username string
	Password string
	Email    string `validate:"required,email"`
	NewEmail string `validate:"required,email"`
}

type UpdateUserBalance struct {
	Username string
	Password string
	Email    string `validate:"required,email"`
	TopUp    int
}

type Transaction struct {
	PayeeEmail string `validate:"required,email"`
	PayorEmail string `validate:"required,email"`
	Amount     int
}

type Account struct {
	OwnerID   uuid.UUID
	Balance   int
	Overdraft bool
}
