package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
