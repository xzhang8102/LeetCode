package biweekly78

func waysToSplitArray(nums []int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		left := prefix[i+1]
		right := prefix[n] - prefix[i+1]
		if left >= right {
			ans++
		}
	}
	return ans
}
