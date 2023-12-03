package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
)

// Logger represents the instance of the Logrus logger used for logging in the application.
var Logger = logrus.New()

// InitLog initializes the application logger by configuring Logrus settings.
// It sets the logger's formatter to JSONFormatter and the log level to InfoLevel.
func InitLog() {
	Logger.Infof("Initializing application logger.")

	// Set Logrus to log the file name and line no
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			fileName := filepath.Base(f.File)
			return "", fmt.Sprintf(" %s:%d", fileName, f.Line)
		},
	})
	Logger.SetLevel(logrus.DebugLevel)
	Logger.Debugf("Application logger initialized successfully.")
}
