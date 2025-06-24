package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	apiURL  string
	apiKey  string
	onceSet sync.Once
)

// SetAPIConfig allows main.go to set both api_url and api_key
func SetAPIConfig(url, key string) {
	onceSet.Do(func() {
		apiURL = url
		apiKey = key
	})
}

// EmitPayload sends to API if configured; otherwise, does nothing
func EmitPayload(event, sessid, value string) {
	payload := EventPayload{
		TsClient: time.Now().UnixMilli(),
		SessID:   sessid,
		Event:    event,
		Value:    value,
	}

	data, err := json.Marshal([]EventPayload{payload})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal payload: %v\n", err)
		return
	}

	if apiURL == "" {
		return // No API, just simulate
	}
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "Cannot send to API: missing X-API-Key\n")
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send payload: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		fmt.Fprintf(os.Stderr, "API responded with non-success status: %s\n", resp.Status)
	}
}
