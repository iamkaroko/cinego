package booking

import "errors"

var (
	ErrSeatAlreadyBooked = errors.New("seat already booked")
	ErrSeatNotFound      = errors.New("seat not found")
)

// Booking represents a reservation made by a user for a specific movie and seat.
type Booking struct {
	ID      string
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

// BookingRepository defines methods to manage and retrieve bookings in the system.
type BookingRepository interface {
	Book(booking *Booking) error
	ListByMovieID(movieID string) ([]Booking, error)
}
