package main

import (
	"log/slog"
	"os"
	"url_shortner/internal/config"
	"url_shortner/internal/lib/logger/sl"
	"url_shortner/internal/storage/sqlite"
)

func main() {
	config.EnvLoad()

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("debug logging enabled")

	storage, err := sqlite.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Error("failed to create storage", sl.Err(err))
		os.Exit(1)
	}

	//id, err := storage.SaveURL("https://google.com", "google")
	//if err != nil {
	//	log.Error("failed to save url", sl.Err(err))
	//	os.Exit(1)
	//}

	//log.Info("saved url", slog.Int64("id", id))

	//err = storage.DeleteURL("google")
	//if err != nil {
	//	log.Error("failed to delete url", sl.Err(err))
	//	os.Exit(1)
	//}

	//alias, err := storage.GetURL("google")
	//if err != nil {
	//	log.Error("failed to get url alias", sl.Err(err))
	//  os.Exit(1)
	//}
	//
	//log.Info("getted alias from 'google'", slog.String("alias", alias))

	_ = storage

	// TODO: init router: chi + net/http, render
	// TODO: run server
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
