// The first server feature to implement is logging.
// The log package in the Go standard library provides a good set of basic features for creating logs
// but needs some additional features to filter those messages for detail.
// This file defines the Logger interface, which specifies methods for logging messages with different levels of severity,
// which is set using a LogLevel value ranging from Trace to Fatal.
// There is also a None level that specifies no logging output.
// For each level of severity, the Logger interface defines one method that accepts a simple string
// and one method that accepts a template string and placeholder values.
// I define interfaces for all the features that the platform provides and use those interfaces to provide default implementations.
// This will allow the application to replace the default implementation if required and also make it possible to provide applications
// with features as services, which I describe later in this chapter.
package logging

type LogLevel int

// 0,1,2,3,4,5
const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}
