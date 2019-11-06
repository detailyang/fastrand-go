package fastrand

import (
	"testing"
)

func TestFastRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		FastRand()
		FastRandN(1)
	}
}
