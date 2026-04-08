package booking

import "context"

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

func (s *Service) ConfirmSeat(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return s.repo.Confirm(ctx, sessionID, userID)
}

func (s *Service) ReleaseSeat(ctx context.Context, sessionID string, userID string) error {
	return s.repo.Release(ctx, sessionID, userID)
}
