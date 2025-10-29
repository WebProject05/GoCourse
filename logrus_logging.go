package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	// Set the log level
	log.SetLevel(logrus.InfoLevel)

	// Set the log format (JSON format)
	log.SetFormatter(&logrus.JSONFormatter{})

	// logging examples
	log.Info("This is an info message")
	log.Warn("This is an warning message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "Santosh",
		"method": "GET",
	}).Info("User logged in")
}
