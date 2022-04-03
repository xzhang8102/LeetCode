package weekly287

import "sort"

func findWinners(matches [][]int) [][]int {
	type record struct{ win, lose int }
	records := map[int]*record{}
	for _, match := range matches {
		if records[match[0]] == nil {
			records[match[0]] = &record{1, 0}
		} else {
			records[match[0]].win++
		}
		if records[match[1]] == nil {
			records[match[1]] = &record{0, 1}
		} else {
			records[match[1]].lose++
		}
	}
	ans := make([][]int, 2)
	ans[0], ans[1] = []int{}, []int{}
	for k, v := range records {
		if v.lose == 0 {
			ans[0] = append(ans[0], k)
		} else if v.lose == 1 {
			ans[1] = append(ans[1], k)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
