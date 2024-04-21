package main

import (
	"net/http"
)

var ErrUserInvalid = apiError{Err: "User Not Valid", Status: http.StatusNotFound}

type apiFunc func(http.ResponseWriter, *http.Request) error

func (e apiError) Error() string {
	return e.Err
}
