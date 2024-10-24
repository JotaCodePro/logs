package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

// Values for colors
const (
	Blue       = "\033[34m"
	Yellow     = "\033[33m"
	Red        = "\033[31m"
	Reset      = "\033[0m"
	DateFormat = "2006-01-02 15:04:05"
)

// Models
type Handler struct {
	time    time.Time
	level   string
	message string
	format  int
}

type message struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"msg"`
}

// NewLogger ...
func NewLogger(format int) *Handler {
	return &Handler{format: format}
}

// Internal Functions
// logWithColor ...
func (l *Handler) logWithColor() {

	var color string

	switch l.level {
	case "INFO":
		color = Blue
	case "WARNING":
		color = Yellow
	case "ERROR":
		color = Red
	default:
		color = Reset
	}

	messages := message{
		Level:   l.level,
		Time:    l.time.Format(DateFormat),
		Message: l.message,
	}

	// 1 = JSON
	// 2 = Color
	// 3 = Default
	var dataMessage string
	switch l.format {
	case 1:
		a, _ := json.Marshal(messages)
		dataMessage = string(a)
	case 2:
		dataMessage = fmt.Sprintf("[%s] Time: %s - Message: %s", l.level, l.time.Format(DateFormat), l.message)
	default:
		color = Reset
		fmt.Println(l.message)
	}
	fmt.Println(color + dataMessage)
}

// Public Functions
// Info ...
func (l *Handler) Info(data string) {
	l.level = "INFO"
	l.time = time.Now()
	l.message = data
	l.logWithColor()
}

// Error ...
func (l *Handler) Error(data string) {
	l.level = "ERROR"
	l.time = time.Now()
	l.message = data
	l.logWithColor()
}

// Warning ...
func (l *Handler) Warning(data string) {
	l.level = "WARNING"
	l.time = time.Now()
	l.message = data
	l.logWithColor()
}
