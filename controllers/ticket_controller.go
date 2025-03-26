package controllers

import (
	"net/http"

	"movieTicket/models"
	"movieTicket/services"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns a simple success message
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is up and running"})
}

// BookTicket handles booking a movie ticket
func BookTicket(c *gin.Context) {
	var request models.BookTicketRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := services.BookTicketService(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket booked successfully", "ticket": ticket})
}

// ViewTicket retrieves ticket details by email
func ViewTicket(c *gin.Context) {
	email := c.Query("email")
	ticket, err := services.ViewTicketService(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

// ViewAttendees returns a list of attendees for a movie and showtime
func ViewAttendees(c *gin.Context) {
	movieTitle := c.Query("movie_title")
	showtime := c.Query("showtime")

	attendees, err := services.ViewAttendeesService(movieTitle, showtime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}

// CancelTicket handles ticket cancellation
func CancelTicket(c *gin.Context) {
	var request models.CancelTicketRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CancelTicketService(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket successfully canceled."})
}

// ModifySeat handles seat modification requests
func ModifySeat(c *gin.Context) {
	var request models.ModifySeatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.ModifySeatService(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat updated successfully", "new_seat_number": request.NewSeatNumber})
}
