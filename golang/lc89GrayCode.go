package golang

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
import "strconv"

func grayCode(n int) []int {
	ans := []int{0}
	visited := map[int]bool{
		0: true,
	}
	pick := make([]byte, n)
	for i := range pick {
		pick[i] = '0'
	}
	var backtrack func(index int, bit byte)
	backtrack = func(index int, bit byte) {
		prevBit := pick[index]
		pick[index] = bit
		parsed, _ := strconv.ParseInt(string(pick), 2, n+1)
		num := int(parsed)
		if !visited[num] {
			ans = append(ans, num)
			visited[num] = true
			for i := 0; i < n; i++ {
				if i != index {
					if pick[i] == '1' {
						backtrack(i, '0')
					} else {
						backtrack(i, '1')
					}
				}
			}
		} else {
			pick[index] = prevBit
		}
	}
	backtrack(n-1, '1')
	return ans
}

// @lc code=end
