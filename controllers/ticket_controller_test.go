package controllers

import (
	"bytes"
	"encoding/json"
	"movieTicket/models"
	"movieTicket/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBookTicket(t *testing.T) {
	mockService := &services.MockMovieTicketService{}
	controller := NewController(mockService)
    gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/book-ticket", controller.BookTicket)

	requestBody := models.BookTicketRequest{
		Name:       "John Doe",
		Email:      "newuser@example.com",
		MovieTitle: "Avengers",
		Showtime:   "7:00 PM",
	}

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/book-ticket", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ticket booked successfully")
}

func TestViewTicket(t *testing.T) {
	mockService := &services.MockMovieTicketService{}
	controller := NewController(mockService)
    gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/view-ticket", controller.ViewTicket)

	req, _ := http.NewRequest("GET", "/view-ticket?email=test@example.com", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestViewAttendees(t *testing.T) {
	mockService := &services.MockMovieTicketService{}
	controller := NewController(mockService)
    gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/view-attendees", controller.ViewAttendees)

	req, _ := http.NewRequest("GET", "/view-attendees?movie_title=Avengers&showtime=7:00PM", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCancelTicket(t *testing.T) {
	mockService := &services.MockMovieTicketService{}
	controller := NewController(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.DELETE("/cancel-ticket", controller.CancelTicket)

	requestBody := models.CancelTicketRequest{
		Email:    "test@example.com",
		Showtime: "7:00 PM",
	}

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("DELETE", "/cancel-ticket", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ticket successfully canceled.")
}

func TestModifySeat(t *testing.T) {
	mockService := &services.MockMovieTicketService{}
	controller := NewController(mockService)
    gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.PUT("/modify-seat", controller.ModifySeat)

	requestBody := models.ModifySeatRequest{
		Email:         "test@example.com",
		Showtime:      "7:00 PM",
		NewSeatNumber: "B12",
	}

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/modify-seat", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Seat updated successfully")
}