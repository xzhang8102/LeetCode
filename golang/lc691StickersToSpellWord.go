package golang

import "strings"

/*
 * @lc app=leetcode.cn id=691 lang=golang
 *
 * [691] Stickers to Spell Word
 */

// @lc code=start
func minStickers(stickers []string, target string) int {
	set := map[rune]bool{}
	for _, ch := range target {
		set[ch] = true
	}
	n := len(stickers)
	available := make([]map[rune]int, 0, n)
	for _, sticker := range stickers {
		cnt := map[rune]int{}
		for _, ch := range sticker {
			if set[ch] {
				cnt[ch]++
			}
		}
		if len(cnt) > 0 {
			available = append(available, cnt)
		}
	}
	q, visited := []string{target}, map[string]int{target: 0}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, avl := range available {
			for _, char := range curr {
				if avl[char] > 0 {
					next := curr
					for c, cnt := range avl {
						next = strings.Replace(next, string(c), "", cnt)
					}
					if next == "" {
						return visited[curr] + 1
					}
					if _, ok := visited[next]; !ok {
						q = append(q, next)
						visited[next] = visited[curr] + 1
					}
					break
				}
			}
		}
	}
	return -1
}

// @lc code=end
