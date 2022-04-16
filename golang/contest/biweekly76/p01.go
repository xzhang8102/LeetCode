package biweekly76

import "math"

func findClosestNumber(nums []int) int {
	dist := math.MaxInt64
	ans := -1
	for _, num := range nums {
		d := num
		if num < 0 {
			d = -d
		}
		if d < dist {
			dist = d
			ans = num
		} else if d == dist {
			if num > ans {
				ans = num
			}
		}
	}
	return ans
}
