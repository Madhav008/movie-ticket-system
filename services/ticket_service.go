package services

import (
	"errors"
	"time"

	"movieTicket/models"
	"movieTicket/repository"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	BookTicketService(request models.BookTicketRequest) (models.TicketConfirmation, error)
	ViewTicketService(email string) ([]models.Ticket, error)
	ViewAttendeesService(movieTitle, showtime string) ([]models.Attendees, error)
	CancelTicketService(request models.CancelTicketRequest) error
	ModifySeatService(request models.ModifySeatRequest) error
}

type MovieTicketService struct {
	repo *repository.MovieTicketRepository
}

type MockMovieTicketService struct{}

func NewMovieTicketService(repo *repository.MovieTicketRepository) *MovieTicketService {
	return &MovieTicketService{repo: repo}
}

func NewMockMovieTicketService() *MockMovieTicketService {
	return &MockMovieTicketService{}
}

// Real Service Implementation
func (s *MovieTicketService) BookTicketService(request models.BookTicketRequest) (models.TicketConfirmation, error) {
	if request.Name == "" || request.Email == "" || request.MovieTitle == "" || request.Showtime == "" {
		return models.TicketConfirmation{}, errors.New("all fields are required")
	}
	ticketID := uuid.New().String()
	seatNumber := "A" + ticketID[len(ticketID)-2:]

	ticket := models.Ticket{
		ID:         uint(time.Now().Unix()),
		Name:       request.Name,
		Email:      request.Email,
		MovieTitle: request.MovieTitle,
		Showtime:   request.Showtime,
		SeatNumber: seatNumber,
		Status:     "Confirmed",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := s.repo.BookTicket(&ticket)
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

func (s *MovieTicketService) ViewTicketService(email string) ([]models.Ticket, error) {
	if email == "" {
		return []models.Ticket{}, errors.New("email is required")
	}
	ticket, err := s.repo.GetTicketByEmail(email)
	if err != nil {
		return []models.Ticket{}, err
	}
	return ticket, nil
}

func (s *MovieTicketService) ViewAttendeesService(movieTitle, showtime string) ([]models.Attendees, error) {
	if movieTitle == "" || showtime == "" {
		return nil, errors.New("movie title and showtime are required")
	}
	attendees, err := s.repo.GetAttendeesByMovie(movieTitle, showtime)
	if err != nil {
		return nil, err
	}
	return attendees, nil
}

func (s *MovieTicketService) CancelTicketService(request models.CancelTicketRequest) error {
	if request.Email == "" || request.Showtime == "" {
		return errors.New("email and showtime are required")
	}
	return s.repo.CancelTicket(request.Email, request.Showtime)
}

func (s *MovieTicketService) ModifySeatService(request models.ModifySeatRequest) error {
	if request.Email == "" || request.Showtime == "" || request.NewSeatNumber == "" {
		return errors.New("email, showtime, and new seat number are required")
	}
	return s.repo.ModifySeat(request.Email, request.Showtime, request.NewSeatNumber)
}

// Mock Service Implementation
func (m *MockMovieTicketService) BookTicketService(request models.BookTicketRequest) (models.TicketConfirmation, error) {
	return models.TicketConfirmation{
		Name:       request.Name,
		Email:      request.Email,
		MovieTitle: request.MovieTitle,
		Showtime:   request.Showtime,
		SeatNumber: "A1",
		Status:     "Confirmed",
	}, nil
}

func (m *MockMovieTicketService) ViewTicketService(email string) ([]models.Ticket, error) {
	return []models.Ticket{}, nil
}

func (m *MockMovieTicketService) ViewAttendeesService(movieTitle, showtime string) ([]models.Attendees, error) {
	return []models.Attendees{}, nil
}

func (m *MockMovieTicketService) CancelTicketService(request models.CancelTicketRequest) error {
	return nil
}

func (m *MockMovieTicketService) ModifySeatService(request models.ModifySeatRequest) error {
	return nil
}
