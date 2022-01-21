package golang

import "math"

/*
 * @lc app=leetcode.cn id=1345 lang=golang
 *
 * [1345] Jump Game IV
 */

// @lc code=start
func minJumps(arr []int) int {
	n := len(arr)
	mapping := map[int][]int{} // value - index
	for i, v := range arr {
		mapping[v] = append(mapping[v], i)
	}
	visited := map[int]bool{0: true} // visited index
	ans, count := math.MaxInt32, 0
	var backtrack func(int)
	backtrack = func(index int) {
		if index == n-1 {
			if count < ans {
				ans = count
			}
			return
		}
		positions := map[int]bool{}
		if index < n-1 {
			positions[index+1] = true
		}
		for _, pos := range mapping[arr[index]] {
			positions[pos] = true
		}
		if index > 0 {
			positions[index-1] = true
		}
		for next := range positions {
			if !visited[next] && count < ans {
				count++
				visited[next] = true
				backtrack(next)
				count--
				visited[next] = false
			}
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
