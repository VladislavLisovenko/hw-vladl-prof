package memorystorage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
)

var (
	ErrorEventAlreadyExists = fmt.Errorf("event already exists")
	ErrorEventDoesntExists  = fmt.Errorf("event doesn't exist")
	ErrorDateBusy           = fmt.Errorf("time is already occupied by another event")
)

type Storage struct {
	data map[string]storage.Event
	mu   sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{
		mu:   sync.RWMutex{},
		data: make(map[string]storage.Event),
	}
}

func (s *Storage) eventByTimeExists(id string, evTime time.Time) bool {
	for _, ev := range s.data {
		if ev.ID != id && ev.EventDate.Equal(evTime) {
			return true
		}
	}

	return false
}

func (s *Storage) Add(ctx context.Context, ev storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.eventByTimeExists(ev.ID, ev.EventDate) {
		return ErrorDateBusy
	}
	if _, ok := s.data[ev.ID]; ok {
		return ErrorEventAlreadyExists
	}
	s.data[ev.ID] = ev

	return nil
}

func (s *Storage) Update(ctx context.Context, ev storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.eventByTimeExists(ev.ID, ev.EventDate) {
		return ErrorDateBusy
	}
	if _, ok := s.data[ev.ID]; !ok {
		return ErrorEventDoesntExists
	}
	s.data[ev.ID] = ev

	return nil
}

func (s *Storage) Delete(ctx context.Context, ev storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[ev.ID]; !ok {
		return ErrorEventDoesntExists
	}
	delete(s.data, ev.ID)

	return nil
}

func (s *Storage) ListDayEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	list := make([]storage.Event, 0)
	for _, ev := range s.data {
		if ev.EventDate.Year() == startTime.Year() && ev.EventDate.YearDay() == startTime.YearDay() {
			list = append(list, ev)
		}
	}

	return list, nil
}

func (s *Storage) ListWeekEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	list := make([]storage.Event, 0)
	startDateYear := startTime.Year()
	startDateDay := startTime.YearDay()
	finishDateYear := startTime.AddDate(0, 0, 7).Year()
	finishDateDay := startTime.AddDate(0, 0, 7).YearDay()

	for _, ev := range s.data {
		if ev.EventDate.Year() >= startDateYear && ev.EventDate.YearDay() >= startDateDay && ev.EventDate.Year() <= finishDateYear && ev.EventDate.Day() <= finishDateDay {
			list = append(list, ev)
		}
	}

	return list, nil
}

func (s *Storage) ListMonthEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	list := make([]storage.Event, 0)
	startDateYear := startTime.Year()
	startDateDay := startTime.YearDay()
	finishDateYear := startTime.AddDate(0, 1, 0).Year()
	finishDateDay := startTime.AddDate(0, 1, 0).YearDay()

	for _, ev := range s.data {
		if ev.EventDate.Year() >= startDateYear && ev.EventDate.YearDay() >= startDateDay && ev.EventDate.Year() <= finishDateYear && ev.EventDate.Day() <= finishDateDay {
			list = append(list, ev)
		}
	}

	return list, nil
}
