package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID      uint   `gorm:"id,omitempty"`
	Summary string `json:"summary,omitempty"`
	UserID  User   `gorm:"foreignKey:id"`
}
