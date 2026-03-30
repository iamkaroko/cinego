package booking

type Service struct {
	repo BookingRepository
}

func NewService(repo BookingRepository) *Service {
	return &Service{repo}
}

func (s *Service) Book(booking Booking) error {
	return s.repo.Book(&booking)
}
