package logging

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var once sync.Once
var logger *zap.Logger

func createLogger(isDevMode bool) {
	once.Do(func() {
		var err error
		if isDevMode {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}
		if err != nil {
			logger = zap.NewNop()
		}
		fmt.Println("logger is created")
	})
}

func GetLogger(isDevMode bool) *zap.Logger {
	if logger == nil {
		createLogger(isDevMode)
	}
	return logger
}
