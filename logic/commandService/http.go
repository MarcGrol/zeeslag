package commandService

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *CommandService) {
	subRouter := router.PathPrefix("/api/command").Subrouter()
	subRouter.HandleFunc("/new", service.invite()).Methods("POST")
	subRouter.HandleFunc("/{gameId}/salvo", service.fireSalvo()).Methods("PUT")
	subRouter.HandleFunc("/{gameId}", service.getGame()).Methods("GET")
}

func (s *CommandService) invite() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into command

		// TODO Push command into service

	}
}

func (s *CommandService) fireSalvo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into command

		// TODO Push command into service

	}
}

func (s *CommandService) getGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into query

		// TODO Call service and feed result back

	}
}
