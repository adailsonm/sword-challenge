package models

import "time"

type User struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Role         string    `json:"role,omitempty"`
	Email        string    `json:"email,omitempty"`
	PasswordHash string    `json:"passwordHash,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
