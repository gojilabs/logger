// +build debug info notice warn

package logger

func Warn(fields ...interface{}) {
	writeLine(warn, WARN, fields...)
}

func WarnErr(err error) {
	Warn(err.Error())
}
