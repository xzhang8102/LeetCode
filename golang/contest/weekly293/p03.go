package weekly293

func largestCombination(candidates []int) int {
	ans := 0
	for i := 0; i < 24; i++ {
		cnt := 0
		for _, candidate := range candidates {
			cnt += (candidate >> i) & 1
		}
		ans = max(ans, cnt)
	}
	return ans
}
