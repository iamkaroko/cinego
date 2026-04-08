package booking

import (
	"errors"
	"time"
)

var (
	ErrSeatAlreadyBooked = errors.New("seat already booked")
)

// Booking represents a reservation made by a user for a specific movie and seat.
type Booking struct {
	ID        string
	MovieID   string
	SeatID    string
	UserID    string
	Status    string
	ExpiresAt time.Time
}

// BookingRepository defines methods to manage and retrieve bookings in the system.
type BookingRepository interface {
	Book(b Booking) (Booking, error)
	ListByMovieID(movieID string) []Booking
}
