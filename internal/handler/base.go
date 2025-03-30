package handler

import (
	"github.com/rs/zerolog"

	"github.com/pro0o/second-life/internal/data"
)

type Handler struct {
	Repo data.Repository
	Log  zerolog.Logger
}

func NewHandler(repo data.Repository, logger zerolog.Logger) *Handler {
	return &Handler{
		Repo: repo,
		Log:  logger,
	}
}
