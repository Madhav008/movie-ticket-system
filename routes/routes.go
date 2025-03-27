package routes

import (
	"movieTicket/controllers"
	"movieTicket/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all API routes for the movie ticket booking system
func SetupRoutes(router *gin.Engine, service services.ServiceInterface) {
	ctrl := controllers.NewController(service)

	// Home route
	router.GET("/", ctrl.HealthCheck)

	// Book Movie Ticket API
	router.POST("/api/book-ticket", ctrl.BookTicket)

	// View Movie Ticket Details API
	router.GET("/api/view-ticket", ctrl.ViewTicket)

	// View All Attendees for a Movie API
	router.GET("/api/view-attendees", ctrl.ViewAttendees)

	// Cancel Movie Ticket API
	router.DELETE("/api/cancel-ticket", ctrl.CancelTicket)

	// Modify Seat Assignment API
	router.PUT("/api/modify-seat", ctrl.ModifySeat)
}
