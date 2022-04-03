package biweekly75

func triangularSum(nums []int) int {
	for len(nums) > 1 {
		tmp := make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			tmp[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = tmp
	}
	return nums[0]
}
