package repository

import (
	"errors"
	"fmt"
	"log"
	"movieTicket/config"
	"movieTicket/models"
	"time"
)

// MovieTicketRepository struct for handling ticket-related DB operations
type MovieTicketRepository struct{}

var TicketRepo = MovieTicketRepository{}

// BookTicket saves a new movie ticket to the database
func (r *MovieTicketRepository) BookTicket(ticket *models.Ticket) error {
	var count int64
	if err := config.DB.Model(&models.Ticket{}).
		Where("email = ? AND movie_title = ? AND showtime = ?", ticket.Email, ticket.MovieTitle, ticket.Showtime).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("email already booked for the same showtime and movie")
	}

	// Check if seats exist for this showtime and movie; if not, create them
	var seatCount int64
	config.DB.Model(&models.Seat{}).Where("movie_title = ? AND showtime = ?", ticket.MovieTitle, ticket.Showtime).Count(&seatCount)
	log.Println("---------------------------")
	log.Println(seatCount)
	log.Println("---------------------------")

	if seatCount == 0 {
		// Assuming a fixed number of seats per show (e.g., 50)
		if err := CreateSeatsForShowtime(ticket.MovieTitle, ticket.Showtime, 50); err != nil {
			return errors.New("failed to create seats for the new showtime")
		}
	}

	// Find next available seat
	seatNumber, err := FindNextAvailableSeat(ticket.MovieTitle, ticket.Showtime)
	if err != nil {
		return errors.New("no available seats")
	}
	log.Println("---------------------------")
	log.Println(seatNumber)
	log.Println("---------------------------")

	// Create ticket entry
	_ticket := models.Ticket{
		Name:       ticket.Name,
		Email:      ticket.Email,
		MovieTitle: ticket.MovieTitle,
		Showtime:   ticket.Showtime,
		SeatNumber: seatNumber,
		Status:     "Confirmed",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	log.Println("---------------------------")
	log.Println(_ticket)
	log.Println("---------------------------")

	err = config.DB.Model(&models.Ticket{}).Create(&_ticket).Error
	if err != nil {
		return err
	}
	// Mark seat as booked
	config.DB.Model(&models.Seat{}).Where("seat_number = ? AND showtime = ?", seatNumber, _ticket.Showtime).
		Update("is_booked", true)

	return nil
}

func CreateSeatsForShowtime(movieTitle, showtime string, totalSeats int) error {
	log.Printf("Creating %d seats for %s at %s", totalSeats, movieTitle, showtime)

	var seats []models.Seat
	for i := 1; i <= totalSeats; i++ {
		seats = append(seats, models.Seat{
			MovieTitle: movieTitle,
			Showtime:   showtime,
			SeatNumber: fmt.Sprintf("A%d", i),
			IsBooked:   false,
		})
	}
	return config.DB.Create(&seats).Error
}

func FindNextAvailableSeat(movieTitle, showtime string) (string, error) {
	var seat models.Seat
	err := config.DB.Where("movie_title = ? AND showtime = ? AND is_booked = false", movieTitle, showtime).
		Order("seat_number ASC").
		First(&seat).Error
	if err != nil {
		return "", errors.New("no available seats")
	}
	return seat.SeatNumber, nil
}

// GetTicketByEmail retrieves a ticket by email
func (r *MovieTicketRepository) GetTicketByEmail(email string) ([]models.Ticket, error) {
	var ticket []models.Ticket
	err := config.DB.Where("email = ?", email).First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

// GetAttendeesByMovie retrieves all attendees for a specific movie and showtime
func (r *MovieTicketRepository) GetAttendeesByMovie(movieTitle string, showtime string) ([]models.Attendees, error) {
	var attendees []models.Attendees
	// var ticket []models.Ticket
	// err := config.DB.Where("movie_title = ? AND showtime = ?", movieTitle, showtime).Find(&ticket).Error
	// if err != nil {
	// 	return nil, err
	// }

	// attendees = make([]models.Attendees, len(ticket))
	// for i, t := range ticket {
	// 	attendees[i] = models.Attendees{
	// 		Name:       t.Name,
	// 		Email:      t.Email,
	// 		MovieTitle: t.MovieTitle,
	// 		Showtime:   t.Showtime,
	// 		SeatNumber: t.SeatNumber,
	// 	}
	// }

	err := config.DB.Model(&models.Ticket{}).Where("movie_title = ? AND showtime = ?", movieTitle, showtime).Find(&attendees).Error
	if err != nil {
		return nil, err
	}

	return attendees, nil
}

// CancelTicket deletes a ticket by email and showtime
func (r *MovieTicketRepository) CancelTicket(email string, showtime string) error {
	return config.DB.Where("email = ? AND showtime = ?", email, showtime).Delete(&models.Ticket{}).Error
}

// ModifySeat updates the seat assignment for a ticket
func (r *MovieTicketRepository) ModifySeat(email string, showtime string, newSeat string) error {
	return config.DB.Model(&models.Ticket{}).Where("email = ? AND showtime = ?", email, showtime).
		Update("seat_number", newSeat).Error
}
