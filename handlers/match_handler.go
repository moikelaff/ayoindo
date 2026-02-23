package handlers

import (
	"net/http"

	"ayoindo/config"
	"ayoindo/models"
	"ayoindo/utils"

	"github.com/gin-gonic/gin"
)

type MatchInput struct {
	HomeTeamID uint   `json:"home_team_id" binding:"required"`
	AwayTeamID uint   `json:"away_team_id" binding:"required"`
	MatchDate  string `json:"match_date" binding:"required"` // YYYY-MM-DD
	MatchTime  string `json:"match_time" binding:"required"` // HH:MM
}

// GetAllMatches godoc
// GET /api/matches
func GetAllMatches(c *gin.Context) {
	var matches []models.Match
	query := config.DB.Preload("HomeTeam").Preload("AwayTeam").Preload("MatchResult").Preload("MatchResult.Goals").Preload("MatchResult.Goals.Player")

	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	config.DB.Model(&models.Match{}).Count(&total)
	query.Order("match_date ASC, match_time ASC").Find(&matches)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Matches retrieved successfully",
		"data":    matches,
		"total":   total,
	})
}

// GetMatchByID godoc
// GET /api/matches/:id
func GetMatchByID(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := config.DB.
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("MatchResult").
		Preload("MatchResult.Goals").
		Preload("MatchResult.Goals.Player").
		First(&match, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Match retrieved successfully", match)
}

// CreateMatch godoc
// POST /api/matches
func CreateMatch(c *gin.Context) {
	var input MatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	if input.HomeTeamID == input.AwayTeamID {
		utils.ValidationErrorResponse(c, "Home team and away team cannot be the same")
		return
	}

	// Validate teams exist
	var homeTeam, awayTeam models.Team
	if err := config.DB.First(&homeTeam, input.HomeTeamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Home team not found")
		return
	}
	if err := config.DB.First(&awayTeam, input.AwayTeamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Away team not found")
		return
	}

	match := models.Match{
		HomeTeamID: input.HomeTeamID,
		AwayTeamID: input.AwayTeamID,
		MatchDate:  input.MatchDate,
		MatchTime:  input.MatchTime,
		Status:     models.MatchStatusScheduled,
	}

	if err := config.DB.Create(&match).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create match")
		return
	}

	config.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, match.ID)
	utils.SuccessResponse(c, http.StatusCreated, "Match created successfully", match)
}

// UpdateMatch godoc
// PUT /api/matches/:id
func UpdateMatch(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := config.DB.First(&match, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match not found")
		return
	}

	if match.Status == models.MatchStatusCompleted {
		utils.ErrorResponse(c, http.StatusBadRequest, "Cannot update a completed match schedule")
		return
	}

	var input MatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	if input.HomeTeamID == input.AwayTeamID {
		utils.ValidationErrorResponse(c, "Home team and away team cannot be the same")
		return
	}

	var homeTeam, awayTeam models.Team
	if err := config.DB.First(&homeTeam, input.HomeTeamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Home team not found")
		return
	}
	if err := config.DB.First(&awayTeam, input.AwayTeamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Away team not found")
		return
	}

	match.HomeTeamID = input.HomeTeamID
	match.AwayTeamID = input.AwayTeamID
	match.MatchDate = input.MatchDate
	match.MatchTime = input.MatchTime

	if err := config.DB.Save(&match).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update match")
		return
	}

	config.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, match.ID)
	utils.SuccessResponse(c, http.StatusOK, "Match updated successfully", match)
}

// DeleteMatch godoc
// DELETE /api/matches/:id — soft delete
func DeleteMatch(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := config.DB.First(&match, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Match not found")
		return
	}

	if err := config.DB.Delete(&match).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete match")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Match deleted successfully", nil)
}
