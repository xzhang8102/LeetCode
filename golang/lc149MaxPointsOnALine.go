package golang

import "fmt"

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] Max Points on a Line
 */

// @lc code=start
func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	ans := 2
	for i, p1 := range points {
		slope := map[string]int{}
		for _, p2 := range points[i+1:] {
			x, y := p1[0]-p2[0], p1[1]-p2[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := lc149Gcd(lc149Abs(x), lc149Abs(y))
				x /= g
				y /= g
			}
			k := fmt.Sprintf("%d-%d", y, x)
			if count, ok := slope[k]; ok {
				slope[k] = count + 1
			} else {
				slope[k] = 2
			}
		}
		for _, c := range slope {
			if c > ans {
				ans = c
			}
		}
	}
	return ans
}

func lc149Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lc149Gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}
	return y
}

// @lc code=end
