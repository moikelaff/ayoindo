package models

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	MatchResultID uint           `json:"match_result_id" gorm:"not null"`
	PlayerID      uint           `json:"player_id" gorm:"not null"`
	Player        *Player        `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
	Minute        int            `json:"minute" gorm:"not null"` // minute when goal occurred
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
