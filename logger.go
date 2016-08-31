package goutil

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Loglevel typedef for integer
type Loglevel int

// Enum with logging levels
const (
	FATAL Loglevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)

// Logger struct
type Logger struct {
	loglev Loglevel
}

// ParseLoglevel function converts string representation of log levels to int
func ParseLoglevel(level string) Loglevel {
	switch {
	case strings.EqualFold(level, "debug") || level == strconv.Itoa(int(DEBUG)):
		return DEBUG
	case strings.EqualFold(level, "info") || level == strconv.Itoa(int(INFO)):
		return INFO
	case strings.EqualFold(level, "warn") || level == strconv.Itoa(int(WARN)):
		return WARN
	case strings.EqualFold(level, "error") || level == strconv.Itoa(int(ERROR)):
		return ERROR
	case strings.EqualFold(level, "fatal") || level == strconv.Itoa(int(FATAL)):
		return FATAL
	}
	return DEBUG
}

func getPrefix(tag string, f string) string {
	return fmt.Sprintf("[%s] %s", tag, f)
}

// NewLogger creates new Logger object
func NewLogger(l Loglevel) *Logger {
	return &Logger{
		loglev: l,
	}
}

// Fatal wraps log.Fatalf with loglevel check
func (l *Logger) Fatal(fmt string, s ...interface{}) {
	if FATAL <= l.loglev {
		log.Fatalf(getPrefix("FATAL", fmt), s...)
	}
}

// Error wraps log.Printf with loglevel check
func (l *Logger) Error(fmt string, s ...interface{}) {
	if ERROR <= l.loglev {
		log.Printf(getPrefix("ERROR", fmt), s...)
	}
}

// Warn wraps log.Printf with loglevel check
func (l *Logger) Warn(fmt string, s ...interface{}) {
	if WARN <= l.loglev {
		log.Printf(getPrefix("WARN", fmt), s...)
	}
}

// Info wraps log.Printf with loglevel check
func (l *Logger) Info(fmt string, s ...interface{}) {
	if INFO <= l.loglev {
		log.Printf(getPrefix("INFO", fmt), s...)
	}
}

// Debug wraps log.Printf with loglevel check
func (l *Logger) Debug(fmt string, s ...interface{}) {
	if DEBUG <= l.loglev {
		log.Printf(getPrefix("DEBUG", fmt), s...)
	}
}
