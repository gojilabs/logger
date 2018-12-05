// +build !debug,!info,!notice,!warn,!error,!critical,!alert

package logger

func Alert(_ ...interface{}) {
}

func AlertErr(_ error) {
}
