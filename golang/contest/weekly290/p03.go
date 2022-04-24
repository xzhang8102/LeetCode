package weekly290

import (
	"sort"
)

func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][1] > rectangles[j][1]
	})
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] > points[j][1] })

	ans := make([]int, len(points))
	i, n, xSelection := 0, len(rectangles), []int{}
	for _, p := range points {
		start := i
		// keep finding the rectangles which have larger heights
		for ; i < n && p[1] <= rectangles[i][1]; i++ {
			// append corresponding length
			xSelection = append(xSelection, rectangles[i][0])
		}
		// sort only when adding new length
		if start < i {
			sort.Ints(xSelection)
		}
		ans[p[2]] = len(xSelection) - sort.SearchInts(xSelection, p[0])
	}
	return ans
}
