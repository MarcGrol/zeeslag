package eventService

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *EventService) {
	subRouter := router.PathPrefix("/api/event/game").Subrouter()
	subRouter.HandleFunc("/new", service.onInvitedForGame()).Methods("POST")
	subRouter.HandleFunc("/{gameId}/salvo/{salvoId}", service.onSalvoFired()).Methods("PUT")
	subRouter.HandleFunc("/{gameId}/salvo/{salvoId}impact", service.onSalvoImpactAssessed()).Methods("PUT")
}

func (s *EventService) onInvitedForGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}

func (s *EventService) onSalvoFired() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}

func (s *EventService) onSalvoImpactAssessed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}
