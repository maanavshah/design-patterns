package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Simulates resource-intensive logger creation
func New() *logrus.Logger {
	time.Sleep(1 * time.Second) // resource initialization -> network connections, file connections, etc
	return logrus.New()
}

// Function to create a new logger instance every time
func GetNewLogger() *logrus.Logger {
	return New()
}

// Singleton pattern ensures that only one instance of the logger is created
var (
	loggerOnce sync.Once
	logger     *logrus.Logger
)

func GetSingletonLogger() *logrus.Logger {
	loggerOnce.Do(func() {
		logger = New()
	})
	return logger
}

func main() {
	fmt.Println("Starting application...")

	// Example 1: New logger instance every time
	fmt.Println("Using a new logger instance each time:")
	for i := 0; i < 5; i++ {
		log := GetNewLogger()
		log.Infof("Log message %d from new logger", i+1)
	}

	// Example 2: Singleton logger
	fmt.Println("Using a singleton logger instance:")
	for i := 0; i < 5; i++ {
		log := GetSingletonLogger()
		log.Infof("Log message %d from singleton logger", i+1)
	}
}
