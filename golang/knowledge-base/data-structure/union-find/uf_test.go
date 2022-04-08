package main

import "testing"

func TestUnionFind_find(t *testing.T) {
	type fields struct {
		parents     []int
		Connections int
	}
	tests := []struct {
		name string
		want int
	}{
		{
			name: "make sure path compression is working",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewUnionFind(20)
			uf.Union(0, 1)
			uf.Union(0, 2)
			uf.Union(1, 2)
			uf.Union(2, 10)
			uf.Union(10, 15)
			uf.Union(16, 17)
			uf.Union(18, 19)
			uf.Union(16, 19)
			uf.Union(10, 19)
			if got := uf.find(10); got != tt.want {
				t.Errorf("UnionFind.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
