package environment

import (
	"errors"
	"os"
)

var env string

const Production = "production"
const Development = "development"
const Demo = "demo"
const Staging = "staging"
const Test = "test"

func SetFromEnv(environmentVariable string) error {
	if env != "" {
		return errors.New("environment: cannot change environment after setting it")
	}

	if environmentVariable == "" {
		return errors.New("environment: no environment variable specified")
	}

	return SetFromString(os.Getenv(environmentVariable))
}

func SetFromString(value string) error {
	if value == "" {
		return errors.New("environment: environment variable is empty")
	}
	env = value

	if IsUnknown() {
		return errors.New("environment: unknown environment type " + env)
	}
	os.Stdout.WriteString("Initializing " + env + " environment...\n")

	return nil
}

func String() string {
	return env
}

func SetDevelopment() error {
	return SetFromString(Development)
}

func IsDevelopment() bool {
	return env == Development
}

func SetStaging() error {
	return SetFromString(Staging)
}

func IsStaging() bool {
	return env == Staging
}

func SetDemo() error {
	return SetFromString(Demo)
}

func IsDemo() bool {
	return env == Demo
}

func SetProduction() error {
	return SetFromString(Production)
}

func IsProduction() bool {
	return env == Production
}

func SetTest() error {
	return SetFromString(Test)
}

func IsTest() bool {
	return env == Test
}

func IsKnown() bool {
	return IsProduction() || IsDemo() || IsStaging() || IsTest() || IsDevelopment()
}

func IsLocal() bool {
	return IsDevelopment() || IsTest()
}

func IsRemote() bool {
	return !IsLocal()
}

func IsUnknown() bool {
	return !IsKnown()
}
