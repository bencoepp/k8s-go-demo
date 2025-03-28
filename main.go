package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	message string
	config  string
	logFile *os.File
)

func main() {
	// Setup logging
	var err error
	logFile, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	multiLog := log.New(io.MultiWriter(os.Stdout, logFile), "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Read environment variable
	message = os.Getenv("MESSAGE")
	if message == "" {
		message = "Default message"
	}

	// Read config file
	configBytes, err := ioutil.ReadFile("/etc/demo/config")
	if err != nil {
		multiLog.Printf("Error reading config: %v", err)
		config = "Could not read config"
	} else {
		config = string(configBytes)
	}

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		multiLog.Println("Accessed /env endpoint")
		fmt.Fprintf(w, "Message: %s", message)
	})

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		multiLog.Println("Accessed /config endpoint")
		fmt.Fprintf(w, "Config: %s", config)
	})

	multiLog.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
