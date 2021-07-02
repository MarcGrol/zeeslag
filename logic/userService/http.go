package userService

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *UserService) {
	subRouter := router.PathPrefix("/api/user").Subrouter()
	subRouter.HandleFunc("/new", service.inviteForGame()).Methods("POST")
	subRouter.HandleFunc("/{gameId}", service.getGame()).Methods("GET")
	subRouter.HandleFunc("/{gameId}/salvo", service.fireSalvo()).Methods("POST")
}

func (s *UserService) inviteForGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into command

		// TODO Push command into service

	}
}

func (s *UserService) fireSalvo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into command

		// TODO Push command into service

	}
}

func (s *UserService) getGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into query

		// TODO Call service and feed result back

	}
}
