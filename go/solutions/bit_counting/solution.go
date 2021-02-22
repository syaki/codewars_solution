package solution

import "math/bits"

func CountBits(n uint) int {
	count := 0
	if n < 0 {
		n &= 0xffffffff
	}
	for n != 0 {
		count++
		n = (n - 1) & n
	}
	return count
}

func CleverCountBits(n uint) int {
	return bits.OnesCount(n)
}
