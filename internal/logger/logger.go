package logger

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	debug *log.Logger
}

func New(level string) *Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile

	return &Logger{
		info:  log.New(os.Stdout, "INFO  ", flags),
		warn:  log.New(os.Stdout, "WARN  ", flags),
		error: log.New(os.Stderr, "ERROR ", flags),
		debug: log.New(os.Stdout, "DEBUG ", flags),
	}
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Warn(msg string) {
	l.warn.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.error.Println(msg)
}

func (l *Logger) Debug(msg string) {
	l.debug.Println(msg)
}
