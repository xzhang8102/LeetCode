package golang

/*
 * @lc app=leetcode.cn id=765 lang=golang
 *
 * [765] Couples Holding Hands
 */

// @lc code=start
type lc765UF struct {
	parents     []int
	connections int
}

func newLC765UF(n int) *lc765UF {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = -1
	}
	return &lc765UF{parents, n}
}

func (uf *lc765UF) find(x int) int {
	if uf.parents[x] >= 0 {
		uf.parents[x] = uf.find(uf.parents[x])
		return uf.parents[x]
	}
	return x
}

func (uf *lc765UF) union(a, b int) {
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

func minSwapsCouples(row []int) int {
	n := len(row)
	uf := newLC765UF(n / 2)
	for i := 0; i < n; i += 2 {
		// if the adjacent two ids are a couple
		// row[i] / 2 == row[i+1] / 2
		uf.union(row[i]/2, row[i+1]/2)
	}
	return n/2 - uf.connections
}

// @lc code=end
