package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code   string `json:"code"`
	Detail string `json:"detail,omitempty"`
}

func Success(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var appErr *apperrors.AppError
	if !errors.As(err, &appErr) {

		appErr = apperrors.NewInternalError("Erro inesperado", err)
	}

	w.WriteHeader(appErr.StatusCode)

	json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: appErr.Message,
		Error: &ErrorInfo{
			Code:   appErr.Code,
			Detail: appErr.Detail,
		},
	})

	log.Printf("Erro tratado: %v", appErr.Err)
}

func Created(w http.ResponseWriter, message string, data interface{}) {
	Success(w, http.StatusCreated, message, data)
}

func OK(w http.ResponseWriter, message string, data interface{}) {
	Success(w, http.StatusOK, message, data)
}
