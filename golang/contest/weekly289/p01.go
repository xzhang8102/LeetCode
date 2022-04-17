package weekly289

import (
	"fmt"
	"strings"
)

func digitSum(s string, k int) string {
	n := len(s)
	if n <= k {
		return s
	}
	var b strings.Builder
	for i := 0; i < n; {
		num := 0
		hi := i + k
		if hi > n {
			hi = n
		}
		for _, ch := range s[i:hi] {
			num += int(ch - '0')
		}
		b.WriteString(fmt.Sprintf("%d", num))
		i += k
	}
	return digitSum(b.String(), k)
}
