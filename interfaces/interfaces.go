package interfaces

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string
	Username string
	Password string
	Balence  int
}

type Register struct {
	Username string
	Email    string
	Password string
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
