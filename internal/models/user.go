package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	Email     string    `json:"email,omitempty" gorm:"unique,not null"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
