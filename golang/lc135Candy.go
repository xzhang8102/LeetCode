package golang

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {
	n := len(ratings)
	dec := 1
	ans := 1
	pre := 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			if dec > 1 {
				// start: 1, end: dec - 1, length: dec - 1
				ans += dec * (dec - 1) / 2
				if dec > pre {
					ans += dec - pre
				}
				pre = 1
			}
			dec = 1
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
		} else {
			dec++
		}
	}
	if dec > 1 {
		ans += dec * (dec - 1) / 2
		if dec > pre {
			ans += dec - pre
		}
	}
	return ans
}

// @lc code=end
