package logger

import (
	"os"
	"sync"
	"time"

	"github.com/gojilabs/environment"
)

const debug = 0
const info = 1
const warn = 2
const err = 3

type Level rune

const DEBUG Level = 'D'
const INFO Level = 'I'
const WARN Level = 'W'
const ERROR Level = 'E'

type Logger struct {
	level            int
	msgLevel         Level
	mutex            sync.Mutex
	prefix           string
	timestampEnabled bool
}

func New(logLevel Level) *Logger {

	level := debug
	msgLevel := DEBUG

	if logLevel == ERROR {
		level = err
		msgLevel = ERROR
	} else if logLevel == WARN {
		level = warn
		msgLevel = WARN
	} else if logLevel == INFO {
		level = info
		msgLevel = INFO
	}

	return &Logger{level: level, msgLevel: msgLevel, timestampEnabled: environment.IsDevelopment() || environment.IsTest()}
}

func (l *Logger) SetTimestampEnabled(enabled bool) {
	l.timestampEnabled = enabled
}

func (l *Logger) AddPrefix(key string, value string) {
	l.prefix = l.prefix + key + "=" + value + " "
}

func (l *Logger) shouldWrite(level int) bool {
	return level >= l.level
}

func (l *Logger) writeLine(level int, levelRune Level, msg string) {
	if l.shouldWrite(level) {
		timestamp := ""
		if l.timestampEnabled {
			timestamp = time.Now().Format(time.StampMicro) + " "
		}

		l.mutex.Lock()
		defer l.mutex.Unlock()

		os.Stdout.WriteString(timestamp + string(levelRune) + " " + l.prefix + msg + "\n")
	}
}

func (l *Logger) Debug(msg string) {
	l.writeLine(debug, DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.writeLine(info, INFO, msg)
}

func (l *Logger) Warn(msg string) {
	l.writeLine(warn, WARN, msg)
}

func (l *Logger) Error(msg string) {
	l.writeLine(err, ERROR, msg)
}

func (l *Logger) DebugErr(err error) {
	l.Debug(err.Error())
}

func (l *Logger) InfoErr(err error) {
	l.Info(err.Error())
}

func (l *Logger) WarnErr(err error) {
	l.Warn(err.Error())
}

func (l *Logger) ErrorErr(err error) {
	l.Error(err.Error())
}
