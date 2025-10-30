package Intermediate

// Logging is used to keep track of the development changes made at the specfic time and by the person (in some cases)

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a logging print statment")

	log.SetPrefix("LOG:")
	log.Println("This is an info message")

	// Flags
	// log.SetFlags(log.Ldate)	// This will only log the date
	log.SetFlags(log.Ldate | log.Ltime) // This will only log the date and time
	log.Println("This is a log with the date")

	infoLogger.Println("This is an info logger")
	warnLogger.Println("This is an warning message")
	errorLogger.Println("This is an error message")

	file, err := os.OpenFile("writeString.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 0666 meaning we are granting read and write to every one
	if err != nil {
		log.Fatalf("Failed to open log file.")
	}
	defer file.Close()

	// The debug logger will be printed in the file
	debugLogger := log.New(file, "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message")

}

var (
	// These three logs will be shown in the console as we are using the stdout
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
