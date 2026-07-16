package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type EventPayload struct {
	TsClient int64  `json:"tsclient"`
	SessID   string `json:"sessid"`
	Event    string `json:"event"`
	Value    string `json:"value"`
}

var (
	apiURL     string
	apiKey     string
	onceSet    sync.Once
	httpClient *http.Client
)

// SetAPIConfig allows main.go to set both api_url and api_key
func SetAPIConfig(url, key string) {
	onceSet.Do(func() {
		apiURL = url
		apiKey = key
		// FIX: Use a dedicated client with a defined timeout to avoid hanging forever
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	})
}

// EmitPayload sends to API if configured; otherwise, does nothing
func EmitPayload(event, sessid, value string) {
	// IMPROVEMENT: Fast-fail immediately before allocating memory or running json.Marshal
	if apiURL == "" {
		return // No API, just simulate
	}
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "Cannot send to API: missing X-API-Key\n")
		return
	}

	payload := EventPayload{
		TsClient: time.Now().UnixMilli(),
		SessID:   sessid,
		Event:    event,
		Value:    value,
	}

	// Marshaling directly as a single item payload array
	data, err := json.Marshal([]EventPayload{payload})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal payload: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// FIX: Use the customized client instead of DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send payload: %v\n", err)
		return
	}

	// FIX: Clean up response fully to prevent socket leaks and allow connection reuse
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Fprintf(os.Stderr, "API responded with non-success status: %s\n", resp.Status)
	}
}
