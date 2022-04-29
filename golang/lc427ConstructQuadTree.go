package golang

/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] Construct Quad Tree
 */

/**
 * Definition for a QuadTree node.
 */
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// @lc code=start
func construct(grid [][]int) *Node {
	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, start, end int) *Node {
		for _, row := range rows {
			for _, v := range row[start:end] {
				if v != rows[0][start] {
					rMid := len(rows) / 2
					cMid := (start + end) / 2
					return &Node{
						false,
						false,
						dfs(rows[:rMid], start, cMid),
						dfs(rows[:rMid], cMid, end),
						dfs(rows[rMid:], start, cMid),
						dfs(rows[rMid:], cMid, end),
					}
				}
			}
		}
		return &Node{Val: rows[0][start] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}

// @lc code=end
