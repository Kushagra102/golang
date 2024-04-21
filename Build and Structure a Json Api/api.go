package main

import (
	"encoding/json"
	"net/http"
)


func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e)
				return
			}
			writeJSON(w, http.StatusInternalServerError, apiError{Err: "Internal Server Error", Status: http.StatusInternalServerError})
		}
	}
}


func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "Invalid Requested Method", Status: http.StatusMethodNotAllowed}
	}

	user := User{ID: 1, Valid: true}
	if !user.Valid {
		return ErrUserInvalid
	}

	return writeJSON(w, http.StatusOK, User{})
}
