package golang

/*
 * @lc app=leetcode.cn id=1748 lang=golang
 *
 * [1748] Sum of Unique Elements
 */

// @lc code=start
func sumOfUnique(nums []int) int {
	hashmap := map[int]bool{}
	ans := 0
	for _, n := range nums {
		if unique, ok := hashmap[n]; ok {
			if unique {
				ans -= n
				hashmap[n] = false
			}
		} else {
			hashmap[n] = true
			ans += n
		}
	}
	return ans
}

// @lc code=end
