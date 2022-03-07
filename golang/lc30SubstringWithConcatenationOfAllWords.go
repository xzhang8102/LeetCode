package golang

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	window := map[string]int{}
	for _, w := range words {
		window[w]++
	}
	valid := func() bool {
		for _, v := range window {
			if v != 0 {
				return false
			}
		}
		return true
	}
	ans := []int{}
	n := len(s)
	step := len(words[0])
	gap := step * len(words)
	for slow, fast := 0, gap; fast <= n; slow, fast = slow+1, slow+gap+1 {
		i := slow
		for i < fast {
			if cnt := window[s[i:i+step]]; cnt > 0 {
				window[s[i:i+step]] = cnt - 1
				i += step
			} else {
				break
			}
		}
		if i == fast && valid() {
			ans = append(ans, slow)
		}
		for j := slow; j+step <= i; j += step {
			window[s[j:j+step]]++
		}
	}
	return ans
}

// @lc code=end
