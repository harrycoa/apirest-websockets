package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/harrycoa/apirest-websockets.git/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Creamos la respuesta
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Api go version 1.0.0",
			Status:  true,
		})

	}
}
