package main

import (
	"encoding/json"
	"net/http"
)

// INFO: Че так много undefind?
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

func TokenHand(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Message: "invalid json"})
		return
	}

	if user.ID == "" {
		writeJson(w, http.StatusBadRequest, ErrorResponse{Message: "user_id is required"})
		return
	}

	writeJson(w, http.StatusOK, Token{Access_token: "sdfsawerwarjls"})

}
