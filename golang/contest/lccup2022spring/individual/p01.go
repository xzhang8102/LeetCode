package individual

func giveGem(gem []int, operations [][]int) int {
	for _, op := range operations {
		giver, taker := op[0], op[1]
		gift := gem[giver] / 2
		gem[giver] -= gift
		gem[taker] += gift
	}
	min, max := gem[0], gem[0]
	for _, g := range gem[1:] {
		if g > max {
			max = g
		}
		if g < min {
			min = g
		}
	}
	return max - min
}
