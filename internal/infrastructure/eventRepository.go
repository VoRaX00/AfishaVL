package infrastructure

import (
	"Afisha/internal/domain"
	"github.com/jmoiron/sqlx"
)

type EventRepository struct {
	db sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: *db,
	}
}

func (repo *EventRepository) Create(event domain.Event) (int, error) {

}

func (repo *EventRepository) Update(event domain.Event) error {
	return nil
}

func (repo *EventRepository) GetById(eventId string) (domain.Event, error) {
	return domain.Event{}, nil
}

func (repo *EventRepository) Delete(eventId string) error {
	return nil
}
