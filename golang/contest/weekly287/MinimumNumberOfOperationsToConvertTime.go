package weekly287

import (
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
	if current == correct {
		return 0
	}
	arr1 := strings.Split(current, ":")
	arr2 := strings.Split(correct, ":")
	hour1, _ := strconv.Atoi(arr1[0])
	min1, _ := strconv.Atoi(arr1[1])
	time1 := hour1*60 + min1
	hour2, _ := strconv.Atoi(arr2[0])
	min2, _ := strconv.Atoi(arr2[1])
	time2 := hour2*60 + min2
	if time1 > time2 {
		time2 += 23*60 + 59
	}
	diff := time2 - time1
	ans := 0
	if diff/60 > 0 {
		ans += diff / 60
		diff = diff % 60
	}
	if diff/15 > 0 {
		ans += diff / 15
		diff %= 15
	}
	if diff/5 > 0 {
		ans += diff / 5
		diff %= 5
	}
	if diff/1 > 0 {
		ans += diff
	}
	return ans
}
