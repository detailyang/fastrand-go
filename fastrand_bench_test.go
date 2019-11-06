package fastrand

import (
	"math/rand"
	"testing"
)

func BenchmarkFastRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := FastRand()
		_ = j
	}
}

func BenchmarkMathRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
