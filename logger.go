package logger

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gojilabs/environment"
)

type levelByte byte

const debug = levelByte(0x00)
const info = levelByte(0x01)
const warn = levelByte(0x02)
const err = levelByte(0x03)

type Level rune

const DEBUG = Level('D')
const INFO = Level('I')
const WARN = Level('W')
const ERROR = Level('E')

const space = ' '
const newline = '\n'
const equals = '='

var prefix bytes.Buffer
var buf bytes.Buffer
var prefixMap map[string]string

var level = debug
var msgLevel = DEBUG
var timestampEnabled = false

var mutex sync.Mutex
var writer io.Writer

func Initialize(logLevel Level, writers ...io.Writer) {
	level = debug
	msgLevel = DEBUG

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

	if len(writers) > 0 {
		writer = io.MultiWriter(writers...)
	} else {
		writer = os.Stdout
	}

	prefixMap = make(map[string]string)

	timestampEnabled = environment.IsDevelopment() || environment.IsTest()
}

func SetTimestampEnabled(enabled bool) {
	timestampEnabled = enabled
}

func AddPrefix(key string, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	prefixMap[key] = value

	prefix.Reset()

	for k, v := range prefixMap {
		prefix.WriteString(k)
		prefix.WriteRune(equals)
		prefix.WriteString(v)
		prefix.WriteRune(space)
	}
}

func shouldWrite(myLevel levelByte) bool {
	return myLevel >= level
}

func writeLine(level levelByte, levelRune Level, msg string) {
	if shouldWrite(level) {
		defer buf.Reset()

		if timestampEnabled {
			buf.WriteString(time.Now().Format(time.StampMicro))
			buf.WriteRune(space)
		}

		mutex.Lock()
		defer mutex.Unlock()

		buf.WriteRune(rune(levelRune))
		buf.WriteRune(space)
		buf.Write(prefix.Bytes())
		buf.WriteString(msg)
		buf.WriteRune(newline)
		buf.WriteTo(writer)
	}
}

func Debug(msg string) {
	writeLine(debug, DEBUG, msg)
}

func Info(msg string) {
	writeLine(info, INFO, msg)
}

func Warn(msg string) {
	writeLine(warn, WARN, msg)
}

func Error(msg string) {
	writeLine(err, ERROR, msg)
}

func DebugErr(err error) {
	Debug(err.Error())
}

func InfoErr(err error) {
	Info(err.Error())
}

func WarnErr(err error) {
	Warn(err.Error())
}

func ErrorErr(err error) {
	Error(err.Error())
}
