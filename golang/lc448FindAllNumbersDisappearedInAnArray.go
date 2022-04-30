package golang

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	freq := make([]int, n+1)
	for _, num := range nums {
		freq[num]++
	}
	ans := []int{}
	for i := 1; i <= n; i++ {
		if freq[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end
