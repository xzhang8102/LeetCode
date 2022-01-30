package golang

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	if m > n {
		return ""
	}
	min := n + 1
	var ans string
	cache := map[byte]int{}
	for i := 0; i < m; i++ {
		cache[t[i]]++
	}
	check := func() bool {
		for _, v := range cache {
			if v > 0 {
				return false
			}
		}
		return true
	}
	stack := []int{}
	for right := 0; right < n; right++ {
		if _, ok := cache[s[right]]; ok {
			stack = append(stack, right)
			cache[s[right]]--
		}
		for len(stack) >= m && check() {
			if right-stack[0]+1 < min {
				ans = s[stack[0] : right+1]
				min = len(ans)
			}
			cache[s[stack[0]]]++
			stack = stack[1:]
		}
	}
	return ans
}

// @lc code=end
