package peerService

import (
	"encoding/json"
	"github.com/MarcGrol/zeeslag/core"
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
		request := core.PeerHasInvitedYouForGame{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.callService(w, request.ToPdu())
	}
}

func (s *PeerService) onAcceptedByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unpack request
		request := core.PeerHasAcceptedGame{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.callService(w, request.ToPdu())
	}
}

func (s *PeerService) onRejectedByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unpack request
		request := core.PeerHasRejectedInvitation{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.callService(w, request.ToPdu())
	}
}
func (s *PeerService) onSalvoFiredByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unpack request
		request := core.PeerHasFiredSalvo{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.callService(w, request.ToPdu())
	}
}

func (s *PeerService) onSalvoImpactAssessedByPeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unpack request
		request := core.PeerHasAssessedImpactOfSalvo{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.callService(w, request.ToPdu())
	}
}

func (s *PeerService) callService(w http.ResponseWriter, pdu core.GameMsgPdu) {
	//  Push command into service
	err := s.OnPeerEvent(pdu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
