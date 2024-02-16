package main

import (
	"github.com/pazifical/colorful-logging/clog"
)

func main() {
	logger := clog.NewLogger(clog.Info)

	logger.Debug("This is a debug log")
	logger.Info("This is an info log")
	logger.Warning("This is a warning log")
	logger.Error("This is an error log")
}
