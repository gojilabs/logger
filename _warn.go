// +build !debug,!info,!notice,!warn

package logger

func Warn(_ ...interface{}) {
}

func WarnErr(_ error) {
}
