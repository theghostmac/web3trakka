package housekeeper

import (
	"log"
	"os"
)

// LogLevel constants.
const (
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

// CustomLogger struct holds log related settings.
type CustomLogger struct {
	logger *log.Logger
}

// NewCustomLogger creates a new instance of CustomLogger.
func NewCustomLogger() *CustomLogger {
	//return &CustomLogger{logger: log.New(logger.Writer(), "", logger.Flags())}
	return &CustomLogger{logger: log.New(os.Stdout, "", log.LstdFlags)}
}

// Info logs info messages.
func (c *CustomLogger) Info(message string) {
	c.logger.Printf("%s: %s\n", INFO, message)
}

// Warning logs warning messages.
func (c *CustomLogger) Warning(message string) {
	c.logger.Printf("%s: %s\n", WARNING, message)
}

// Error logs error messages.
func (c *CustomLogger) Error(message string) {
	c.logger.Printf("%s: %s\n", ERROR, message)
}

// Fatal logs fatal messages.
func (c *CustomLogger) Fatal(message string) {
	c.logger.Printf("%s: %s\n", FATAL, message)
	os.Exit(1)
}
