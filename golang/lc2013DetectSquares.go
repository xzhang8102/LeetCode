package golang

/*
 * @lc app=leetcode.cn id=2013 lang=golang
 *
 * [2013] Detect Squares
 */

// @lc code=start
type DetectSquares struct {
	coords map[int]map[int]int // x : y : count
}

func Constructor() DetectSquares {
	return DetectSquares{
		map[int]map[int]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if this.coords[x] == nil {
		this.coords[x] = map[int]int{}
	}
	this.coords[x][y]++
}

func (this *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	if this.coords[x] == nil {
		return 0
	}
	ans := 0
	yMap := this.coords[x]
	for edgeX, edgeYMap := range this.coords {
		if edgeX != x {
			d := x - edgeX
			ans += yMap[y+d] * edgeYMap[y+d] * edgeYMap[y]
			ans += yMap[y-d] * edgeYMap[y-d] * edgeYMap[y]
		}
	}
	return ans
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
// @lc code=end
