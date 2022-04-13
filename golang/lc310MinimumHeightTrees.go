package golang

/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = []int{}
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// 往下的最大高度
	down1 := make([]int, n)
	// 往下的次大高度
	down2 := make([]int, n)
	// 往上的最大高度
	up := make([]int, n)
	// 记录down1[cur]由哪个孩子节点转移
	p := make([]int, n)
	for i := range p {
		p[i] = -1
	}
	var dfs1 func(cur, fa int) int
	dfs1 = func(cur, fa int) int {
		for _, neighbour := range graph[cur] {
			if neighbour == fa {
				continue
			}
			height := dfs1(neighbour, cur) + 1
			if height > down1[cur] {
				down2[cur] = down1[cur]
				down1[cur] = height
				p[cur] = neighbour
			} else if height > down2[cur] {
				down2[cur] = height
			}
		}
		return down1[cur]
	}

	var dfs2 func(cur, fa int)
	dfs2 = func(cur, fa int) {
		for _, neighbour := range graph[cur] {
			if neighbour == fa {
				continue
			}
			if p[cur] != neighbour {
				up[neighbour] = lc310Max(up[neighbour], down1[cur]+1)
			} else {
				up[neighbour] = lc310Max(up[neighbour], down2[cur]+1)
			}
			up[neighbour] = lc310Max(up[neighbour], up[cur]+1)
			dfs2(neighbour, cur)
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)
	ans := []int{}
	min := n + 1
	for i := 0; i < n; i++ {
		height := lc310Max(down1[i], up[i])
		if height < min {
			min = height
			ans = []int{i}
		} else if height == min {
			ans = append(ans, i)
		}
	}
	return ans
}

func lc310Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
