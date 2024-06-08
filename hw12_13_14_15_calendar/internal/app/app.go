package app

import (
	"context"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	log     Logger
	Storage Storage
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
	return &App{log: logger, Storage: storage}
}

func (a *App) CreateEvent(ctx context.Context, ev storage.Event) error {
	return a.Storage.Add(ctx, ev)
}

func (a *App) UpdateEvent(ctx context.Context, ev storage.Event) error {
	return a.Storage.Update(ctx, ev)
}

func (a *App) DeleteEvent(ctx context.Context, ev storage.Event) error {
	return a.Storage.Delete(ctx, ev)
}

func (a *App) ListDayEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	return a.Storage.ListDayEvents(ctx, startTime)
}

func (a *App) ListWeekEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	return a.Storage.ListWeekEvents(ctx, startTime)
}

func (a *App) ListMonthEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error) {
	return a.Storage.ListMonthEvents(ctx, startTime)
}
