package main

import (
	ef "EffectiveMobile"
	"EffectiveMobile/internal/config"
	"EffectiveMobile/internal/handler"
	"EffectiveMobile/internal/repository"
	"EffectiveMobile/internal/service"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

// @title Effective Mobile
// @version 1.0.0
// @description test task

// @host localhost:8080
// @BasePath /
func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, opts))

	cfg := config.InitConfig()
	err := cfg.Validate()
	if err != nil {
		log.Error("failed to validate config", err.Error())
	}
	log.Info("init config success")

	dbConfig := repository.NewConfigDB(cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.DBName, cfg.DBConfig.SSLMode)
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		log.Error("failed to connect to database", err.Error())
	}
	log.Info("connect to database")

	repos := repository.NewRepos(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(ef.Server)
	go func() {
		if err := srv.Run(cfg.HTTPServerConfig.Address, handlers.InitRouter()); err != nil {
			log.Error("error occured while running http server", err.Error())
		}
	}()

	log.Info("http server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Server is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("server shutdown failed", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Error("failed to close database", err.Error())
	}
}
