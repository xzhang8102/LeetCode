package golang

/*
 * @lc app=leetcode.cn id=1868 lang=golang
 *
 * [1868] Product of Two Run-Length Encoded Arrays
 */

// @lc code=start
func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	ans := [][]int{}
	for p1, p2 := 0, 0; p1 < len(encoded1) && p2 < len(encoded2); {
		val1, freq1 := encoded1[p1][0], encoded1[p1][1]
		val2, freq2 := encoded2[p2][0], encoded2[p2][1]
		if len(ans) == 0 || ans[len(ans)-1][0] != val1*val2 {
			ans = append(ans, []int{val1 * val2, lc1868Min(freq1, freq2)})
		} else {
			ans[len(ans)-1][1] += lc1868Min(freq1, freq2)
		}
		if freq1 < freq2 {
			p1++
			encoded2[p2][1] -= freq1
		} else {
			p2++
			encoded1[p1][1] -= freq2
		}
		if encoded1[p1][1] == 0 {
			p1++
		}
	}
	return ans
}

func lc1868Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
