// +build debug info

package logger

func InfoErr(err error) {
	Info(err.Error())
}

func Info(fields ...interface{}) {
	writeLine(info, INFO, fields...)
}
