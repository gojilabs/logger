// +build debug info notice

package logger

func NoticeErr(err error) {
	Notice(err.Error())
}

func Notice(fields ...interface{}) {
	writeLine(notice, NOTICE, fields...)
}
