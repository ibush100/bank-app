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
