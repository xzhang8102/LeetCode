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
		for j := 1; j <= len(right); j++ {
			var n1, n2, n3, n4 int
			if i == 0 {
				n1 = 1
			} else {
				n1, _ = strconv.Atoi(left[:i])
			}
			n2, _ = strconv.Atoi(left[i:])
			if j == len(right) {
				n4 = 1
			} else {
				n4, _ = strconv.Atoi(right[j:])
			}
			n3, _ = strconv.Atoi(right[:j])
			if n1*(n2+n3)*n4 < min {
				min = n1 * (n2 + n3) * n4
				ans = fmt.Sprintf("(%d+%d)", n2, n3)
				if i > 0 {
					ans = fmt.Sprintf("%d%s", n1, ans)
				}
				if j < len(right) {
					ans = fmt.Sprintf("%s%d", ans, n4)
				}
			}
		}
	}
	return ans
}
