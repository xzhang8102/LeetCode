package lccup2022spring

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNumber(root *TreeNode, ops [][]int) int {
	if root == nil {
		return 0
	}
	paintRed := [][2]int{}
	start := false
	for _, op := range ops {
		if !start {
			if op[0] == 1 {
				start = true
				paintRed = append(paintRed, [2]int{op[1], op[2]})
			}
		} else {
			newInterval := [2]int{op[1], op[2]}
			tmp := [][2]int{}
			if op[0] == 0 { // paint blue
				for _, interval := range paintRed {
					if interval[1] < newInterval[0] || interval[0] > newInterval[1] {
						tmp = append(tmp, interval)
					} else {
						if newInterval[0] > interval[0] {
							tmp = append(tmp, [2]int{interval[0], newInterval[0] - 1})
						}
						if newInterval[1] < interval[1] {
							tmp = append(tmp, [2]int{newInterval[1] + 1, interval[1]})
						}
					}
				}
			} else {
				added := false
				for _, interval := range paintRed {
					if interval[1] < newInterval[0]-1 {
						tmp = append(tmp, interval)
					} else if interval[0] > newInterval[1]+1 {
						if !added {
							added = true
							tmp = append(tmp, newInterval)
						}
						tmp = append(tmp, interval)
					} else {
						newInterval[0] = p03Min(newInterval[0], interval[0])
						newInterval[1] = p03Max(newInterval[1], interval[1])
					}
				}
				if !added {
					tmp = append(tmp, newInterval)
				}
			}
			paintRed = tmp
		}
	}
	ans := 0
	var paint func(node *TreeNode, interval [2]int)
	paint = func(node *TreeNode, interval [2]int) {
		if node == nil {
			return
		}
		if node.Val < interval[0] {
			paint(node.Right, interval)
		} else if node.Val > interval[1] {
			paint(node.Left, interval)
		} else {
			ans++
			paint(node.Left, interval)
			paint(node.Right, interval)
		}
	}
	for _, interval := range paintRed {
		paint(root, interval)
	}
	return ans
}

func p03Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func p03Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
