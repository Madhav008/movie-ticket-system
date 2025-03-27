package controllers

import (
	"net/http"

	"movieTicket/models"
	"movieTicket/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service services.ServiceInterface
}

func NewController(service services.ServiceInterface) *Controller {
	return &Controller{service: service}
}

// HealthCheck returns a simple success message
func (ctrl *Controller) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is up and running"})
}

// BookTicket handles booking a movie ticket
func (ctrl *Controller) BookTicket(c *gin.Context) {
	var request models.BookTicketRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket, err := ctrl.service.BookTicketService(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket booked successfully", "ticket": ticket})
}

// ViewTicket retrieves ticket details by email
func (ctrl *Controller) ViewTicket(c *gin.Context) {
	email := c.Query("email")
	ticket, err := ctrl.service.ViewTicketService(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

// ViewAttendees returns a list of attendees for a movie and showtime
func (ctrl *Controller) ViewAttendees(c *gin.Context) {
	movieTitle := c.Query("movie_title")
	showtime := c.Query("showtime")

	attendees, err := ctrl.service.ViewAttendeesService(movieTitle, showtime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}

// CancelTicket handles ticket cancellation
func (ctrl *Controller) CancelTicket(c *gin.Context) {
	var request models.CancelTicketRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CancelTicketService(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket successfully canceled."})
}

// ModifySeat handles seat modification requests
func (ctrl *Controller) ModifySeat(c *gin.Context) {
	var request models.ModifySeatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.ModifySeatService(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat updated successfully", "new_seat_number": request.NewSeatNumber})
}
