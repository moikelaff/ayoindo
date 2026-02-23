package handlers

import (
	"net/http"

	"ayoindo/config"
	"ayoindo/models"
	"ayoindo/utils"

	"github.com/gin-gonic/gin"
)

type PlayerInput struct {
	TeamID       uint                  `json:"team_id" binding:"required"`
	Name         string                `json:"name" binding:"required,min=2,max=100"`
	Height       float64               `json:"height" binding:"required,min=100,max=250"`
	Weight       float64               `json:"weight" binding:"required,min=30,max=200"`
	Position     models.PlayerPosition `json:"position" binding:"required"`
	JerseyNumber int                   `json:"jersey_number" binding:"required,min=1,max=99"`
}

func isValidPosition(pos models.PlayerPosition) bool {
	switch pos {
	case models.PositionPenyerang, models.PositionGelandang,
		models.PositionBertahan, models.PositionPenjagaGawang:
		return true
	}
	return false
}

// GetAllPlayers godoc
// GET /api/players
func GetAllPlayers(c *gin.Context) {
	var players []models.Player
	query := config.DB.Preload("Team")

	if teamID := c.Query("team_id"); teamID != "" {
		query = query.Where("team_id = ?", teamID)
	}
	if pos := c.Query("position"); pos != "" {
		query = query.Where("position = ?", pos)
	}

	var total int64
	config.DB.Model(&models.Player{}).Count(&total)
	query.Find(&players)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Players retrieved successfully",
		"data":    players,
		"total":   total,
	})
}

// GetPlayerByID godoc
// GET /api/players/:id
func GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	var player models.Player

	if err := config.DB.Preload("Team").First(&player, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Player not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Player retrieved successfully", player)
}

// GetPlayersByTeam godoc
// GET /api/teams/:id/players
func GetPlayersByTeam(c *gin.Context) {
	teamID := c.Param("id")
	var team models.Team
	if err := config.DB.First(&team, teamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Team not found")
		return
	}

	var players []models.Player
	config.DB.Where("team_id = ?", teamID).Find(&players)

	utils.SuccessResponse(c, http.StatusOK, "Players retrieved successfully", players)
}

// CreatePlayer godoc
// POST /api/players
func CreatePlayer(c *gin.Context) {
	var input PlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	// Validate position enum
	if !isValidPosition(input.Position) {
		utils.ValidationErrorResponse(c, "Invalid position. Must be one of: penyerang, gelandang, bertahan, penjaga_gawang")
		return
	}

	// Validate team exists
	var team models.Team
	if err := config.DB.First(&team, input.TeamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Team not found")
		return
	}

	// Check jersey number uniqueness within the team
	var existing models.Player
	if err := config.DB.Where("team_id = ? AND jersey_number = ?", input.TeamID, input.JerseyNumber).
		First(&existing).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Jersey number already taken in this team")
		return
	}

	player := models.Player{
		TeamID:       input.TeamID,
		Name:         input.Name,
		Height:       input.Height,
		Weight:       input.Weight,
		Position:     input.Position,
		JerseyNumber: input.JerseyNumber,
	}

	if err := config.DB.Create(&player).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create player")
		return
	}

	config.DB.Preload("Team").First(&player, player.ID)
	utils.SuccessResponse(c, http.StatusCreated, "Player created successfully", player)
}

// UpdatePlayer godoc
// PUT /api/players/:id
func UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player

	if err := config.DB.First(&player, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Player not found")
		return
	}

	var input PlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	if !isValidPosition(input.Position) {
		utils.ValidationErrorResponse(c, "Invalid position. Must be one of: penyerang, gelandang, bertahan, penjaga_gawang")
		return
	}

	// Validate team exists
	var team models.Team
	if err := config.DB.First(&team, input.TeamID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Team not found")
		return
	}

	// Check jersey uniqueness within team, excluding this player
	var existing models.Player
	if err := config.DB.Where("team_id = ? AND jersey_number = ? AND id != ?",
		input.TeamID, input.JerseyNumber, player.ID).First(&existing).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Jersey number already taken in this team")
		return
	}

	player.TeamID = input.TeamID
	player.Name = input.Name
	player.Height = input.Height
	player.Weight = input.Weight
	player.Position = input.Position
	player.JerseyNumber = input.JerseyNumber

	if err := config.DB.Save(&player).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update player")
		return
	}

	config.DB.Preload("Team").First(&player, player.ID)
	utils.SuccessResponse(c, http.StatusOK, "Player updated successfully", player)
}

// DeletePlayer godoc
// DELETE /api/players/:id — soft delete
func DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	var player models.Player

	if err := config.DB.First(&player, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Player not found")
		return
	}

	if err := config.DB.Delete(&player).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete player")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Player deleted successfully", nil)
}
