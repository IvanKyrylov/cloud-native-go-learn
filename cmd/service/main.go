package main

import "log/slog"

func main() {
	slog.Info("Start service")

	run()

	defer slog.Info("End service")
}

func run() {
	// ...
}
