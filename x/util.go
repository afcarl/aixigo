package x

import (
	"math"
	"math/rand"
	"runtime"
	"time"
)

// NewPRN creates a new rand.Rand object with its own seed
func NewPRN() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

//ArgMax returns the argmax (index) for an array of floats
func ArgMax(A []float64) (int, float64) {
	max := math.Inf(-1)
	idx := -1
	for i, v := range A {
		if v > max {
			max = v
			idx = i
		}
	}
	return idx, max
}

//ToInt helper for Observation objects
func ToInt(o ObservationBits) Observation {
	n := 0
	for _, b := range o {
		n <<= 1
		if b {
			n++
		}
	}
	return Observation(n)
}

// Equals checks equality of percepts
func Equals(e, p *Percept) bool {
	return p.R == e.R && p.O == e.O
}

// NumCPU does what is expected
func NumCPU() int {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)
	return cpu
}

// RLUtility is the utility function for normal reward-based reinforcement learners
func RLUtility(o Observation, r Reward, dfr int) float64 {
	return float64(r)
}

// Log2 for integers to hackishly improve performance in MCTS
func Log2(v uint) int {
	n := -1
	for ; v > 0; n++ {
		v >>= 1
	}
	return n
}

// Entropy (Shannon)
func Entropy(p []float64) float64 {
	ent := 0.0
	for _, v := range p {
		if v == 0 {
			continue
		}
		ent -= v * math.Log2(v)
	}
	return ent
}
