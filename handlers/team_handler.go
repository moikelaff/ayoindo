package handlers

import (
	"net/http"

	"ayoindo/config"
	"ayoindo/models"
	"ayoindo/utils"

	"github.com/gin-gonic/gin"
)

type TeamInput struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Logo        string `json:"logo"`
	FoundedYear int    `json:"founded_year" binding:"required,min=1800,max=2100"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
}

// GetAllTeams godoc
// GET /api/teams
func GetAllTeams(c *gin.Context) {
	var teams []models.Team
	query := config.DB.Model(&models.Team{})

	// Optional filter by city
	if city := c.Query("city"); city != "" {
		query = query.Where("city ILIKE ?", "%"+city+"%")
	}

	var total int64
	query.Count(&total)
	query.Find(&teams)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Teams retrieved successfully",
		"data":    teams,
		"total":   total,
	})
}

// GetTeamByID godoc
// GET /api/teams/:id
func GetTeamByID(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := config.DB.Preload("Players").First(&team, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Team not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Team retrieved successfully", team)
}

// CreateTeam godoc
// POST /api/teams
func CreateTeam(c *gin.Context) {
	var input TeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	team := models.Team{
		Name:        input.Name,
		Logo:        input.Logo,
		FoundedYear: input.FoundedYear,
		Address:     input.Address,
		City:        input.City,
	}

	if err := config.DB.Create(&team).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create team")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Team created successfully", team)
}

// UpdateTeam godoc
// PUT /api/teams/:id
func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := config.DB.First(&team, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Team not found")
		return
	}

	var input TeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	team.Name = input.Name
	team.Logo = input.Logo
	team.FoundedYear = input.FoundedYear
	team.Address = input.Address
	team.City = input.City

	if err := config.DB.Save(&team).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update team")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Team updated successfully", team)
}

// DeleteTeam godoc
// DELETE /api/teams/:id — soft delete via GORM
func DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := config.DB.First(&team, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Team not found")
		return
	}

	if err := config.DB.Delete(&team).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete team")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Team deleted successfully", nil)
}
