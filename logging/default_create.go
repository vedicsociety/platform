// The NewDefaultLogger function creates a DefaultLogger with a minimum severity level and log.
// Loggers that write messages to standard out.
// For a test, change the main function so that it writes out its message using the logging feature.
package logging

import (
	"log"
	"os"
	"strings"

	"github.com/vedicsociety/platform/config"
)

func NewDefaultLogger(cfg config.Configuration) Logger {
	// Obtain the logging level from the configuration system
	var level LogLevel = Debug
	if configLevelString, found := cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}

	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}

// There is no good way to represent iota values in JSON, so we used a string
// and defined the LogLevelFromString function to convert the configuration setting to a LogLevel value
func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":
		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}
