package goutil

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
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
	syslog *syslog.Writer
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
func NewLogger(l Loglevel, s *syslog.Writer) *Logger {
	return &Logger{
		loglev: l,
		syslog: s,
	}
}

// Fatal wraps log.Fatalf with loglevel check
func (l *Logger) Fatal(f string, s ...interface{}) {
	if FATAL <= l.loglev {
		if l.syslog != nil {
			l.syslog.Crit(fmt.Sprintf(f, s...))
			os.Exit(1)
		} else {
			log.Fatalf(getPrefix("FATAL", f), s...)
		}
	}
}

// Error wraps log.Printf with loglevel check
func (l *Logger) Error(f string, s ...interface{}) {
	if ERROR <= l.loglev {
		if l.syslog != nil {
			l.syslog.Err(fmt.Sprintf(f, s...))
		} else {
			log.Printf(getPrefix("ERROR", f), s...)
		}
	}
}

// Warn wraps log.Printf with loglevel check
func (l *Logger) Warn(f string, s ...interface{}) {
	if WARN <= l.loglev {
		if l.syslog != nil {
			l.syslog.Warning(fmt.Sprintf(f, s...))
		} else {
			log.Printf(getPrefix("WARN", f), s...)
		}
	}
}

// Info wraps log.Printf with loglevel check
func (l *Logger) Info(f string, s ...interface{}) {
	if INFO <= l.loglev {
		if l.syslog != nil {
			l.syslog.Info(fmt.Sprintf(f, s...))
		} else {
			log.Printf(getPrefix("INFO", f), s...)
		}
	}
}

// Debug wraps log.Printf with loglevel check
func (l *Logger) Debug(f string, s ...interface{}) {
	if DEBUG <= l.loglev {
		if l.syslog != nil {
			l.syslog.Debug(fmt.Sprintf(f, s...))
		} else {
			log.Printf(getPrefix("DEBUG", f), s...)
		}
	}
}
