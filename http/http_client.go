package http

import (
    "net/http"
    "time"
)

// NewHTTPClient creates a new HTTP client with custom settings
func NewHTTPClient() *http.Client {
    return &http.Client{
        Timeout: 10 * time.Second, // Example configuration
    }
}
