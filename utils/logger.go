package utils

import (
	"log"
)

// LogEvent prints a formatted terminal log for each emitted event
func LogEvent(index int, tsclient int64, sessionID, event, value string) {
	log.Printf("data[%d]:: tsclient: %d, sessid: %s, event: %s, value: %s",
		index, tsclient, sessionID, event, value)
}
