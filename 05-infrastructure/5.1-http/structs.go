package main

type User struct {
	ID string `json:"id,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type Token struct {
	Access_token string `json:"access_token,omitempty"`
}
