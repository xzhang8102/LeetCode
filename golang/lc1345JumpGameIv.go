package golang

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
	type pair struct {
		idx, step int
	}
	q := []pair{{0, 0}}
	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		if point.idx == n-1 {
			return point.step
		}
		for _, pos := range mapping[arr[point.idx]] {
			if !visited[pos] {
				visited[pos] = true
				q = append(q, pair{pos, point.step + 1})
			}
		}
		delete(mapping, arr[point.idx])
		if !visited[point.idx+1] {
			visited[point.idx+1] = true
			q = append(q, pair{point.idx + 1, point.step + 1})
		}
		if point.idx > 0 && !visited[point.idx-1] {
			visited[point.idx-1] = true
			q = append(q, pair{point.idx - 1, point.step + 1})
		}
	}
	return -1
}

// @lc code=end
