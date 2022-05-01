package weekly291

import "math"

func minimumCardPickup(cards []int) int {
	ans := math.MaxInt32
	dict := map[int]int{}
	for i, card := range cards {
		if index, ok := dict[card]; ok {
			ans = min(ans, i-index+1)
		}
		dict[card] = i
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
