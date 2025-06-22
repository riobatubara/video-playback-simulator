package main

import (
	"os"
	"strconv"
	"sync"

	"video-playback-simulator/simulator"
	// "video-playback-simulator/utils"
)

func main() {
	// Load environment configuration
	concurrentStr := os.Getenv("CONCURRENT")
	if concurrentStr == "" {
		concurrentStr = "1"
	}
	concurrent, err := strconv.Atoi(concurrentStr)
	if err != nil || concurrent < 1 {
		concurrent = 1
	}

	var wg sync.WaitGroup
	wg.Add(concurrent)

	// Run multiple sessions concurrently
	for i := 0; i < concurrent; i++ {
		go func(index int) {
			sess := simulator.NewSession(index)
			sess.Run()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
