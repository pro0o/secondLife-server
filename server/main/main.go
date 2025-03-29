package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	storagev1connect "secondLife/gen/bowie/v1/storagev1connect"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"secondLife/internal/data"
)

const (
	defaultPort = "8080"
)

type application struct {
	repo data.Repository
	log  zerolog.Logger
}

func main() {
	os.Exit(run())
}

func run() int {
	ctx := context.Background()

	logger := log.With().Str("service", "secondLife").Logger()
	err := godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	addr := "0.0.0.0:" + port

	dsn := os.Getenv("DSN")
	log.Printf("dsn is:", dsn)
	if dsn == "" {
		logger.Error().Msg("DSN environment variable is required")
		return 1
	}

	// jwtSecret := os.Getenv("JWT_SECRET")
	// if jwtSecret == "" {
	// 	logger.Error().Msg("JWT_SECRET environment variable is required")
	// 	return 1
	// }

	db, err := setupDatabase(dsn)
	if err != nil {
		logger.Error().Err(err).Msg("failed to setup database")
		return 1
	}
	defer db.Close()

	app := &application{
		repo: data.NewRepository(ctx, db),
		log:  logger,
	}

	router := app.setupRouter()

	logger.Info().Str("addr", addr).Msg("starting server")
	err = http.ListenAndServe(
		addr,
		h2c.NewHandler(router, &http2.Server{}),
	)

	if err != nil {
		logger.Error().Err(err).Msg("server error")
		return 1
	}

	return 0
}

func setupDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return db, nil
}
func (app *application) setupRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Mount(storagev1connect.NewStorageServiceHandler(app, connect.WithInterceptors(
		ServiceVersionInterceptor("StorageService", "v1"),
	)))
	return r
}
