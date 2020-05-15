package mixer

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLCG(t *testing.T) {
	g := NewLGC(1024)
	list := make([]int64, 0)
	for i := 0; i < 100; i++ {
		list = append(list, g.Int64())
	}
	fmt.Println(list)
}

func TestLCGRandom_Perm(t *testing.T) {
	for i := 0; i < 10; i++ {
		g := NewLGC(1024)
		perms := g.Perm(20)
		t.Logf("i: %v", perms)
	}
}

func TestRandPerm(t *testing.T) {
	for i := 0; i < 10; i++ {
		rn := rand.New(rand.NewSource(1024))
		perms := rn.Perm(20)
		t.Logf("i: %v", perms)
	}
}
