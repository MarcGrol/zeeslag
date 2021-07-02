package peerService

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *PeerService) {
	subRouter := router.PathPrefix("/api/event/game").Subrouter()
	subRouter.HandleFunc("/{gameId}", service.onInvitedForGame()).Methods("POST")
	subRouter.HandleFunc("/{gameId}", service.onAcceptedByPeer()).Methods("PUT")
	subRouter.HandleFunc("/{gameId}", service.onRejectedByPeer()).Methods("DELETE")
	subRouter.HandleFunc("/{gameId}/salvo/{salvoId}", service.onSalvoFiredByPeer()).Methods("POST")
	subRouter.HandleFunc("/{gameId}/salvo/{salvoId}", service.onSalvoImpactAssessedByPeer()).Methods("PUT")
}

func (s *PeerService) onInvitedForGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}

func (s *PeerService) onAcceptedByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}

func (s *PeerService) onRejectedByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}
func (s *PeerService) onSalvoFiredByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}

func (s *PeerService) onSalvoImpactAssessedByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}
