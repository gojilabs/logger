package logger

import (
  "log"
  "strings"
  "time"
)

const debug = 0
const info = 1
const warn = 2
const err = 3

const DEBUG = 'D'
const INFO = 'I'
const WARN = 'W'
const ERROR = 'E'

type Logger struct {
  level int
  msgLevel rune
}

func New(logLevel string) *Logger {

  level := debug
  msgLevel := DEBUG

  if logLevel != "" {
    levelStr := strings.ToLower(logLevel)
    if levelStr == "error" {
      level = err
      msgLevel = ERROR
    } else if levelStr == "warn" {
      level = warn
      msgLevel = WARN
    } else if levelStr == "info" {
      level = info
      msgLevel = INFO
    }
  }

  return &Logger{level: level, msgLevel: msgLevel}
}

func (l *Logger) shouldWrite(level int) bool {
  return level >= l.level
}

func (l *Logger) writeLine(level int, msg string) {
  if l.shouldWrite(level) {
    log.Println(string(l.msgLevel) + " " + time.Now().Format(time.StampMicro) + " " + msg)
  }
}

func (l *Logger) Debug(msg string) {
  l.writeLine(debug, msg)
}

func (l *Logger) Info(msg string) {
  l.writeLine(info, msg)
}

func (l *Logger) Warn(msg string) {
  l.writeLine(warn, msg)
}

func (l *Logger) Error(msg string) {
  l.writeLine(err, msg)
}