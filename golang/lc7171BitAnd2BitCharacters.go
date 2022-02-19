package golang

/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1-bit and 2-bit Characters
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i < n-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == n-1
}

// @lc code=end
