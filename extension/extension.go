package extension

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// EventType represents the type of events recieved from /event/next
type EventType string
type ExtensionHeader string

const (
	// Invoke is a lambda invoke
	Invoke EventType = "INVOKE"
	// Shutdown is a shutdown event for the environment
	Shutdown EventType = "SHUTDOWN"

	ExtensionNameHeader ExtensionHeader = "Lambda-Extension-Name"
	ExtensionIdentifierHeader ExtensionHeader = "Lambda-Extension-Identifier"
)

// RegisterResponse is the body of the response for /register
type RegisterResponse struct {
	FunctionName    string `json:"functionName"`
	FunctionVersion string `json:"functionVersion"`
	Handler         string `json:"handler"`
}

type ExtensionService struct {
	baseURL     string
	httpClient  *http.Client
	extensionID string
	subscribedEventTypes []EventType
}

// NewExtensionService creates a new HTTP client with custom settings
func NewExtensionService(httpClient *http.Client, subscribedEventTypes []EventType, awsLambdaRuntimeAPI string) *ExtensionService {
	baseURL := fmt.Sprintf("http://%s/2020-01-01/extension", awsLambdaRuntimeAPI)

	return &ExtensionService{
		baseURL: baseURL,
		httpClient: httpClient,
		subscribedEventTypes: subscribedEventTypes,
	}
}

func (e *ExtensionService) Register(ctx context.Context, extensionName string) (*RegisterResponse, error) {
    // Construct the full URL for the registration request
    url := e.baseURL + "/register"

    // Create and marshal the request body to JSON
    reqBody, err := json.Marshal(map[string]interface{}{
        "events": e.subscribedEventTypes,
    })

    if err != nil {
        return nil, err
    }

    // Build the HTTP POST request with the context, URL, and request body
    httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBody))
    if err != nil {
        return nil, err
    }

    httpReq.Header.Set(string(ExtensionNameHeader), extensionName) // Set extension name header

    // Perform the HTTP request
    httpRes, err := e.httpClient.Do(httpReq)
    if err != nil || httpRes.StatusCode != 200 {
        return nil, fmt.Errorf("request failed with status %s", httpRes.Status)
    }
    defer httpRes.Body.Close() // Close response body

    // Parse the response body into a RegisterResponse struct
    var res RegisterResponse
    if err = json.NewDecoder(httpRes.Body).Decode(&res); err != nil {
        return nil, err
    }

    // Store the extension ID from the response headers
    e.extensionID = httpRes.Header.Get(string(ExtensionIdentifierHeader))
    fmt.Println("Extension id:", e.extensionID) // Log the extension ID

    return &res, nil // Return the parsed response
}