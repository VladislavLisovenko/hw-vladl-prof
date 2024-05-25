package app

import (
	"context"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	log     Logger
	storage Storage
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

type Storage interface {
	Add(ctx context.Context, ev storage.Event) error
	Update(ctx context.Context, ev storage.Event) error
	Delete(ctx context.Context, ev storage.Event) error
	ListDayEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error)
	ListWeekEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error)
	ListMonthEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{log: logger, storage: storage}
}

func (a *App) CreateEvent(ctx context.Context, ev storage.Event) error {
	return a.storage.Add(ctx, ev)
}

func (a *App) UpdateEvent(ctx context.Context, ev storage.Event) error {
	return a.storage.Update(ctx, ev)
}

func (a *App) DeleteEvent(ctx context.Context, ev storage.Event) error {
	return a.storage.Delete(ctx, ev)
}

func (a *App) ListDayEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	return a.storage.ListDayEvents(ctx, startTime)
}

func (a *App) ListWeekEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	return a.storage.ListWeekEvents(ctx, startTime)
}

func (a *App) ListMonthEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	return a.storage.ListMonthEvents(ctx, startTime)
}
