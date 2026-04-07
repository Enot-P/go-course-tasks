package main

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func UserID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Message: "user_id is required"})
		return
	}

	writeJson(w, http.StatusOK, User{ID: id}) // INFO: Нужно передавать струтуру или мапу?

}
