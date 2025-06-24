package main

import (
	"sync"
	"time"

	"video-playback-simulator/simulator"
)

const (
	numSessions = 1 // You can increase this for load simulation
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numSessions)

	for i := 0; i < numSessions; i++ {
		go func(index int) {
			defer wg.Done()
			session := simulator.NewSession(index)
			session.Run()
		}(i)
		// Optional staggered start for realism
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}
