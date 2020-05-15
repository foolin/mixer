package mixer

import (
	"sync"
)

const (
	LcgDefaultMudu int64 = 4294967296 //math.Pow(2.0, 32.0)
	LcgDefaultMult int64 = 1103515245
	LcgDefaultInc  int64 = 12345
	LcgMaxRand     int64 = 4294967295
)

//linear congruential generator
type LCGRandom struct {
	state      int64
	modulus    int64
	multiplier int64
	increment  int64
	lock       sync.Mutex
}

func NewLGC(seed int64) *LCGRandom {
	return NewLGCWith(seed, LcgDefaultMudu, LcgDefaultMult, LcgDefaultInc)
}

func NewLGCWith(seed, modulus, multiplier, increment int64) *LCGRandom {
	return &LCGRandom{
		state:      seed % LcgMaxRand,
		modulus:    modulus,
		multiplier: multiplier,
		increment:  increment,
	}
}

func (g *LCGRandom) Int64() int64 {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.state = (g.multiplier*g.state + g.increment) % g.modulus
	return g.state
}

func (g *LCGRandom) Int64n(n int64) int64 {
	return g.Int64() % n
}

func (g *LCGRandom) Intn(n int) int {
	return int(g.Int64n(int64(n)))
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (g *LCGRandom) Perm(n int) []int {
	m := make([]int, n)
	// In the following loop, the iteration when i=0 always swaps m[0] with m[0].
	// A change to remove this useless iteration is to assign 1 to i in the init
	// statement.
	for i := 0; i < n; i++ {
		j := g.Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}
