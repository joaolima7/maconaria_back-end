package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"seccess"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func Error(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	json.NewEncoder(w).Encode(response)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	Success(w, http.StatusCreated, message, data)
}

func OK(w http.ResponseWriter, message string, data interface{}) {
	Success(w, http.StatusOK, message, data)
}

func BadRequest(w http.ResponseWriter, message string, err error) {
	Error(w, http.StatusBadRequest, message, err)
}

func NotFound(w http.ResponseWriter, message string, err error) {
	Error(w, http.StatusNotFound, message, err)
}

func InternalServerError(w http.ResponseWriter, message string, err error) {
	Error(w, http.StatusInternalServerError, message, err)
}

func Unauthorized(w http.ResponseWriter, message string, err error) {
	Error(w, http.StatusUnauthorized, message, err)
}

func Forbidden(w http.ResponseWriter, message string, err error) {
	Error(w, http.StatusForbidden, message, err)
}
