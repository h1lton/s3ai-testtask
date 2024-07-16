package logger

import (
	"fmt"
	"log/slog"
	"os"
	"s3ai-testtask/internal/infrastructure/logger/handlers/slogpretty"
)

// Set sets the default logger based on the given environment.
//
// Parameters:
// - env: the environment to set the logger for. It can be "local", "docker", or "prod".
//
// Returns:
// - error: an error if the environment is unknown.
func Set(env string) error {
	switch env {
	case "local":
		slog.SetDefault(slogpretty.New())
		//slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		//	Level: slog.LevelDebug,
		//})))

	case "docker":
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})))

	case "prod":
		slog.SetDefault(slog.New(slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		)))

	default:
		return fmt.Errorf("logger.Set: unknown env: %s", env)
	}

	return nil
}
