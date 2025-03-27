package repository

import (
	"errors"
	"fmt"
	"log"
	"movieTicket/config"
	"movieTicket/models"
	"sync"
	"time"
)

// MovieTicketRepository struct for handling ticket-related DB operations
type MovieTicketRepository struct{}

var TicketRepo = MovieTicketRepository{}

// In-memory fallback storage
var (
	tickets      = make(map[string]models.Ticket) // Key: (email + showtime)
	ticketsMutex sync.Mutex
)

// NewMovieTicketRepository returns a new instance of MovieTicketRepository
func NewMovieTicketRepository() *MovieTicketRepository {
	return &MovieTicketRepository{}
}

// BookTicket saves a new movie ticket to the database or memory
func (r *MovieTicketRepository) BookTicket(ticket *models.Ticket) error {
	if config.DBAvailable {
		// Check if the user already booked for the same movie and showtime
		var count int64
		if err := config.DB.Model(&models.Ticket{}).
			Where("email = ? AND movie_title = ? AND showtime = ?", ticket.Email, ticket.MovieTitle, ticket.Showtime).
			Count(&count).Error; err != nil {
			config.DBAvailable = false
			log.Println("⚠️  Database error, switching to in-memory mode")
		}

		if count > 0 {
			return errors.New("email already booked for the same showtime and movie")
		}

		// Ensure seats exist
		var seatCount int64
		config.DB.Model(&models.Seat{}).Where("movie_title = ? AND showtime = ?", ticket.MovieTitle, ticket.Showtime).Count(&seatCount)

		if seatCount == 0 {
			if err := CreateSeatsForShowtime(ticket.MovieTitle, ticket.Showtime, 50); err != nil {
				return errors.New("failed to create seats for the new showtime")
			}
		}

		// Get the next available seat
		seatNumber, err := FindNextAvailableSeat(ticket.MovieTitle, ticket.Showtime)
		if err != nil {
			return errors.New("no available seats")
		}

		// Create ticket entry
		ticket.SeatNumber = seatNumber
		ticket.Status = "Confirmed"
		ticket.CreatedAt = time.Now()
		ticket.UpdatedAt = time.Now()

		if err := config.DB.Model(&models.Ticket{}).Create(ticket).Error; err != nil {
			config.DBAvailable = false
			log.Println("⚠️  Database error, saving ticket in-memory instead")
		} else {
			// Mark seat as booked
			config.DB.Model(&models.Seat{}).Where("seat_number = ? AND showtime = ?", seatNumber, ticket.Showtime).
				Update("is_booked", true)
			return nil
		}
	}

	// In-memory storage fallback
	ticketsMutex.Lock()
	defer ticketsMutex.Unlock()

	key := ticket.Email + ticket.Showtime
	if _, exists := tickets[key]; exists {
		return errors.New("email already booked for this showtime and movie")
	}

	tickets[key] = *ticket
	log.Println("✅ Ticket booked in-memory due to DB failure")

	return nil
}

// CreateSeatsForShowtime initializes seats for a new movie showtime
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

// FindNextAvailableSeat finds the next available seat for a given showtime
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
	if config.DBAvailable {
		var tickets []models.Ticket
		err := config.DB.Where("email = ?", email).Find(&tickets).Error
		if err == nil {
			return tickets, nil
		}
		config.DBAvailable = false
		log.Println("⚠️  Database error, switching to in-memory mode")
	}

	// In-memory fallback
	ticketsMutex.Lock()
	defer ticketsMutex.Unlock()

	var results []models.Ticket
	for _, ticket := range tickets {
		if ticket.Email == email {
			results = append(results, ticket)
		}
	}

	if len(results) == 0 {
		return nil, errors.New("no tickets found")
	}

	return results, nil
}

// GetAttendeesByMovie retrieves all attendees for a specific movie and showtime
func (r *MovieTicketRepository) GetAttendeesByMovie(movieTitle, showtime string) ([]models.Attendees, error) {
	if config.DBAvailable {
		var attendees []models.Attendees
		err := config.DB.Model(&models.Ticket{}).
			Where("movie_title = ? AND showtime = ?", movieTitle, showtime).
			Find(&attendees).Error
		if err == nil {
			return attendees, nil
		}
		config.DBAvailable = false
		log.Println("⚠️  Database error, switching to in-memory mode")
	}

	// In-memory fallback
	ticketsMutex.Lock()
	defer ticketsMutex.Unlock()

	var attendees []models.Attendees
	for _, ticket := range tickets {
		if ticket.MovieTitle == movieTitle && ticket.Showtime == showtime {
			attendees = append(attendees, models.Attendees{
				Name:       ticket.Name,
				SeatNumber: ticket.SeatNumber,
			})
		}
	}

	if len(attendees) == 0 {
		return nil, errors.New("no attendees found")
	}

	return attendees, nil
}

// CancelTicket deletes a ticket by email and showtime
func (r *MovieTicketRepository) CancelTicket(email, showtime string) error {
	if config.DBAvailable {
		err := config.DB.Where("email = ? AND showtime = ?", email, showtime).Delete(&models.Ticket{}).Error
		if err == nil {
			return nil
		}
		config.DBAvailable = false
		log.Println("⚠️  Database error, switching to in-memory mode")
	}

	// In-memory fallback
	ticketsMutex.Lock()
	defer ticketsMutex.Unlock()

	key := email + showtime
	if _, exists := tickets[key]; exists {
		delete(tickets, key)
		log.Println("✅ Ticket canceled from in-memory storage")
		return nil
	}

	return errors.New("ticket not found")
}

// ModifySeat updates the seat assignment for a ticket
func (r *MovieTicketRepository) ModifySeat(email, showtime, newSeat string) error {
	if config.DBAvailable {
		err := config.DB.Model(&models.Ticket{}).
			Where("email = ? AND showtime = ?", email, showtime).
			Update("seat_number", newSeat).Error
		if err == nil {
			return nil
		}
		config.DBAvailable = false
		log.Println("⚠️  Database error, switching to in-memory mode")
	}

	// In-memory fallback
	ticketsMutex.Lock()
	defer ticketsMutex.Unlock()

	key := email + showtime
	if ticket, exists := tickets[key]; exists {
		ticket.SeatNumber = newSeat
		tickets[key] = ticket
		log.Println("✅ Seat modified in-memory")
		return nil
	}

	return errors.New("ticket not found")
}
