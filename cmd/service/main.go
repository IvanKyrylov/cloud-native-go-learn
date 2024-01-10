package main

import (
	"log/slog"
	"os"

	"github.com/IvanKyrylov/cloud-native-go-learn/basic"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	logger.Info("Start service")
	defer logger.Info("End service")

	run(logger)
}

func run(logger *slog.Logger) {
	basic.Run(logger)
}
