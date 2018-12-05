// +build debug

package logger

func Debug(fields ...interface{}) {
	writeLine(debug, DEBUG, fields...)
}

func DebugErr(err error) {
	Debug(err.Error())
}
