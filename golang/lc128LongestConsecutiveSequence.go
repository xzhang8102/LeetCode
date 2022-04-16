package golang

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	dict := map[int]int{}
	ans := 0
	for _, num := range nums {
		if _, ok := dict[num]; !ok {
			left := dict[num-1]
			right := dict[num+1]
			streak := left + 1 + right
			dict[num] = streak
			if streak > ans {
				ans = streak
			}
			dict[num-left] = streak
			dict[num+right] = streak
		}
	}
	return ans
}

// @lc code=end
