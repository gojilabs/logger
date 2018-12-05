// +build !debug,!info,!notice

package logger

func NoticeErr(_ error) {
}

func Notice(_ ...interface{}) {
}
