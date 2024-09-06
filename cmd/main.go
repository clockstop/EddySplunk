package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Event types
const (
	InitEvent     = "INIT"
	InvokeEvent   = "INVOKE"
	ShutdownEvent = "SHUTDOWN"
)

// Log to file
func logToFile(eventType, details string) {
	filename := "/tmp/extension_logs.txt"
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("%s: %s - %s\n", timestamp, eventType, details)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// WriteString returns 2 args, first is number of bytes (ignored), second is error.
	// If err != nil, then it logs the fatal error.
	if _, err := f.WriteString(logMessage); err != nil {
		log.Fatal(err)
	}
}

// Main function
func main() {
	fmt.Println("Extension started, attempting registration.")
	// extensionId := registerExtension()
	// fmt.Println("Extension registration successful. extensionId=" + extensionId)
	// for {
	// 	event := nextEvent(extensionID)

	// 	switch event.Type {
	// 	case InitEvent:
	// 		logToFile(InitEvent, "Initialization event")
	// 	case InvokeEvent:
	// 		logToFile(InvokeEvent, "Invocation event")
	// 	case ShutdownEvent:
	// 		logToFile(ShutdownEvent, "Shutdown event")
	// 		return
	// 	}
	// }
}

// Register the extension with the Lambda environment
func registerExtension() string {
	var (
		extensionApiHost string = os.Getenv("AWS_LAMBDA_RUNTIME_API")
		extensionUrl     string = fmt.Sprintf("http://%s/2020-01-01/extension", extensionApiHost)
		registerPayload  string = `{"events":["INVOKE"]}` //TODO Add Shutdown phase as well.
	)
	reader := strings.NewReader(registerPayload)

	fmt.Println("Lambda Runtime API: " + extensionApiHost)
	fmt.Println("Registration Payload: \n" + registerPayload)

	// Create a new POST request
	req, err := http.NewRequest("POST", extensionUrl+"/register", reader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "nil"
	}

	fmt.Println("Request: ", req)

	// Set the Lambda-Extension-Name header
	req.Header.Set("Lambda-Extension-Name", "splunkeddy")

	// Send the request using http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		log.Fatal(err)
	}

	resBody, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	// Handle the response
	fmt.Println("Response status:", resp.Status)
	fmt.Printf("Response body: %s\n", resBody)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the header value for "Lambda-Extension-Identifier"
	var extensionId string = resp.Header.Get("Lambda-Extension-Identifier")

	if extensionId == "" {
		log.Fatalf("Lambda-Extension-Identifier header not found")
	}

	// Use the extensionID as needed
	fmt.Println("Extension ID:", extensionId)

	return extensionId
}

// Get the next event from the Lambda runtime
func nextEvent(extensionID string) (event struct{ Type string }) {
	request, err := http.NewRequest("GET", "http://sandbox:/2020-01-01/extension/event/next", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Lambda-Extension-Identifier", extensionID)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &event)

	return event
}
