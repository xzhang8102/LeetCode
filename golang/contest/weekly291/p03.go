package weekly291

import (
	"strconv"
	"strings"
)

func countDistinct(nums []int, k int, p int) int {
	set := map[string]bool{}
	var b strings.Builder
	for i := range nums {
		cnt := 0
		for j := i; j < len(nums); j++ {
			if nums[j]%p == 0 {
				cnt++
			}
			if cnt > k {
				break
			}
			if b.Len() > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(strconv.Itoa(nums[j]))
			set[b.String()] = true
		}
		b.Reset()
	}
	return len(set)
}
