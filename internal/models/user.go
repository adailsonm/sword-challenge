package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"id,omitempty" json:"id"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
