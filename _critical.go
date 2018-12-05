// +build !debug,!info,!notice,!warn,!error,!critical

package logger

func CriticalErr(_ error) {
}

func Critical(_ ...interface{}) {
}
