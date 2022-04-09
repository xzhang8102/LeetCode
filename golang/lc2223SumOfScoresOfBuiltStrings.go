package golang

/*
 * @lc app=leetcode.cn id=2223 lang=golang
 *
 * [2223] Sum of Scores of Built Strings
 */

// @lc code=start
func sumScores(s string) int64 {
	var sum int64 = 0
	n := len(s)
	z := make([]int, n)
	left, right := 0, 0
	for i := 1; i < n; i++ {
		if i > right {
			left, right = i, i
			for right < n && s[right] == s[right-left] {
				right++
			}
			z[i] = right - left
			right--
		} else {
			if i+z[i-left]-1 < right {
				z[i] = z[i-left]
			} else {
				left = i
				for right < n && s[right] == s[right-left] {
					right++
				}
				z[i] = right - left
				right--
			}
		}
		sum += int64(z[i])
	}
	sum += int64(n)
	return sum
}

// @lc code=end
