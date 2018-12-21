package logger

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gojilabs/environment"
)

type Level uint8

const debug = Level(7)
const info = Level(6)
const notice = Level(5)
const warn = Level(4)
const err = Level(3)
const critical = Level(2)
const alert = Level(1)
const emergency = Level(0)

type Severity rune

const DEBUG = Severity('7')
const INFO = Severity('6')
const NOTICE = Severity('5')
const WARN = Severity('4')
const ERROR = Severity('3')
const CRITICAL = Severity('2')
const ALERT = Severity('1')
const EMERGENCY = Severity('0')

const space = ' '
const newline = '\n'
const equals = '='
const openBracket = '<'
const closeBracket = '>'

var prefix strings.Builder
var buf strings.Builder
var prefixMap map[string]string

var timestampEnabled = false

var mutex sync.Mutex
var writer io.Writer

func Initialize(writers ...io.Writer) {
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
	if key == "" || value == "" {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	prefixMap[key] = value

	prefix.Reset()

	var keys []string
	for k := range prefixMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			prefix.WriteRune(space)
		}

		prefix.WriteString(k)
		prefix.WriteRune(equals)
		prefix.WriteString(prefixMap[k])
	}
}

func writeLine(l Level, severity Severity, fields ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	buf.WriteRune(openBracket)
	buf.WriteRune(rune(severity))
	buf.WriteRune(closeBracket)

	if timestampEnabled {
		buf.WriteString(time.Now().Format(time.StampMicro))
		buf.WriteRune(space)
	}

	buf.WriteString(prefix.String())

	var key string
	var shouldInclude bool

	for i, field := range fields {
		if i%2 == 0 {
			key, shouldInclude = field.(string)
		} else if shouldInclude {
			value := fmt.Sprint(field)
			if value != "" {
				buf.WriteRune(space)
				buf.WriteString(key)
				buf.WriteRune(equals)
				buf.WriteString(value)
			}
		}
	}

	buf.WriteRune(newline)
	writer.Write([]byte(buf.String()))
	buf.Reset()
}
