package handlers

import "log"

type MathLogger struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) *MathLogger {
	return &MathLogger{l}
}

type Triangle struct {
	Type string  `json:"type,omitempty"`
	A    float64 `json:"a,omitempty"`
	B    float64 `json:"b,omitempty"`
	C    float64 `json:"c,omitempty"`
}

type Four struct {
	Type string  `json:"type,omitempty"`
	A    float64 `json:"a,omitempty"`
}

type Rectangle struct {
	Type string  `json:"type,omitempty"`
	A    float64 `json:"a,omitempty"`
	B    float64 `json:"b,omitempty"`
}

type Circle struct {
	Type string  `json:"type,omitempty"`
	R    float64 `json:"r,omitempty"`
}

type InputQE struct {
	A float64 `json:"a,omitempty"`
	B float64 `json:"b,omitempty"`
	C float64 `json:"c,omitempty"`
}
