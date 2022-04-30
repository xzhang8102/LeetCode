package golang

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		i := (num - 1) % n
		nums[i] += n
	}
	ans := []int{}
	for i, num := range nums {
		if num <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// @lc code=end
