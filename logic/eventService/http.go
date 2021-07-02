package eventService

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *EventService ) {
	subRouter := router.PathPrefix("/api/event/game").Subrouter()
	subRouter.HandleFunc("/new", service.onInvitedForGame()).Methods("POST")
	subRouter.HandleFunc("/{gameId}/salvo", service.onSalvoFired()).Methods("PUT")
}

func (s *EventService) onInvitedForGame()http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *EventService) onSalvoFired()http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
