package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: make(map[string]Booking),
	}
}

func (m *ConcurrentStore) Book(booking *Booking) error {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.bookings[booking.ID]; exists {
		return ErrSeatAlreadyBooked
	}
	m.bookings[booking.ID] = *booking
	return nil
}
func (m *ConcurrentStore) ListByMovieID(movieID string) ([]Booking, error) {
	m.RLock()
	defer m.RUnlock()

	var result []Booking

	for _, booking := range m.bookings {
		if booking.MovieID == movieID {
			result = append(result, booking)
		}
	}
	return result, nil
}
