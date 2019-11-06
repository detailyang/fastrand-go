//Package fastrand exports runtime.fastrand to public
package fastrand

import (
	_ "unsafe"
)

//go:linkname FastRand runtime.fastrand
func FastRand() uint32

//go:linkname FastRandN runtime.fastrandn
func FastRandN(n uint32) uint32
