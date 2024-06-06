package internalgrpc

import (
	context "context"
	"fmt"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateEvent(ctx context.Context, e *AddEventRequest) (*emptypb.Empty, error) {
	event := storage.Event{
		ID:                       uuid.NewString(),
		Title:                    e.Event.Title,
		EventDate:                e.Event.EventDate.AsTime(),
		ExpirationDate:           e.Event.ExpirationDate.AsTime(),
		Description:              e.Event.Description,
		UserID:                   e.Event.UserID,
		SecondsUntilNotification: e.Event.SecondsUntilNotification,
	}

	err := s.app.Storage.Add(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("grpc create event error: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateEvent(ctx context.Context, e *UpdateEventRequest) (*emptypb.Empty, error) {
	event := storage.Event{
		ID:                       e.Event.ID,
		Title:                    e.Event.Title,
		EventDate:                e.Event.EventDate.AsTime(),
		ExpirationDate:           e.Event.ExpirationDate.AsTime(),
		Description:              e.Event.Description,
		UserID:                   e.Event.UserID,
		SecondsUntilNotification: e.Event.SecondsUntilNotification,
	}

	err := s.app.Storage.Update(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("grpc update event error: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteEvent(ctx context.Context, e *RemoveEventRequest) (*emptypb.Empty, error) {
	event := storage.Event{
		ID: e.Event.ID,
	}

	err := s.app.Storage.Delete(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("grpc delete event error: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func eventsToGrpcEvents(events []storage.Event) []*GrpcEvent {
	protoEvents := make([]*GrpcEvent, len(events))
	for i, event := range events {
		protoEvents[i] = &GrpcEvent{
			ID:                       event.ID,
			Title:                    event.Title,
			EventDate:                timestamppb.New(event.EventDate),
			ExpirationDate:           timestamppb.New(event.ExpirationDate),
			Description:              event.Description,
			UserID:                   event.UserID,
			SecondsUntilNotification: event.SecondsUntilNotification,
		}
	}

	return protoEvents
}

func (s *Server) ListDayEvents(ctx context.Context, e *GetEventsRequest) (*GetEventsResponse, error) {
	events, err := s.app.Storage.ListDayEvents(ctx, e.StartDate.AsTime())

	response := &GetEventsResponse{
		Events: eventsToGrpcEvents(events),
	}

	return response, err
}

func (s *Server) ListWeekEvents(ctx context.Context, e *GetEventsRequest) (*GetEventsResponse, error) {
	events, err := s.app.Storage.ListWeekEvents(ctx, e.StartDate.AsTime())

	response := &GetEventsResponse{
		Events: eventsToGrpcEvents(events),
	}

	return response, err
}

func (s *Server) ListMonthEvents(ctx context.Context, e *GetEventsRequest) (*GetEventsResponse, error) {
	events, err := s.app.Storage.ListMonthEvents(ctx, e.StartDate.AsTime())

	response := &GetEventsResponse{
		Events: eventsToGrpcEvents(events),
	}

	return response, err
}

func (s *Server) mustEmbedUnimplementedEventsServer() {
}
