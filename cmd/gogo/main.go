package main

import (
	"log/slog"
	"os"

	"github.com/radisvaliullin/gogo/internal/gogo"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("gogo app run")

	gogo := gogo.New(logger)
	if err := gogo.Start(); err != nil {
		logger.Error("gogo service start", "error", err)
		return 1
	}

	return 0
}
