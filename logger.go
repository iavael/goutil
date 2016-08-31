package goutil

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Loglevel int

const (
	FATAL Loglevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)

type Logger struct {
	loglev Loglevel
}

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

func NewLogger(l Loglevel) *Logger {
	return &Logger{
		loglev: l,
	}
}

func (self *Logger) Fatal(fmt string, s ...interface{}) {
	if FATAL <= self.loglev {
		log.Fatalf(getPrefix("FATAL", fmt), s...)
	}
}

func (self *Logger) Error(fmt string, s ...interface{}) {
	if ERROR <= self.loglev {
		log.Printf(getPrefix("ERROR", fmt), s...)
	}
}

func (self *Logger) Warn(fmt string, s ...interface{}) {
	if WARN <= self.loglev {
		log.Printf(getPrefix("WARN", fmt), s...)
	}
}

func (self *Logger) Info(fmt string, s ...interface{}) {
	if INFO <= self.loglev {
		log.Printf(getPrefix("INFO", fmt), s...)
	}
}

func (self *Logger) Debug(fmt string, s ...interface{}) {
	if DEBUG <= self.loglev {
		log.Printf(getPrefix("DEBUG", fmt), s...)
	}
}
