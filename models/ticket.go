package models

import "time"

// Ticket represents a movie ticket booking
type Ticket struct {
	ID         uint      `json:"id"`          // Unique identifier for the ticket
	Name       string    `json:"name"`        // User's name
	Email      string    `json:"email"`       // User's email
	MovieTitle string    `json:"movie_title"` // Title of the movie
	Showtime   string    `json:"showtime"`    // Showtime of the movie
	SeatNumber string    `json:"seat_number"` // Assigned seat number
	Status     string    `json:"status"`      // Status of the booking (e.g., Confirmed, Cancelled)
	CreatedAt  time.Time `json:"created_at"`  // Timestamp of ticket creation
	UpdatedAt  time.Time `json:"updated_at"`  // Timestamp of last update
}

// Seat represents a seat in a theater
type Seat struct {
	ID         uint      `json:"id"`          // Unique identifier for the seat
	MovieTitle string    `json:"movie_title"` // Movie associated with the seat
	Showtime   string    `json:"showtime"`    // Showtime for which the seat is reserved
	SeatNumber string    `json:"seat_number"` // Unique seat number
	IsBooked   bool      `json:"is_booked"`   // Indicates if the seat is booked
	CreatedAt  time.Time `json:"created_at"`  // Timestamp of seat creation
	UpdatedAt  time.Time `json:"updated_at"`  // Timestamp of last update
}

// Request models for API calls

// BookTicketRequest represents the request body for booking a ticket
type BookTicketRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	MovieTitle string `json:"movie_title" binding:"required"`
	Showtime   string `json:"showtime" binding:"required"`
}

// ModifySeatRequest represents the request body for modifying a seat
type ModifySeatRequest struct {
	Email         string `json:"email" binding:"required,email"`
	Showtime      string `json:"showtime" binding:"required"`
	NewSeatNumber string `json:"new_seat_number" binding:"required"`
}

// CancelTicketRequest represents the request body for canceling a ticket
type CancelTicketRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Showtime string `json:"showtime" binding:"required"`
}

type Attendees struct {
	Name       string `json:"name"`        // User's name
	SeatNumber string `json:"seat_number"` // Assigned seat number
}

type TicketConfirmation struct {
	Name       string `json:"name"`        // User's name
	Email      string `json:"email"`       // User's email
	MovieTitle string `json:"movie_title"` // Title of the movie
	Showtime   string `json:"showtime"`    // Showtime of the movie
	SeatNumber string `json:"seat_number"` // Assigned seat number
	Status     string `json:"status"`      // Status of the booking (e.g., Confirmed, Cancelled)
}
