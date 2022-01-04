package golang

/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] Day of the Week
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	// 31/12/1970 Thursday
	days := 0
	weekDays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	monthDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	if year > 1971 {
		for y := 1971; y < year; y++ {
			if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
				days += 366
			} else {
				days += 365
			}
		}
	}
	for m := 0; m < month-1; m++ {
		days += monthDays[m]
		if m == 1 {
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				days += 1
			}
		}
	}
	return weekDays[(days+day+3)%7]
}

// @lc code=end
