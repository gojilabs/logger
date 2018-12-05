// +build debug info notice warn error critical

package logger

func CriticalErr(err error) {
	Critical(err.Error())
}

func Critical(fields ...interface{}) {
	writeLine(critical, CRITICAL, fields...)
}
