package internalhttp

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
	logger Logger
}

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Debug(msg string)
}

type Application interface {
	CreateEvent(ctx context.Context, ev storage.Event) error
	UpdateEvent(ctx context.Context, ev storage.Event) error
	DeleteEvent(ctx context.Context, ev storage.Event) error
	ListDayEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error)
	ListWeekEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error)
	ListMonthEvents(ctx context.Context, startTime time.Time) ([]storage.Event, error)
}

func NewServer(logger Logger, app Application, host string, port string) *Server {
	router := chi.NewRouter()
	server := &Server{
		server: &http.Server{
			Handler: loggingMiddleware(router, logger),
			Addr:    net.JoinHostPort(host, port),
		},
		logger: logger,
	}

	router.Get("/", server.Hello)
	router.Get("/hello", server.Hello)
	router.Post("/events", server.AddEvent)
	router.Put("/events/{id:[0-9a-z-]+}", server.UpdateEvent)
	router.Delete("/events/{id:[0-9a-z-]+}", server.DeleteEvent)
	router.Get("/events/day", server.ListDayEvents)
	router.Get("/events/week", server.ListWeekEvents)
	router.Get("/events/month", server.ListMonthEvents)

	return server
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.server.ListenAndServe(); err != nil {
		return s.server.Shutdown(ctx)
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
