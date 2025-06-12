package debug_log

import (
	"log"
	"os"
)

func LogToFile(message string) {
	// Open the log file in append mode
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return
	}
	defer f.Close()

	// Create a new logger
	logger := log.New(f, "[LOG] ", log.Ldate|log.Ltime|log.Lshortfile)

	// Log the message
	logger.Println(message)
}
