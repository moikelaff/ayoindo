package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"not null"`
	Logo        string         `json:"logo"`
	FoundedYear int            `json:"founded_year" gorm:"not null"`
	Address     string         `json:"address" gorm:"not null"`
	City        string         `json:"city" gorm:"not null"`
	Players     []Player       `json:"players,omitempty" gorm:"foreignKey:TeamID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
