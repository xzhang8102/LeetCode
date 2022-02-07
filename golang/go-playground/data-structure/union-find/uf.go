package main

import "fmt"

type UnionFind struct {
	parents     []int
	Connections int
}

func NewUnionFind(size int) *UnionFind {
	parents := make([]int, size)
	for i := range parents {
		parents[i] = -1 // every node is its root and the size of the tree is abs(-1)
	}
	return &UnionFind{parents, size}
}

func (uf *UnionFind) find(val int) int {
	if uf.parents[val] >= 0 {
		uf.parents[val] = uf.find(uf.parents[val]) // path compression
		return uf.parents[val]
	}
	return val
}

func (uf *UnionFind) Union(a, b int) {
	rootA, rootB := uf.find(a), uf.find(b)
	if rootA != rootB {
		sizeA, sizeB := -uf.parents[rootA], -uf.parents[rootB]
		if sizeA >= sizeB {
			uf.parents[rootB] = rootA
			uf.parents[rootA] -= sizeB
		} else {
			uf.parents[rootA] = rootB
			uf.parents[rootB] -= sizeA
		}
		uf.Connections--
	}
}

func (uf *UnionFind) IsConnected(a, b int) bool {
	return uf.find(a) == uf.find(b)
}

func main() {
	uf := NewUnionFind(10) // 0..9
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 9)
	fmt.Println(uf.IsConnected(0, 9))
	fmt.Println(uf.Connections)
}
