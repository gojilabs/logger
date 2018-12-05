// +build debug info notice warn error critical alert

package logger

func Alert(fields ...interface{}) {
	writeLine(alert, ALERT, fields...)
}

func AlertErr(err error) {
	Alert(err.Error())
}
