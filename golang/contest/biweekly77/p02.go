package biweekly77

import "math"

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	ans := 0
	min := math.MaxInt64
	for i := range nums {
		avg1 := prefixSum[i+1] / (i + 1)
		avg2 := 0
		if i < n-1 {
			avg2 = (prefixSum[n] - prefixSum[i+1]) / (n - i - 1)
		}
		if diff := abs(avg1 - avg2); diff < min {
			if diff == 0 {
				return i
			}
			ans = i
			min = diff
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
