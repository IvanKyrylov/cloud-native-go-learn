package basic

import "log/slog"

func Run(logger *slog.Logger) {
	runBasic(logger)
	runGeneric(logger)
}
