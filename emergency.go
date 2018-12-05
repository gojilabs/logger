package logger

func EmergencyErr(err error) {
	Emergency(err.Error())
}

func Emergency(fields ...interface{}) {
	writeLine(emergency, EMERGENCY, fields...)
}
