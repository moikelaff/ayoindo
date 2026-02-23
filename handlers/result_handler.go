package handlers

import (
	"net/http"

	"ayoindo/config"
	"ayoindo/models"
	"ayoindo/utils"

	"github.com/gin-gonic/gin"
)

type GoalInput struct {
	PlayerID uint `json:"player_id" binding:"required"`
	Minute   int  `json:"minute" binding:"required,min=1,max=120"`
}

type MatchResultInput struct {
	HomeScore int         `json:"home_score" binding:"min=0"`
	AwayScore int         `json:"away_score" binding:"min=0"`
	Goals     []GoalInput `json:"goals"`
}

// SubmitMatchResult godoc
// POST /api/matches/:id/result
func SubmitMatchResult(c *gin.Context) {
	matchID := c.Param("id")

	var match models.Match
	if err := config.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, matchID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match not found")
		return
	}

	var input MatchResultInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	// Check if result already exists — update if so
	var existingResult models.MatchResult
	resultExists := config.DB.Where("match_id = ?", match.ID).First(&existingResult).Error == nil

	// Validate goals: each player must belong to one of the two teams
	// and goal count must match scores
	homeGoalCount := 0
	awayGoalCount := 0

	for _, g := range input.Goals {
		var player models.Player
		if err := config.DB.First(&player, g.PlayerID).Error; err != nil {
			utils.ErrorResponse(c, http.StatusNotFound, "Player not found: player_id "+string(rune(g.PlayerID)))
			return
		}
		if player.TeamID != match.HomeTeamID && player.TeamID != match.AwayTeamID {
			utils.ValidationErrorResponse(c, "Player does not belong to either team in this match")
			return
		}
		if player.TeamID == match.HomeTeamID {
			homeGoalCount++
		} else {
			awayGoalCount++
		}
	}

	if homeGoalCount != input.HomeScore || awayGoalCount != input.AwayScore {
		utils.ValidationErrorResponse(c, "Number of goals does not match the provided scores")
		return
	}

	// Use a transaction
	tx := config.DB.Begin()

	var result models.MatchResult
	if resultExists {
		// Delete old goals first
		tx.Where("match_result_id = ?", existingResult.ID).Delete(&models.Goal{})
		existingResult.HomeScore = input.HomeScore
		existingResult.AwayScore = input.AwayScore
		if err := tx.Save(&existingResult).Error; err != nil {
			tx.Rollback()
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update match result")
			return
		}
		result = existingResult
	} else {
		result = models.MatchResult{
			MatchID:   match.ID,
			HomeScore: input.HomeScore,
			AwayScore: input.AwayScore,
		}
		if err := tx.Create(&result).Error; err != nil {
			tx.Rollback()
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create match result")
			return
		}
	}

	// Insert goals
	for _, g := range input.Goals {
		goal := models.Goal{
			MatchResultID: result.ID,
			PlayerID:      g.PlayerID,
			Minute:        g.Minute,
		}
		if err := tx.Create(&goal).Error; err != nil {
			tx.Rollback()
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to save goal")
			return
		}
	}

	// Mark match as completed
	match.Status = models.MatchStatusCompleted
	if err := tx.Save(&match).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update match status")
		return
	}

	tx.Commit()

	// Reload with associations
	config.DB.Preload("Goals").Preload("Goals.Player").First(&result, result.ID)

	utils.SuccessResponse(c, http.StatusOK, "Match result submitted successfully", result)
}

// GetMatchResult godoc
// GET /api/matches/:id/result
func GetMatchResult(c *gin.Context) {
	matchID := c.Param("id")

	var match models.Match
	if err := config.DB.First(&match, matchID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match not found")
		return
	}

	var result models.MatchResult
	if err := config.DB.
		Preload("Goals").
		Preload("Goals.Player").
		Where("match_id = ?", matchID).
		First(&result).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "No result found for this match")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Match result retrieved successfully", result)
}
