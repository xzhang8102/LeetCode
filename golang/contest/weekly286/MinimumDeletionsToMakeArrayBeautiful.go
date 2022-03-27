package weekly286

func minDeletion(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := 0
	for i := 1; i < n; {
		if nums[i] == nums[i-1] {
			j := i
			for j < n && nums[j] == nums[i] {
				j++
				ans++
			}
			i = j + 2
		} else {
			i += 2
		}
	}
	if (n-ans)&1 == 1 {
		return ans + 1
	}
	return ans
}
