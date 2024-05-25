package internalhttp

import (
	"net/http"
)

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
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`UpdateEvent`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("UpdateEvent")
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`DeleteEvent`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("DeleteEvent")
}

func (s *Server) ListDayEvents(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`ListDayEvents`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("ListDayEvents")
}

func (s *Server) ListWeekEvents(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`ListWeekEvents`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("ListWeekEvents")
}

func (s *Server) ListMonthEvents(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`ListMonthEvents`)); err != nil {
		s.logger.Error(err.Error())
	}
	s.logger.Info("ListMonthEvents")
}
