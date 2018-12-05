// +build debug info notice warn error

package logger

func ErrorErr(err error) {
	Error(err.Error())
}

func Error(fields ...interface{}) {
	writeLine(err, ERROR, fields...)
}
