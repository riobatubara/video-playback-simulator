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

// Global shared instance
var Randomize = &Init{
	Rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	Mtx: &sync.Mutex{},
}

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Thread-safe random integer
func (r *Init) Intn(n int) int {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()
	return r.Rnd.Intn(n)
}

// Thread-safe random float64 (0.0 to <1.0)
func (r *Init) Float64() float64 {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()
	return r.Rnd.Float64()
}

// Thread-safe shuffle
func (r *Init) Shuffle(n int, swap func(i, j int)) {
	r.Mtx.Lock()
	defer r.Mtx.Unlock()
	r.Rnd.Shuffle(n, swap)
}

func GenerateSessionID() string {
	const length = 32
	id := make([]rune, length)

	for i := range id {
		id[i] = runes[Randomize.Intn(len(runes))]
	}

	return string(id)
}

func RandomBitrate() (int, int) {
	videoBitrates := []int{1500, 2400, 3600, 4500, 6000} // example video bitrates in kbps
	audioBitrates := []int{96, 128, 160, 192, 256}       // example audio bitrates in kbps

	v := videoBitrates[Randomize.Intn(len(videoBitrates))]
	a := audioBitrates[Randomize.Intn(len(audioBitrates))]

	return v, a
}
