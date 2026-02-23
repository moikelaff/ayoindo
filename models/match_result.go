package models

import (
	"time"

	"gorm.io/gorm"
)

type MatchResult struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	MatchID   uint           `json:"match_id" gorm:"uniqueIndex;not null"`
	HomeScore int            `json:"home_score" gorm:"default:0"`
	AwayScore int            `json:"away_score" gorm:"default:0"`
	Goals     []Goal         `json:"goals,omitempty" gorm:"foreignKey:MatchResultID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
