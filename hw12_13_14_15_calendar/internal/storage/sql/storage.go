package sqlstorage

import (
	"context"
	"database/sql"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
	_ "github.com/jackc/pgx/stdlib"
)

type Storage struct {
	db *sql.DB
}

func New(dataSourceName string) (*Storage, error) {
	database, err := sql.Open("pgx", dataSourceName)
	return &Storage{db: database}, err
}

func (s *Storage) Connect(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Storage) Add(ctx context.Context, ev storage.Event) error {
	query := `
	insert into events
		(id, title, event_date, expiration_date, description, user_id, seconds_until_notification)
	values
		($1, $2, $3, $4, $5, %6, %7)
	returning id`
	_, err := s.db.Exec(query, ev.ID, ev.Title, ev.EventDate, ev.Description, ev.UserID, ev.SecondsUntilNotification)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Update(ctx context.Context, ev storage.Event) error {
	query := `
	update events 
	set 
		title = $1,
		event_date = $2,
		expiration_date = $3,
		description = $4,
		user_id = $5,
		seconds_until_notification = $6
	where
		id = $1`
	_, err := s.db.Exec(query, ev.ID, ev.Title, ev.EventDate, ev.Description, ev.UserID, ev.SecondsUntilNotification)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(ctx context.Context, ev storage.Event) error {
	query := `delete events where id = $1`
	_, err := s.db.Exec(query, ev.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) ListDayEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	y, m, d := startTime.Date()
	startDate := time.Date(y, m, d, 0, 0, 0, 0, startTime.Location())
	finishDate := startDate.AddDate(0, 0, 1).Add(-1 * time.Second)
	return s.ListEvents(ctx, startDate, finishDate)
}

func (s *Storage) ListWeekEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	y, m, d := startTime.Date()
	startDate := time.Date(y, m, d, 0, 0, 0, 0, startTime.Location())
	finishDate := startDate.AddDate(0, 0, 7).Add(-1 * time.Second)
	return s.ListEvents(ctx, startDate, finishDate)
}

func (s *Storage) ListMonthEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	y, m, d := startTime.Date()
	startDate := time.Date(y, m, d, 0, 0, 0, 0, startTime.Location())
	finishDate := startDate.AddDate(0, 1, 0).Add(-1 * time.Second)
	return s.ListEvents(ctx, startDate, finishDate)
}

func (s *Storage) ListEvents(ctx context.Context, startTime time.Time, finishTime time.Time) ([]storage.Event, error) {
	query := `
	select 
		id, title, event_date, expiration_date, description, user_id, seconds_until_notification 
	from 
		events 
	where 
		event_date >= $1 and event_date <= $2`
	rows, err := s.db.QueryContext(ctx, query, startTime, finishTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]storage.Event, 0)
	for rows.Next() {
		ev := storage.Event{}
		err = rows.Scan(&ev.ID, &ev.Title, &ev.EventDate, &ev.ExpirationDate, &ev.Description, &ev.UserID, &ev.SecondsUntilNotification)
		if err != nil {
			return nil, err
		}
		list = append(list, ev)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}
