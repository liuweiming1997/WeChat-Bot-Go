package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Error(s ...interface{}) {
	logrus.Error(s...)
}

func Fatal(s ...interface{}) {
	logrus.Fatal(s...)
}

func Warning(s ...interface{}) {
	logrus.Warning(s...)
}

func Info(s ...interface{}) {
	logrus.Info(s...)
}

func Debug(s ...interface{}) {
	fmt.Println(s...)
}
