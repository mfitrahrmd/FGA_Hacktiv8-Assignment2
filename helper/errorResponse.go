package helper

import (
	"net/http"
)

type MyError struct {
	Code int
	Msg  string
	Data any
}

type ServerErrorResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message" example:"server error, please try again later"`
}

var (
	INVALID_REQUEST_BODY = MyError{http.StatusBadRequest, "Invalid request body", nil}
	NOT_FOUND            = MyError{http.StatusNotFound, "Resources do not exist", nil}
	SERVER_ERROR         = MyError{http.StatusInternalServerError, "Unexpected server error", nil}
)

func (e MyError) Error() string {
	return e.Msg
}
