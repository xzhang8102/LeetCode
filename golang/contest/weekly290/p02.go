package weekly290

func countLatticePoints(circles [][]int) int {
	set := map[[2]int]bool{}
	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		for a := x - r; a <= x+r; a++ {
			for b := y - r; b <= y+r; b++ {
				if (a-x)*(a-x)+(y-b)*(y-b) <= r*r {
					set[[2]int{a, b}] = true
				}
			}
		}
	}
	return len(set)
}
