package golang

import "math"

/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	ans := []string{}
	dict := map[string]int{}
	for i, s := range list1 {
		if _, ok := dict[s]; !ok {
			dict[s] = i
		}
	}
	sum := math.MaxInt32
	for i2, s := range list2 {
		if i1, ok := dict[s]; ok {
			if i1+i2 < sum {
				sum = i1 + i2
				ans = []string{s}
			} else if i1+i2 == sum {
				ans = append(ans, s)
			}
		}
	}
	return ans
}

// @lc code=end
