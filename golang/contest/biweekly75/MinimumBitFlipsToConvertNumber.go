package biweekly75

import "math/bits"

func minBitFlips(start int, goal int) int {
	if start == goal {
		return 0
	}
	return bits.OnesCount(uint(start ^ goal))
}
