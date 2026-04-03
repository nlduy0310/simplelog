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

func (l *Logger) Debug(format string, args ...any) {
	l.logAt(DEBUG, format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.logAt(INFO, format, args...)
}

func (l *Logger) Warning(format string, args ...any) {
	l.logAt(WARNING, format, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.logAt(ERROR, format, args...)
}

// Fatal logs at level FATAL then calls os.Exit(1)
func (l *Logger) Fatal(format string, args ...any) {
	l.logAt(FATAL, format, args...)
	os.Exit(1)
}

func (l *Logger) logAt(level Level, format string, args ...any) {
	if l.level.Allow(level) {
		l.logger.Printf(
			"%s - [%s] [%s]: %s\n",
			time.Now().Format(timestampFormat),
			l.Name,
			level.Name(),
			fmt.Sprintf(format, args...),
		)
	}
}
