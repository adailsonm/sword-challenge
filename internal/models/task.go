package models

import "time"

type Task struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Summary   string    `json:"summary,omitempty"`
	User      User      `json:"user"`
	UserID    uint      `json:"-"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
