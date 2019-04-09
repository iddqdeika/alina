package logger

import (
	"alina/alina"
	"fmt"
)

var DefaultLogger alina.Logger

func InitDefaultLogger() {
	DefaultLogger = &logger{}
}

type logger struct {
}

func (l *logger) Info(args ...interface{}) {
	fmt.Println(args)
}

func (l *logger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (l *logger) Error(args ...interface{}) {
	fmt.Println(args)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}
