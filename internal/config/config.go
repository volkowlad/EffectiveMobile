package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Config struct {
	HTTPServerConfig HTTPServer
	DBConfig         DB
}

type HTTPServer struct {
	Address string
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func InitConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file", err.Error())
	}

	return &Config{
		HTTPServerConfig: HTTPServer{
			Address: os.Getenv("HTTP_PORT"),
		},

		DBConfig: DB{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}
}

func (c *Config) Validate() error {
	if c.HTTPServerConfig.Address == "" {
		slog.Error("server: missing port")
		return errors.New("server: missing port")
	}

	if c.DBConfig.User == "" {
		slog.Error("DB: missing user")
		return errors.New("DB: missing user")
	}

	if c.DBConfig.Password == "" {
		slog.Error("DB: missing password")
		return errors.New("DB: missing password")
	}

	if c.DBConfig.Host == "" {
		slog.Error("DB: missing host")
		return errors.New("DB: missing host")
	}

	if c.DBConfig.Port == "" {
		slog.Error("DB: missing port")
		return errors.New("DB: missing port")
	}

	if c.DBConfig.DBName == "" {
		slog.Error("DB: missing name")
		return errors.New("DB: missing name")
	}

	if c.DBConfig.SSLMode == "" {
		slog.Error("DB: missing mode")
		return errors.New("DB: missing mode")
	}

	return nil
}
