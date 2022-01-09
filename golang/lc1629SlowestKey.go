package golang

/*
 * @lc app=leetcode.cn id=1629 lang=golang
 *
 * [1629] Slowest Key
 */

// @lc code=start
func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans := keysPressed[0]
	maxTime := releaseTimes[0]
	for i := 1; i < len(keysPressed); i++ {
		time := releaseTimes[i] - releaseTimes[i-1]
		if time > maxTime {
			ans = keysPressed[i]
			maxTime = time
		} else if time == maxTime && ans < keysPressed[i] {
			ans = keysPressed[i]
		}
	}
	return ans
}

// @lc code=end
