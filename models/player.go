package models

import (
	"time"

	"gorm.io/gorm"
)

// PlayerPosition defines allowed positions
type PlayerPosition string

const (
	PositionPenyerang     PlayerPosition = "penyerang"
	PositionGelandang     PlayerPosition = "gelandang"
	PositionBertahan      PlayerPosition = "bertahan"
	PositionPenjagaGawang PlayerPosition = "penjaga_gawang"
)

type Player struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	TeamID       uint           `json:"team_id" gorm:"not null"`
	Team         *Team          `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Name         string         `json:"name" gorm:"not null"`
	Height       float64        `json:"height" gorm:"not null"` // cm
	Weight       float64        `json:"weight" gorm:"not null"` // kg
	Position     PlayerPosition `json:"position" gorm:"not null"`
	JerseyNumber int            `json:"jersey_number" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
