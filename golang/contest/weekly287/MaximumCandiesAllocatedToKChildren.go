package weekly287

func maximumCandies(candies []int, k int64) int {
	total := 0
	largest := 0
	for _, candy := range candies {
		total += candy
		largest = max(largest, candy)
	}
	if int64(total) < k {
		return 0
	}
	lo, hi := 1, largest
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if check(candies, k, mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo - 1
}

func check(candies []int, k int64, avg int) bool {
	for _, candy := range candies {
		k -= int64(candy / avg)
		if k <= 0 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
