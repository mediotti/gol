package main

import (
	"time"

	"github.com/mediotti/gol/pkg/logger"
)

func main() {

	logger := logger.Logger{
		TimeFormat: time.RFC3339,
	}

	logger.Debug("TESTE - A")
	logger.Info("TESTE - B")
	logger.Warn("TESTE - C")
	logger.Error("TESTE - D")
}
