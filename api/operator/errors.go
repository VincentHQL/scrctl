package operator

import (
	"net/http"

	apiv1 "github.com/VincentHQL/scrctl/api/operator/apiv1"
)

type AppError struct {
	Msg        string
	StatusCode int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) JSONResponse() apiv1.ErrorMsg {
	return apiv1.ErrorMsg{Error: e.Msg, Details: e.Error()}
}

func NewBadRequestError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusBadRequest, Err: e}
}

func NewInternalError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusInternalServerError, Err: e}
}

func NewNotFoundError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusNotFound, Err: e}
}

func NewServiceUnavailableError(msg string, e error) error {
	return &AppError{Msg: msg, StatusCode: http.StatusServiceUnavailable, Err: e}
}
