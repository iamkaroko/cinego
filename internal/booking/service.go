package booking

type Service struct {
	repo BookingRepository
}

func NewService(repo BookingRepository) *Service {
	return &Service{repo}
}

func (s *Service) Book(b Booking) (Booking, error) {
	return s.repo.Book(b)
}

func (s *Service) ListBookings(movieID string) []Booking {
	return s.repo.ListByMovieID(movieID)
}
