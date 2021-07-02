package eventService

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *UserService) {
	subRouter := router.PathPrefix("/api/event").Subrouter()
	subRouter.HandleFunc("", service.onEvent()).Methods("POST")
}

func (s *UserService) onEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO unpack request

		// TODO Convert into event

		// TODO Push event into service
	}
}
