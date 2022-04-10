package weekly288

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
	min := math.MaxInt64
	ans := ""
	arr := strings.Split(expression, "+")
	left, right := arr[0], arr[1]
	for i := 0; i < len(left); i++ {
		n1 := 1
		if i > 0 {
			n1, _ = strconv.Atoi(left[:i])
		}
		n2, _ := strconv.Atoi(left[i:])
		for j := 1; j <= len(right); j++ {
			n4 := 1
			if j < len(right) {
				n4, _ = strconv.Atoi(right[j:])
			}
			n3, _ := strconv.Atoi(right[:j])
			if val := n1 * (n2 + n3) * n4; val < min {
				min = val
				ans = fmt.Sprintf("%s(%s+%s)%s", left[:i], left[i:], right[:j], right[j:])
			}
		}
	}
	return ans
}
