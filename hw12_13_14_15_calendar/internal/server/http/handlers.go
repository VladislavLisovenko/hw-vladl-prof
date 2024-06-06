package internalhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage"
)

type eventListOptions struct {
	StartTime time.Time
}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`Hello`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("Hello")
}

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`AddEvent`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("AddEvent")

	var event storage.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			s.logger.Error(err.Error())
		}
		return
	}

	err = s.app.Storage.Add(context.Background(), event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			s.logger.Error(err.Error())
		}
		return
	}
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`UpdateEvent`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("UpdateEvent")

	var event storage.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	err = s.app.Storage.Update(context.Background(), event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			s.logger.Error(err.Error())
		}
		return
	}
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`DeleteEvent`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("DeleteEvent")

	var event storage.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
		return
	}
	err = s.app.Storage.Delete(context.Background(), event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			s.logger.Error(err.Error())
		}
		return
	}
}

func (s *Server) ListDayEvents(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`ListDayEvents`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("ListDayEvents")

	sendList(w, r, "day", s)
}

func (s *Server) ListWeekEvents(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`ListWeekEvents`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("ListWeekEvents")

	sendList(w, r, "week", s)
}

func (s *Server) ListMonthEvents(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`ListMonthEvents`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("ListMonthEvents")

	sendList(w, r, "month", s)
}

func sendList(w http.ResponseWriter, r *http.Request, period string, s *Server) {
	var options eventListOptions
	err := json.NewDecoder(r.Body).Decode(&options)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			s.logger.Error(err.Error())
		}
		return
	}

	var getListFunc func(context.Context, time.Time) ([]storage.Event, error)
	switch period {
	case "week":
		getListFunc = s.app.Storage.ListWeekEvents
	case "month":
		getListFunc = s.app.Storage.ListMonthEvents
	default:
		getListFunc = s.app.Storage.ListDayEvents
	}

	list, err := getListFunc(context.Background(), options.StartTime)

	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		s.logger.Error(err.Error())
		return
	}

	listDecoded, err := json.Marshal(list)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		s.logger.Error(err.Error())
		return
	}

	_, err = w.Write(listDecoded)
	if err != nil {
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		s.logger.Error(err.Error())
		return
	}
}
