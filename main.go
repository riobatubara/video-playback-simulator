package main

import (
	"flag"
	"fmt"
	"runtime"
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

	// FIX: Explicitly enforce safety warnings
	if *apiURL != "" && *apiKey == "" {
		fmt.Println("[ERROR] --api_url is set but --api_key is missing; payloads will NOT be sent.")
	}

	// Set API config globally
	utils.SetAPIConfig(*apiURL, *apiKey)

	var wg sync.WaitGroup
	start := time.Now()

	totalSessions := *concurrent

	// IMPROVEMENT: Limit active thread exhaustion using a channel-based semaphore token bucket.
	// If concurrent is set to a massive number, it safely limits active execution pools.
	maxParallelWorkers := totalSessions
	if maxParallelWorkers > (runtime.NumCPU() * 100) {
		maxParallelWorkers = runtime.NumCPU() * 100 // Safe upper-limit to protect OS network sockets
	}

	sem := make(chan struct{}, maxParallelWorkers)

	for i := 0; i < totalSessions; i++ {
		wg.Add(1)

		// Block here if active worker ceiling has been breached
		sem <- struct{}{}

		go func(index int) {
			defer func() {
				<-sem // Release worker slot back to bucket
				wg.Done()
			}()

			session := simulator.NewSession(index)
			session.Run()
		}(i)
	}

	wg.Wait()
	fmt.Printf("[DONE] All sessions completed in %s.\n", time.Since(start))
}
