package routes

import (
	"ayoindo/handlers"
	"ayoindo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// ─── Auth (public) ────────────────────────────────────────────────
	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// ─── Protected routes ─────────────────────────────────────────────
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Profile
		protected.GET("auth/me", handlers.GetProfile)

		// Teams
		teams := protected.Group("/teams")
		{
			teams.GET("", handlers.GetAllTeams)
			teams.POST("", handlers.CreateTeam)
			teams.GET("/:id", handlers.GetTeamByID)
			teams.PUT("/:id", handlers.UpdateTeam)
			teams.DELETE("/:id", handlers.DeleteTeam)
			teams.GET("/:id/players", handlers.GetPlayersByTeam)
		}

		// Players
		players := protected.Group("/players")
		{
			players.GET("", handlers.GetAllPlayers)
			players.POST("", handlers.CreatePlayer)
			players.GET("/:id", handlers.GetPlayerByID)
			players.PUT("/:id", handlers.UpdatePlayer)
			players.DELETE("/:id", handlers.DeletePlayer)
		}

		// Matches
		matches := protected.Group("/matches")
		{
			matches.GET("", handlers.GetAllMatches)
			matches.POST("", handlers.CreateMatch)
			matches.GET("/:id", handlers.GetMatchByID)
			matches.PUT("/:id", handlers.UpdateMatch)
			matches.DELETE("/:id", handlers.DeleteMatch)

			// Match Result
			matches.POST("/:id/result", handlers.SubmitMatchResult)
			matches.GET("/:id/result", handlers.GetMatchResult)
		}

		// Reports
		reports := protected.Group("/reports")
		{
			reports.GET("/matches", handlers.GetAllReports)
			reports.GET("/matches/:id", handlers.GetMatchReport)
		}
	}
}
