package http_test

import (
    "testing"
    "time"

    "github.com/clockstop/splunkeddy/http"
)

func TestCreateHTTPClient(t *testing.T) {
    // Call the function to get the HTTP client
    client := http.CreateHTTPClient()

    // Check if the returned client is not nil
    if client == nil {
        t.Fatalf("expected non-nil client, got nil") //Fatal means to stop the test and don't continue forward.
    }

    // Check if the Timeout is correctly set
    expectedTimeout := 10 * time.Second
    if client.Timeout != expectedTimeout {
        t.Errorf("expected timeout %v, got %v", expectedTimeout, client.Timeout)
    }

    // Optionally, you could also check other fields, if any, in the http.Client.
}
