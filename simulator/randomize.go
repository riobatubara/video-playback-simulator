package simulator

import (
	"math/rand"
	"sync"
	"time"
)

type Init struct {
	Rnd *rand.Rand
	Mtx *sync.Mutex
}

// Global pool of non-blocking *rand.Rand generators
var randPool = sync.Pool{
	New: func() interface{} {
		// Seeds a unique, isolated generator per thread to completely avoid mutexes
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

// Global shared instance remains matching your exact original declaration
var Randomize = &Init{
	Rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	Mtx: &sync.Mutex{},
}

// OPTIMIZATION: Converted from rune to byte array to drop memory usage by 75%
var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Thread-safe random integer without global mutex bottlenecks
func (r *Init) Intn(n int) int {
	localRand := randPool.Get().(*rand.Rand)
	val := localRand.Intn(n)
	randPool.Put(localRand) // Return generator back to the pool instantly
	return val
}

// Thread-safe random float64 (0.0 to <1.0)
func (r *Init) Float64() float64 {
	localRand := randPool.Get().(*rand.Rand)
	val := localRand.Float64()
	randPool.Put(localRand)
	return val
}

// Thread-safe shuffle
func (r *Init) Shuffle(n int, swap func(i, j int)) {
	localRand := randPool.Get().(*rand.Rand)
	localRand.Shuffle(n, swap)
	randPool.Put(localRand)
}

// GenerateSessionID prints out exact same format but avoids the 32x loop lock
func GenerateSessionID() string {
	const length = 32
	id := make([]byte, length)

	// OPTIMIZATION: Acquire the local randomizer ONCE instead of 32 separate times
	localRand := randPool.Get().(*rand.Rand)

	for i := 0; i < length; i++ {
		id[i] = letters[localRand.Intn(len(letters))]
	}

	randPool.Put(localRand) // Release back to pool

	// Converts the fast byte array directly to a final output string
	return string(id)
}

func RandomBitrate() (int, int) {
	videoBitrates := []int{1500, 2400, 3600, 4500, 6000}
	audioBitrates := []int{96, 128, 160, 192, 256}

	// Uses the lock-free Intn pool under the hood
	v := videoBitrates[Randomize.Intn(len(videoBitrates))] // example video bitrates in kbps
	a := audioBitrates[Randomize.Intn(len(audioBitrates))] // example audio bitrates in kbps

	return v, a
}
