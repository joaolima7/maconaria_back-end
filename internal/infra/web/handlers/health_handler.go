package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
	Database  string `json:"database"`
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbStatus := "ok"
	if err := h.db.Ping(); err != nil {
		dbStatus = "error"
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	response := HealthResponse{
		Status:    "ok",
		Service:   "maconaria-api",
		Timestamp: time.Now().Format(time.RFC3339),
		Database:  dbStatus,
	}

	json.NewEncoder(w).Encode(response)
}
