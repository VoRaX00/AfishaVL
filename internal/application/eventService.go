package application

import "Afisha/internal/domain"

type EventService struct {
	repo IEventsRepository
}

func NewEventService(repo IEventsRepository) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (s *EventService) Create(event domain.Event) (int, error) {
	return s.repo.Create(event)
}

func (s *EventService) Update(event domain.Event) error {
	return s.repo.Update(event)
}

func (s *EventService) GetById(eventId string) (domain.Event, error) {
	return s.repo.GetById(eventId)
}

func (s *EventService) Delete(eventId string) error {
	return s.repo.Delete(eventId)
}
