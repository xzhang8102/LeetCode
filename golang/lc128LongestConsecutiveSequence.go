package golang

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}
	ans := 0
	for _, num := range nums {
		if !set[num-1] { // num is the start of the sequence
			curr, streak := num, 1
			for set[curr+1] {
				curr++
				streak++
			}
			if ans < streak {
				ans = streak
			}
		}
	}
	return ans
}

// @lc code=end
