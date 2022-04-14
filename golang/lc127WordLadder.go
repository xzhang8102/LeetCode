package golang

import "math"

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	mapping := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := mapping[word]
		if !has {
			id = len(mapping)
			mapping[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}
	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := mapping[endWord]
	if !has {
		return 0
	}
	q := []int{beginId}
	dist := make([]int, len(mapping))
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[beginId] = 0
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, neighbour := range graph[v] {
			if dist[neighbour] == math.MaxInt64 {
				dist[neighbour] = dist[v] + 1
				q = append(q, neighbour)
			}
		}
	}
	return 0
}

// @lc code=end
