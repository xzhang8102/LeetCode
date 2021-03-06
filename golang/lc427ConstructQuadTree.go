package golang

/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

/**
 * Definition for a QuadTree node.
 */
type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

// @lc code=start
func construct(grid [][]int) *QuadTreeNode {
	n := len(grid)
	pre := make([][]int, n+1)
	pre[0] = make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			pre[i+1][j+1] = grid[i][j] + pre[i+1][j] + pre[i][j+1] - pre[i][j]
		}
	}
	var dfs func(r0, r1, c0, c1 int) *QuadTreeNode
	dfs = func(r0, r1, c0, c1 int) *QuadTreeNode {
		total := pre[r1][c1] - pre[r1][c0] - pre[r0][c1] + pre[r0][c0]
		if total == 0 || total == (r1-r0)*(c1-c0) {
			return &QuadTreeNode{Val: grid[r0][c0] == 1, IsLeaf: true}
		}
		return &QuadTreeNode{
			false,
			false,
			dfs(r0, (r0+r1)/2, c0, (c0+c1)/2),
			dfs(r0, (r0+r1)/2, (c0+c1)/2, c1),
			dfs((r0+r1)/2, r1, c0, (c0+c1)/2),
			dfs((r0+r1)/2, r1, (c0+c1)/2, c1),
		}
	}
	return dfs(0, n, 0, n)
}

// @lc code=end
