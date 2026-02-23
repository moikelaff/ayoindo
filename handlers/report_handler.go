package handlers

import (
	"net/http"

	"ayoindo/config"
	"ayoindo/models"
	"ayoindo/utils"

	"github.com/gin-gonic/gin"
)

type TopScorer struct {
	PlayerID   uint   `json:"player_id"`
	PlayerName string `json:"player_name"`
	Goals      int    `json:"goals"`
}

type MatchReportData struct {
	MatchID              uint              `json:"match_id"`
	MatchDate            string            `json:"match_date"`
	MatchTime            string            `json:"match_time"`
	HomeTeam             *models.Team      `json:"home_team"`
	AwayTeam             *models.Team      `json:"away_team"`
	HomeScore            int               `json:"home_score"`
	AwayScore            int               `json:"away_score"`
	FinalStatus          string            `json:"final_status"`
	Goals                []models.Goal     `json:"goals"`
	TopScorers           []TopScorer       `json:"top_scorers"`
	HomeTeamTotalWins    int64             `json:"home_team_total_wins"`
	AwayTeamTotalWins    int64             `json:"away_team_total_wins"`
}

// GetMatchReport godoc
// GET /api/reports/matches/:id
func GetMatchReport(c *gin.Context) {
	matchID := c.Param("id")

	// Load match with teams
	var match models.Match
	if err := config.DB.
		Preload("HomeTeam").
		Preload("AwayTeam").
		First(&match, matchID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match not found")
		return
	}

	if match.Status != models.MatchStatusCompleted {
		utils.ErrorResponse(c, http.StatusBadRequest, "Match has not been completed yet")
		return
	}

	// Load result with goals
	var result models.MatchResult
	if err := config.DB.
		Preload("Goals").
		Preload("Goals.Player").
		Where("match_id = ?", match.ID).
		First(&result).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match result not found")
		return
	}

	// Determine final status
	var finalStatus string
	switch {
	case result.HomeScore > result.AwayScore:
		finalStatus = "Tim Home Menang"
	case result.AwayScore > result.HomeScore:
		finalStatus = "Tim Away Menang"
	default:
		finalStatus = "Draw"
	}

	// Calculate top scorers for this match
	scorerMap := make(map[uint]*TopScorer)
	for _, goal := range result.Goals {
		if _, exists := scorerMap[goal.PlayerID]; !exists {
			name := ""
			if goal.Player != nil {
				name = goal.Player.Name
			}
			scorerMap[goal.PlayerID] = &TopScorer{
				PlayerID:   goal.PlayerID,
				PlayerName: name,
				Goals:      0,
			}
		}
		scorerMap[goal.PlayerID].Goals++
	}

	// Find max goals
	maxGoals := 0
	for _, s := range scorerMap {
		if s.Goals > maxGoals {
			maxGoals = s.Goals
		}
	}

	var topScorers []TopScorer
	for _, s := range scorerMap {
		if s.Goals == maxGoals {
			topScorers = append(topScorers, *s)
		}
	}

	// Accumulate home team wins: all completed matches up to this match_id where home team won
	// Win for home: home_score > away_score in matches where home_team_id = match.HomeTeamID
	var homeTeamWins int64
	config.DB.Model(&models.Match{}).
		Joins("JOIN match_results ON match_results.match_id = matches.id AND match_results.deleted_at IS NULL").
		Where("matches.id <= ? AND matches.deleted_at IS NULL AND matches.status = ? AND matches.home_team_id = ? AND match_results.home_score > match_results.away_score",
			match.ID, models.MatchStatusCompleted, match.HomeTeamID).
		Count(&homeTeamWins)

	// Also count when the same team was away and won
	var homeTeamWinsAsAway int64
	config.DB.Model(&models.Match{}).
		Joins("JOIN match_results ON match_results.match_id = matches.id AND match_results.deleted_at IS NULL").
		Where("matches.id <= ? AND matches.deleted_at IS NULL AND matches.status = ? AND matches.away_team_id = ? AND match_results.away_score > match_results.home_score",
			match.ID, models.MatchStatusCompleted, match.HomeTeamID).
		Count(&homeTeamWinsAsAway)

	homeTeamTotalWins := homeTeamWins + homeTeamWinsAsAway

	// Accumulate away team wins
	var awayTeamWins int64
	config.DB.Model(&models.Match{}).
		Joins("JOIN match_results ON match_results.match_id = matches.id AND match_results.deleted_at IS NULL").
		Where("matches.id <= ? AND matches.deleted_at IS NULL AND matches.status = ? AND matches.home_team_id = ? AND match_results.home_score > match_results.away_score",
			match.ID, models.MatchStatusCompleted, match.AwayTeamID).
		Count(&awayTeamWins)

	var awayTeamWinsAsAway int64
	config.DB.Model(&models.Match{}).
		Joins("JOIN match_results ON match_results.match_id = matches.id AND match_results.deleted_at IS NULL").
		Where("matches.id <= ? AND matches.deleted_at IS NULL AND matches.status = ? AND matches.away_team_id = ? AND match_results.away_score > match_results.home_score",
			match.ID, models.MatchStatusCompleted, match.AwayTeamID).
		Count(&awayTeamWinsAsAway)

	awayTeamTotalWins := awayTeamWins + awayTeamWinsAsAway

	report := MatchReportData{
		MatchID:           match.ID,
		MatchDate:         match.MatchDate,
		MatchTime:         match.MatchTime,
		HomeTeam:          match.HomeTeam,
		AwayTeam:          match.AwayTeam,
		HomeScore:         result.HomeScore,
		AwayScore:         result.AwayScore,
		FinalStatus:       finalStatus,
		Goals:             result.Goals,
		TopScorers:        topScorers,
		HomeTeamTotalWins: homeTeamTotalWins,
		AwayTeamTotalWins: awayTeamTotalWins,
	}

	utils.SuccessResponse(c, http.StatusOK, "Match report retrieved successfully", report)
}

// GetAllReports godoc
// GET /api/reports/matches
func GetAllReports(c *gin.Context) {
	var matches []models.Match
	config.DB.
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("MatchResult").
		Preload("MatchResult.Goals").
		Preload("MatchResult.Goals.Player").
		Where("status = ?", models.MatchStatusCompleted).
		Order("match_date ASC, match_time ASC").
		Find(&matches)

	type ReportSummary struct {
		MatchID     uint         `json:"match_id"`
		MatchDate   string       `json:"match_date"`
		MatchTime   string       `json:"match_time"`
		HomeTeam    *models.Team `json:"home_team"`
		AwayTeam    *models.Team `json:"away_team"`
		HomeScore   int          `json:"home_score"`
		AwayScore   int          `json:"away_score"`
		FinalStatus string       `json:"final_status"`
	}

	var reports []ReportSummary
	for _, m := range matches {
		if m.MatchResult == nil {
			continue
		}
		status := "Draw"
		if m.MatchResult.HomeScore > m.MatchResult.AwayScore {
			status = "Tim Home Menang"
		} else if m.MatchResult.AwayScore > m.MatchResult.HomeScore {
			status = "Tim Away Menang"
		}
		reports = append(reports, ReportSummary{
			MatchID:     m.ID,
			MatchDate:   m.MatchDate,
			MatchTime:   m.MatchTime,
			HomeTeam:    m.HomeTeam,
			AwayTeam:    m.AwayTeam,
			HomeScore:   m.MatchResult.HomeScore,
			AwayScore:   m.MatchResult.AwayScore,
			FinalStatus: status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Reports retrieved successfully",
		"data":    reports,
		"total":   len(reports),
	})
}
