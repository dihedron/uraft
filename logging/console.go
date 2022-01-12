package logging

import (
	"fmt"
	"io"
	"strings"
)

func NewConsoleLogger(stream io.Writer) *ConsoleLogger {
	return &ConsoleLogger{
		stream: stream,
	}
}

type ConsoleLogger struct {
	stream io.Writer
}

func (l *ConsoleLogger) Trace(msg string, args ...interface{}) {
	message := fmt.Sprintf("[TRACE] "+msg, args...)
	message = strings.TrimRight(message, "\n\r")
	fmt.Fprintln(l.stream, message)
}

func (l *ConsoleLogger) Debug(msg string, args ...interface{}) {
	message := fmt.Sprintf("[DEBUG] "+msg, args...)
	message = strings.TrimRight(message, "\n\r")
	fmt.Fprintln(l.stream, message)
}

func (l *ConsoleLogger) Info(msg string, args ...interface{}) {
	message := fmt.Sprintf("[INFO] "+msg, args...)
	message = strings.TrimRight(message, "\n\r")
	fmt.Fprintln(l.stream, message)
}

func (l *ConsoleLogger) Warn(msg string, args ...interface{}) {
	message := fmt.Sprintf("[WARN] "+msg, args...)
	message = strings.TrimRight(message, "\n\r")
	fmt.Fprintln(l.stream, message)
}

func (l *ConsoleLogger) Error(msg string, args ...interface{}) {
	message := fmt.Sprintf("[ERROR] "+msg, args...)
	message = strings.TrimRight(message, "\n\r")
	fmt.Fprintln(l.stream, message)
}
