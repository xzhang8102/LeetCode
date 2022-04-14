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

	distBegin := make([]int, len(mapping))
	distEnd := make([]int, len(mapping))
	for i := range distBegin {
		distBegin[i] = math.MaxInt64
		distEnd[i] = math.MaxInt64
	}
	distBegin[beginId] = 0
	distEnd[endId] = 0
	qBegin := []int{beginId}
	qEnd := []int{endId}
	for len(qBegin) > 0 && len(qEnd) > 0 {
		q := qBegin
		qBegin = nil
		for _, v := range q {
			if distEnd[v] < math.MaxInt64 {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, neighbour := range graph[v] {
				if distBegin[neighbour] == math.MaxInt64 {
					distBegin[neighbour] = distBegin[v] + 1
					qBegin = append(qBegin, neighbour)
				}
			}
		}
		q = qEnd
		qEnd = nil
		for _, v := range q {
			if distBegin[v] < math.MaxInt64 {
				return (distBegin[v]+distEnd[v])/2 + 1
			}
			for _, neighbour := range graph[v] {
				if distEnd[neighbour] == math.MaxInt64 {
					distEnd[neighbour] = distEnd[v] + 1
					qEnd = append(qEnd, neighbour)
				}
			}
		}
	}
	return 0
}

// @lc code=end
