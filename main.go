package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"video-playback-simulator/simulator"
	"video-playback-simulator/utils"
)

func main() {
	// CLI flags
	concurrent := flag.Int("concurrent", 1, "Number of concurrent sessions")
	apiURL := flag.String("api_url", "", "Target REST API URL (optional)")
	apiKey := flag.String("api_key", "", "X-API-Key for REST API (optional)")

	flag.Parse()

	if *apiURL != "" {
		if *apiKey == "" {
			fmt.Println("[ERROR] --api_url is set but --api_key is missing; payloads will NOT be sent.")
		}
	}

	// Set API config globally
	utils.SetAPIConfig(*apiURL, *apiKey)

	// Concurrency control
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < *concurrent; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			session := simulator.NewSession(index)
			session.Run()
		}(i)
	}

	wg.Wait()
	fmt.Printf("[DONE] All sessions completed in %s.\n", time.Since(start))
}
