package service

import "time"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	currentYear := time.Now().Year()
	nyDate := time.Date(currentYear+1, time.January, 1, 0, 0, 0, 0, time.Now().Location())
	dur := time.Until(nyDate)

	return int64(dur.Hours()) / 24
}
