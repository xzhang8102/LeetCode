package golang

/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	graph := make([]map[int]bool, n)
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = map[int]bool{}
		}
		graph[edge[0]][edge[1]] = true
		if graph[edge[1]] == nil {
			graph[edge[1]] = map[int]bool{}
		}
		graph[edge[1]][edge[0]] = true
	}
	leaves := []int{}
	for i, degree := range graph {
		if len(degree) <= 1 {
			leaves = append(leaves, i)
		}
	}
	for n > 2 {
		n -= len(leaves)
		size := len(leaves)
		for i := 0; i < size; i++ {
			leaf := leaves[i]
			for neighbor := range graph[leaf] {
				delete(graph[neighbor], leaf)
				if len(graph[neighbor]) == 1 {
					leaves = append(leaves, neighbor)
				}
			}
		}
		leaves = leaves[size:]
	}
	return leaves
}

// @lc code=end
