package http

import (
    "net/http"
    "time"
)

// CreateHTTPClient creates a new HTTP client with custom settings
func CreateHTTPClient() *http.Client {
    return &http.Client{
        Timeout: 10 * time.Second, // Example configuration
    }
}
