package operator

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	apiv1 "github.com/VincentHQL/scrctl/api/operator/apiv1"
)

// Send a JSON http response with success status code to the client
func ReplyJSONOK(w http.ResponseWriter, obj interface{}) error {
	return ReplyJSON(w, obj, http.StatusOK)
}

// Send a JSON http response with error to the client
func ReplyJSONErr(w http.ResponseWriter, err error) error {
	log.Printf("response with error: %v\n", err)
	var e *AppError
	if errors.As(err, &e) {
		return ReplyJSON(w, e.JSONResponse(), e.StatusCode)
	}
	return ReplyJSON(w, apiv1.ErrorMsg{Error: "Internal Server Error"}, http.StatusInternalServerError)
}

// Send a JSON http response to the client
func ReplyJSON(w http.ResponseWriter, obj interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	return encoder.Encode(obj)
}
