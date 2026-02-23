package models

import (
	"time"

	"gorm.io/gorm"
)

// MatchStatus defines the state of a match
type MatchStatus string

const (
	MatchStatusScheduled MatchStatus = "scheduled"
	MatchStatusCompleted MatchStatus = "completed"
)

type Match struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	HomeTeamID  uint           `json:"home_team_id" gorm:"not null"`
	AwayTeamID  uint           `json:"away_team_id" gorm:"not null"`
	HomeTeam    *Team          `json:"home_team,omitempty" gorm:"foreignKey:HomeTeamID"`
	AwayTeam    *Team          `json:"away_team,omitempty" gorm:"foreignKey:AwayTeamID"`
	MatchDate   string         `json:"match_date" gorm:"not null"` // YYYY-MM-DD
	MatchTime   string         `json:"match_time" gorm:"not null"` // HH:MM
	Status      MatchStatus    `json:"status" gorm:"default:'scheduled'"`
	MatchResult *MatchResult   `json:"match_result,omitempty" gorm:"foreignKey:MatchID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
