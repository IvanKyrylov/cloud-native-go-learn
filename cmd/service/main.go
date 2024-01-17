package main

import (
	"log/slog"
	"os"

	"github.com/IvanKyrylov/cloud-native-go-learn/basic"
	customLogger "github.com/IvanKyrylov/cloud-native-go-learn/pkg/logger"
	"github.com/IvanKyrylov/cloud-native-go-learn/pkg/runner"
)

func main() {
	logger := customLogger.NewSlog(os.Getenv("LOG_LVL"))

	logger.Info("Start app")
	defer logger.Info("End app")

	basicService := basic.New(logger.With(slog.Group("service", slog.String("name", "basic"))))

	run(basicService)
}

func run(services ...runner.Runner) {
	for _, service := range services {
		service.Run()
	}
}
