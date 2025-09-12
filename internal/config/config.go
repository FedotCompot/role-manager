package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseConnection string `envconfig:"DB_CONNECTION"`

	DiscordBotToken string `envconfig:"BOT_TOKEN"`
}

var Data Config

func init() {
	slog.Info("loading config...")
	if err := godotenv.Load(); err != nil {
		slog.Debug("cannot load .env file")
	}
	if err := envconfig.Process("", &Data); err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}
	slog.Debug("env", slog.Any("config", Data))
	slog.Info("config loaded")
}
