package golang

/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] Number of Operations to Make Network Connected
 */

// @lc code=start
type lc1319UF struct {
	parents     []int
	connections int
}

func newUF(size int) *lc1319UF {
	parents := make([]int, size)
	for i := range parents {
		parents[i] = -1
	}
	return &lc1319UF{parents, size}
}

func (uf *lc1319UF) find(x int) int {
	if uf.parents[x] >= 0 {
		uf.parents[x] = uf.find(uf.parents[x])
		return uf.parents[x]
	}
	return x
}

func (uf *lc1319UF) union(a, b int) {
	rootA, rootB := uf.find(a), uf.find(b)
	if rootA != rootB {
		sizeA, sizeB := -uf.parents[rootA], -uf.parents[rootB]
		if sizeA >= sizeB {
			uf.parents[rootA] -= sizeB
			uf.parents[rootB] = rootA
		} else {
			uf.parents[rootB] -= sizeA
			uf.parents[rootA] = rootB
		}
		uf.connections--
	}
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	uf := newUF(n)
	for _, conn := range connections {
		uf.union(conn[0], conn[1])
	}
	return uf.connections - 1
}

// @lc code=end
