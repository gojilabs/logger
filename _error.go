// +build !debug,!info,!notice,!warn,!error

package logger

func ErrorErr(_ error) {
}

func Error(_ ...interface{}) {
}
