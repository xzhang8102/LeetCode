package offer

func findRepeatNumber(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num/n > 1 {
			return i
		}
		num %= n
		if nums[num]/n > 0 {
			return num
		}
		nums[num] += n
	}
	return -1
}
