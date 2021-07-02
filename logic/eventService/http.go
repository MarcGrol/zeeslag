package eventService

import (
	"encoding/json"
	"github.com/MarcGrol/zeeslag/core"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *UserService) {
	subRouter := router.PathPrefix("/api/event").Subrouter()
	subRouter.HandleFunc("", service.onEvent()).Methods("POST")
}

func (s *UserService) onEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event := core.GameEventPdu{}
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = s.OnEvent(event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
