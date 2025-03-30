package handler

import (
	"github.com/rs/zerolog"

	"secondLife/internal/data"
)

// Handler is the base handler structure that implements the StorageService
type Handler struct {
	Repo data.Repository
	Log  zerolog.Logger
}

// NewHandler creates a new handler with dependencies injected
func NewHandler(repo data.Repository, logger zerolog.Logger) *Handler {
	return &Handler{
		Repo: repo,
		Log:  logger,
	}
}
