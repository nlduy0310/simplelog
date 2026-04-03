package simplelog

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	logger *log.Logger
	Name   string
	level  Level
}

func NewLogger(name string, level Level) Logger {
	assertValidLevel(level)
	return Logger{
		logger: log.New(os.Stderr, "", 0),
		Name:   name,
		level:  level,
	}
}

func (l *Logger) Level() Level {
	return l.level
}

func (l *Logger) SetLevel(level Level) {
	assertValidLevel(level)
	l.level = level
}

func (l *Logger) allow(level Level) bool {
	return l.level.Allow(level)
}

// format-style log

func (l *Logger) Debugf(format string, args ...any) {
	l.logfIfNeeded(DEBUG, format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.logfIfNeeded(INFO, format, args...)
}

func (l *Logger) Warningf(format string, args ...any) {
	l.logfIfNeeded(WARNING, format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.logfIfNeeded(ERROR, format, args...)
}

// Fatalf logs at level FATAL then calls os.Exit(1)
func (l *Logger) Fatalf(format string, args ...any) {
	l.logfIfNeeded(FATAL, format, args...)
	os.Exit(1)
}

// log literal strings

func (l *Logger) Debug(msg string) {
	l.logIfNeeded(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.logIfNeeded(INFO, msg)
}

func (l *Logger) Warning(msg string) {
	l.logIfNeeded(WARNING, msg)
}

func (l *Logger) Error(msg string) {
	l.logIfNeeded(ERROR, msg)
}

// Fatal logs the message then calls os.Exit(1)
func (l *Logger) Fatal(msg string) {
	l.logIfNeeded(FATAL, msg)
	os.Exit(1)
}

func (l *Logger) logfIfNeeded(level Level, format string, args ...any) {
	if l.allow(level) {
		l.logAt(level, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) logIfNeeded(level Level, msg string) {
	if l.allow(level) {
		l.logAt(level, msg)
	}
}

func (l *Logger) logAt(level Level, msg string) {
	l.logger.Printf(
		"%s - [%s] [%s]: %s\n",
		time.Now().Format(timestampFormat),
		l.Name,
		level.Name(),
		msg,
	)
}
