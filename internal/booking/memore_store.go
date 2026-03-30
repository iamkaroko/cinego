package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (m *MemoryStore) Book(booking *Booking) error {
	if _, exists := m.bookings[booking.ID]; exists {
		return ErrSeatAlreadyBooked
	}
	m.bookings[booking.ID] = *booking
	return nil
}
func (m *MemoryStore) ListByMovieID(movieID string) ([]Booking, error) {
	var result []Booking

	for _, booking := range m.bookings {
		if booking.MovieID == movieID {
			result = append(result, booking)
		}
	}
	return result, nil
}
