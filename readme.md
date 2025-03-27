# Movie Ticket Booking System

### Running the application

To run the application, navigate to the project directory and type:

    go test ./... && go run main.go

The application will start and listen on port 8080.

## Requirements

### 1. Book Movie Ticket API
- Allow a user to book a movie ticket by providing:
  - Name
  - Email
  - Movie title
  - Showtime
- Assign a seat automatically based on availability.
- Return a ticket confirmation with movie details, seat number, and showtime.

### 2. View Movie Ticket Details API
- Retrieve a user's movie ticket details by providing their email.

### 3. View All Attendees for a Movie API
- Return a list of all attendees for a specific movie showtime, along with their seat numbers.

### 4. Cancel Movie Ticket API
- Allow a user to cancel their movie ticket by providing their email and showtime details.

### 5. Modify Seat Assignment API
- Allow a user to modify their seat for a specific movie by providing their email and showtime.


## Overview
The **Movie Ticket Booking System** provides APIs to book movie tickets, view ticket details, manage seating, and cancel reservations.

## API Endpoints

| Endpoint                     | Method | Description |
|------------------------------|--------|-------------|
| `/api/book-ticket`           | POST   | Book a movie ticket, assign a seat, and return confirmation. |
| `/api/view-ticket`           | GET    | Retrieve a user's ticket details using email. |
| `/api/view-attendees`        | GET    | Get a list of attendees for a specific movie showtime. |
| `/api/cancel-ticket`         | DELETE | Cancel a ticket using email and showtime details. |
| `/api/modify-seat`           | PUT    | Modify seat assignment for a specific movie. |

## API Details

### 1. **Book Movie Ticket**
**Endpoint:** `/api/book-ticket`  
**Method:** `POST`  
**Request Body:**  
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "movie_title": "Inception",
  "showtime": "2025-04-01 18:30"
}
```
**Response:**  
```json
{
  "ticket_id": "12345",
  "name": "John Doe",
  "movie_title": "Inception",
  "showtime": "2025-04-01 18:30",
  "seat_number": "A10",
  "status": "Confirmed"
}
```

### 2. **View Movie Ticket Details**
**Endpoint:** `/api/view-ticket?email=john.doe@example.com`  
**Method:** `GET`  
**Response:**  
```json
{
  "ticket_id": "12345",
  "name": "John Doe",
  "movie_title": "Inception",
  "showtime": "2025-04-01 18:30",
  "seat_number": "A10",
  "status": "Confirmed"
}
```

### 3. **View All Attendees for a Movie**
**Endpoint:** `/api/view-attendees?movie_title=Inception&showtime=2025-04-01 18:30`  
**Method:** `GET`  
**Response:**  
```json
{
    "attendees": [
        {
            "name": "John Doe",
            "seat_number": "A49"
        },
        {
            "name": "Jane Smith",
            "seat_number": "A18"
        }
    ]
}
```

### 4. **Cancel Movie Ticket**
**Endpoint:** `/api/cancel-ticket`  
**Method:** `DELETE`  
**Request Body:**  
```json
{
  "email": "john.doe@example.com",
  "showtime": "2025-04-01 18:30"
}
```
**Response:**  
```json
{
  "message": "Ticket successfully canceled."
}
```

### 5. **Modify Seat Assignment**
**Endpoint:** `/api/modify-seat`  
**Method:** `PUT`  
**Request Body:**  
```json
{
  "email": "john.doe@example.com",
  "showtime": "2025-04-01 18:30",
  "new_seat_number": "A12"
}
```
**Response:**  
```json
{
  "message": "Seat updated successfully.",
  "new_seat_number": "A12"
}
```

## Requirements
- GoLang (Gin, Fiber or any preffered framework)
- Database (PostgreSQL, MySQL, etc.)
- Restfull Api
- JSON Request and Response
- Unit Test
- Docker

## Why to use Gin Framework
- Performance
- Simplicity
- Middleware Support
- Readability


## License
MIT License

