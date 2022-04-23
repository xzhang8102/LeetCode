package team

import (
	"container/list"
)

func conveyorBelt(matrix []string, start []int, end []int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	mapping := map[[2]int]rune{
		{-1, 0}: 'v',
		{0, -1}: '>',
		{0, 1}:  '<',
		{1, 0}:  '^',
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	deque := list.New()
	deque.PushBack(end)
	record := make([]int, n*m)
	for i := range record {
		record[i] = -1
	}
	record[end[0]*m+end[1]] = 0
	for deque.Len() > 0 {
		front := deque.Front()
		deque.Remove(front)
		pair := front.Value.([]int)
		r, c := pair[0], pair[1]
		if r == start[0] && c == start[1] {
			return record[r*m+c]
		}
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < n && newC >= 0 && newC < m {
				ops := record[r*m+c]
				if rune(matrix[newR][newC]) != mapping[dir] {
					ops++
				}
				if record[newR*m+newC] == -1 || record[newR*m+newC] > ops {
					record[newR*m+newC] = ops
					if ops == record[r*m+c] {
						deque.PushFront([]int{newR, newC})
					} else {
						deque.PushBack([]int{newR, newC})
					}
				}
			}
		}
	}
	return -1
}
