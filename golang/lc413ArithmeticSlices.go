package golang

/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	valid := []int{}
	ans := 0
	for i := 1; i < n-1; i++ {
		if nums[i]-nums[i-1] == nums[i+1]-nums[i] {
			valid = append(valid, i)
			ans++
		}
	}
	for _, index := range valid {
		gap := nums[index] - nums[index-1]
		for i := index + 2; i < n; i++ {
			if nums[i]-nums[i-1] == gap {
				ans++
			} else {
				break
			}
		}
	}
	return ans
}

// @lc code=end
