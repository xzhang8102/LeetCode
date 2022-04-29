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
	if len(grid) == 1 {
		return &Node{
			Val:    grid[0][0] == 1,
			IsLeaf: true,
		}
	}
	root := &Node{}
	same := true
	for _, row := range grid {
		for _, v := range row {
			if v != grid[0][0] {
				same = false
				break
			}
		}
		if !same {
			break
		}
	}
	if same {
		root.IsLeaf = true
		root.Val = grid[0][0] == 1
	} else {
		n := len(grid)
		tmp := make([][]int, n/2)
		for i := range tmp {
			tmp[i] = grid[i][:n/2]
		}
		root.TopLeft = construct(tmp)
		for i := range tmp {
			tmp[i] = grid[i][n/2:]
		}
		root.TopRight = construct(tmp)
		for i := range tmp {
			tmp[i] = grid[n/2+i][:n/2]
		}
		root.BottomLeft = construct(tmp)
		for i := range tmp {
			tmp[i] = grid[n/2+i][n/2:]
		}
		root.BottomRight = construct(tmp)
	}
	return root
}

// @lc code=end
