package services

import (
	"errors"
	"time"

	"movieTicket/models"
	"movieTicket/repository"

	"github.com/google/uuid"
)

// BookTicketService handles business logic for booking a ticket
func BookTicketService(request models.BookTicketRequest) (models.TicketConfirmation, error) {
	// Validate input data
	if request.Name == "" || request.Email == "" || request.MovieTitle == "" || request.Showtime == "" {
		return models.TicketConfirmation{}, errors.New("all fields are required")
	}

	// Generate ticket ID and seat number
	ticketID := uuid.New().String()
	seatNumber := "A" + ticketID[len(ticketID)-2:] // Mock seat allocation logic

	// Create a ticket object (replace with DB persistence logic)
	ticket := models.Ticket{
		ID:         uint(time.Now().Unix()), // Mock ID, replace with DB-generated ID
		Name:       request.Name,
		Email:      request.Email,
		MovieTitle: request.MovieTitle,
		Showtime:   request.Showtime,
		SeatNumber: seatNumber,
		Status:     "Confirmed",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := repository.TicketRepo.BookTicket(&ticket)
	if err != nil {
		return models.TicketConfirmation{}, err
	}

	return models.TicketConfirmation{
		Name:       ticket.Name,
		Email:      ticket.Email,
		MovieTitle: ticket.MovieTitle,
		Showtime:   ticket.Showtime,
		SeatNumber: ticket.SeatNumber,
		Status:     ticket.Status,
	}, nil
}

// ViewTicketService retrieves a ticket by email (Mock function)
func ViewTicketService(email string) ([]models.Ticket, error) {
	if email == "" {
		return []models.Ticket{}, errors.New("email is required")
	}

	ticket, err := repository.TicketRepo.GetTicketByEmail(email)
	if err != nil {
		return []models.Ticket{}, err
	}

	return ticket, nil
}

// ViewAttendeesService retrieves all attendees for a movie showtime (Mock function)
func ViewAttendeesService(movieTitle, showtime string) ([]models.Attendees, error) {
	if movieTitle == "" || showtime == "" {
		return nil, errors.New("movie title and showtime are required")
	}

	// Fetch attendees from repository
	attendees, err := repository.TicketRepo.GetAttendeesByMovie(movieTitle, showtime)
	if err != nil {
		return nil, err
	}

	return attendees, nil
}

// CancelTicketService cancels a ticket based on email and showtime
func CancelTicketService(request models.CancelTicketRequest) error {
	if request.Email == "" || request.Showtime == "" {
		return errors.New("email and showtime are required")
	}

	// Cancel ticket using repository
	err := repository.TicketRepo.CancelTicket(request.Email, request.Showtime)
	if err != nil {
		return err
	}

	return nil
}

// ModifySeatService updates seat assignment
func ModifySeatService(request models.ModifySeatRequest) error {
	if request.Email == "" || request.Showtime == "" || request.NewSeatNumber == "" {
		return errors.New("email, showtime, and new seat number are required")
	}

	// Modify seat using repository
	err := repository.TicketRepo.ModifySeat(request.Email, request.Showtime, request.NewSeatNumber)
	if err != nil {
		return err
	}

	return nil
}
