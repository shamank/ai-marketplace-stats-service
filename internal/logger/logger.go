package logger

import (
	"log/slog"
	"os"
)

// InitLogger инициализация логгера
func InitLogger() *slog.Logger {

	// TODO: передавать режим запуска приложения и возвращать соответствующий логгер

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
