package handlers

import "log"

type MathLogger struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) *MathLogger {
	return &MathLogger{l}
}
