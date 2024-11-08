package domain

import (
	"time"
)

type User struct {
	ID       uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name" validate:"required"`
	Email    string    `gorm:"size:255;not null" json:"email" validate:"required,email"`
	Password string    `gorm:"size:255;not null" json:"password" validate:"required"`
	Birthday time.Time `gorm:"not null" json:"birthday" validate:"required"`
}
