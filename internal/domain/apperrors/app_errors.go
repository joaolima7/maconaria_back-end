package apperrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"detail,omitempty"`
	StatusCode int    `json:"-"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Detail)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

const (
	CodeValidationError    = "VALIDATION_ERROR"
	CodeDuplicateEntry     = "DUPLICATE_ENTRY"
	CodeNotFound           = "NOT_FOUND"
	CodeUnauthorized       = "UNAUTHORIZED"
	CodeForbidden          = "FORBIDDEN"
	CodeInternalError      = "INTERNAL_ERROR"
	CodeDatabaseError      = "DATABASE_ERROR"
	CodeInvalidCredentials = "INVALID_CREDENTIALS"
)

func NewValidationError(field string, message string) *AppError {
	return &AppError{
		Code:       CodeValidationError,
		Message:    fmt.Sprintf("Erro de validação: %s", field),
		Detail:     message,
		StatusCode: 400,
	}
}

func NewDuplicateError(field string, value string) *AppError {
	return &AppError{
		Code:       CodeDuplicateEntry,
		Message:    fmt.Sprintf("O %s '%s' já está em uso!", field, value),
		Detail:     "Tente usar outro valor!",
		StatusCode: 409,
	}
}

func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Code:       CodeNotFound,
		Message:    fmt.Sprintf("%s não encontrado!", resource),
		StatusCode: 404,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:       CodeUnauthorized,
		Message:    message,
		StatusCode: 401,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:       CodeForbidden,
		Message:    message,
		StatusCode: 403,
	}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Code:       CodeInternalError,
		Message:    "Erro interno do servidor.",
		Detail:     message,
		StatusCode: 500,
		Err:        err,
	}
}

func WrapDatabaseError(err error, operation string) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return parseMySQLError(mysqlErr)
	}

	return &AppError{
		Code:       CodeDatabaseError,
		Message:    fmt.Sprintf("Erro ao %s", operation),
		Detail:     "Erro de banco de dados.",
		StatusCode: 500,
		Err:        err,
	}
}

func parseMySQLError(err *mysql.MySQLError) *AppError {
	switch err.Number {
	case 1062:
		field, value := extractDuplicateInfo(err.Message)
		return NewDuplicateError(field, value)

	case 1452:
		return &AppError{
			Code:       CodeValidationError,
			Message:    "Referência inválida!",
			Detail:     "O registro relacionado não existe!",
			StatusCode: 400,
			Err:        err,
		}

	case 1054:
		return &AppError{
			Code:       CodeDatabaseError,
			Message:    "Erro de banco de dados.",
			Detail:     "Coluna não encontrada.",
			StatusCode: 500,
			Err:        err,
		}

	default:
		return &AppError{
			Code:       CodeDatabaseError,
			Message:    "Erro de banco de dados.",
			Detail:     err.Message,
			StatusCode: 500,
			Err:        err,
		}
	}
}

func extractDuplicateInfo(msg string) (string, string) {
	parts := strings.Split(msg, "'")
	value := ""
	if len(parts) >= 2 {
		value = parts[1]
	}

	field := "campo"
	if strings.Contains(msg, "for key") {
		keyParts := strings.Split(msg, "for key")
		if len(keyParts) >= 2 {
			field = strings.Trim(strings.Split(keyParts[1], "'")[1], "'")
		}
	}

	fieldTranslations := map[string]string{
		"email":    "e-mail",
		"username": "nome de usuário",
		"phone":    "telefone",
		"cpf":      "CPF",
	}

	if translated, ok := fieldTranslations[field]; ok {
		field = translated
	}

	return field, value
}
