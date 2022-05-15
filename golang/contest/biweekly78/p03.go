package biweekly78

import "sort"

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })
	cover, left := 0, 0
	ans := 0
	for _, t := range tiles {
		tl, tr := t[0], t[1]
		cover += tr - tl + 1
		for tr-tiles[left][1]+1 > carpetLen {
			cover -= tiles[left][1] - tiles[left][0] + 1
			left++
		}
		if total := tr - tiles[left][0] + 1; total <= carpetLen {
			ans = max(ans, cover)
		} else {
			ans = max(ans, cover-(total-carpetLen))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
