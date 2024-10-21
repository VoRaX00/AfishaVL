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

func (repo *EventRepository) Create(event domain.CreateEvent) (int, error) {
	var id int
	query := "INSERT INTO events (user_id, name, description, city_id, category_id, image, price_id, comment, age_limit, registration, contact_id, dateStart, timeStart, Place, Address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id"
	row := repo.db.QueryRow(query, event.UserId, event.NameEvent, event.DiscriptionEvent, event.CityId, event.CategoryId, event.PosterImg, event.Price, event.Comment, event.AgeLimit, event.Registration, event.ContactId, event.DateStart, event.TimeStart, event.Place, event.Address, &id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *EventRepository) Update(event domain.Event) error {
	query := "UPDATE events SET user_id = $1, name = $2, description = $3, city_id = $4, category_id = $5, image = $6, price_id = $7, age_limit = $8, shedule_id = $9, registration = $10, contact_id = $11 RETURNING id"
	_, err := repo.db.Exec(query, event.UserId, event.Name, event.Description, event.CityId, event.CategoryId, event.Image, event.PriceId, event.AgeLimit, event.ScheduleId, event.Registration, event.ContactId, event.Id)
	return err
}

func (repo *EventRepository) GetById(eventId string) (domain.Event, error) {
	var event domain.Event
	query := "SELECT * FROM events WHERE id = $1"
	row := repo.db.QueryRow(query, eventId)
	if err := row.Scan(&event); err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (repo *EventRepository) Delete(eventId string) error {
	query := "DELETE FROM events WHERE id = $1"

	result, err := repo.db.Exec(query, eventId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}
	return nil
}
