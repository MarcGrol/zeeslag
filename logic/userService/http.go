package userService

import (
	"encoding/json"
	"github.com/MarcGrol/zeeslag/core"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoint(router *mux.Router, service *UserService) {
	subRouter := router.PathPrefix("/api/user/game").Subrouter()
	subRouter.HandleFunc("/new", service.inviteForGame()).Methods("POST")
	subRouter.HandleFunc("/{gameId}", service.getGame()).Methods("GET")
	subRouter.HandleFunc("/{gameId}/salvo", service.fireSalvo()).Methods("POST")
}

type HttpErrorResponse struct {
	ErrorMessage string
}

func (s *UserService) inviteForGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unpack request
		request := core.InviteForGame{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpErrorResponse{
				ErrorMessage: err.Error(),
			})
			return
		}

		s.callService(w, request.ToPdu())
	}
}

func (s *UserService) fireSalvo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unpack request
		request := core.FireSalvo{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpErrorResponse{
				ErrorMessage: err.Error(),
			})
			return
		}

		s.callService(w, request.ToPdu())
	}
}

func (s *UserService) getGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameId := mux.Vars(r)["gameId"]

		game, exists, err := s.OnQuery(gameId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(HttpErrorResponse{
				ErrorMessage: err.Error(),
			})
			return
		}

		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(HttpErrorResponse{
				ErrorMessage: err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(game)
	}
}

func (s *UserService) callService(w http.ResponseWriter, pdu core.GameCommandPdu) {
	//  Push command into service
	err := s.OnCommand(pdu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(HttpErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}
