package golang

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	cache := map[byte]int{}
	ans := 0
	left := 0
	// 遍历的是滑动窗口的末端
	for i := 0; i < len(s); i++ {
		if index, seen := cache[s[i]]; seen {
			if index+1 > left {
				left = index + 1
			}
		}
		cache[s[i]] = i
		if i-left+1 > ans {
			ans = i - left + 1
		}
	}
	return ans
}

// @lc code=end
