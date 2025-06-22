package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type EventPayload struct {
	TsClient int64  `json:"tsclient"`
	SessID   string `json:"sessid"`
	Event    string `json:"event"`
	Value    string `json:"value"`
}

// EmitPayload creates a payload and (optionally) sends it to an API.
// For now, we'll just marshal and print the JSON.
func EmitPayload(event, sessid, value string) {
	payload := EventPayload{
		TsClient: time.Now().UnixMilli(),
		SessID:   sessid,
		Event:    event,
		Value:    value,
	}

	data, err := json.Marshal([]EventPayload{payload})
	if err != nil {
		fmt.Printf("Failed to marshal payload: %v\n", err)
		return
	}

	// Placeholder for sending to a REST API
	// Example: sendPayloadToAPI(data)

	// For now, do nothing here — logging is handled by logger.go
	_ = data
}
