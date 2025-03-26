package routes

import (
	"movieTicket/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all API routes for the movie ticket booking system
func SetupRoutes(router *gin.Engine) {
	// Home route
	router.GET("/", controllers.HealthCheck)

	// Book Movie Ticket API
	router.POST("/api/book-ticket", controllers.BookTicket)

	// View Movie Ticket Details API
	router.GET("/api/view-ticket", controllers.ViewTicket)

	// View All Attendees for a Movie API
	router.GET("/api/view-attendees", controllers.ViewAttendees)

	// Cancel Movie Ticket API
	router.DELETE("/api/cancel-ticket", controllers.CancelTicket)

	// Modify Seat Assignment API
	router.PUT("/api/modify-seat", controllers.ModifySeat)
}
