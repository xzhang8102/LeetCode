package golang

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	slow, fast := n, lc202GetNext(n)
	for fast != 1 && fast != slow {
		slow = lc202GetNext(slow)
		fast = lc202GetNext(lc202GetNext(fast))
	}
	return fast == 1
}

func lc202GetNext(n int) int {
	tmp := 0
	for n > 0 {
		tmp += (n % 10) * (n % 10)
		n /= 10
	}
	return tmp
}

// @lc code=end
