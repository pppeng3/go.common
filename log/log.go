package log

import (
	"io"
	"sync"
)

type Logger struct {
	Level  Level
	mu     sync.Mutex
	prefix string
	out    io.Writer
	buf    []byte
}

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func New(out io.Writer, prefix string) *Logger {
	return &Logger{}
}

func Default() *Logger {
	return &Logger{}
}

func (l *Logger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {

}

func (l *Logger) SetOutput(w io.Writer) {

}
