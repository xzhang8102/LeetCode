package golang

/*
 * @lc app=leetcode.cn id=359 lang=golang
 *
 * [359] Logger Rate Limiter
 */

// @lc code=start
type Logger struct {
	history map[string]int
}

func lc359Constructor() Logger {
	return Logger{map[string]int{}}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := this.history[message]; !ok {
		this.history[message] = timestamp
		return true
	} else {
		if timestamp < this.history[message]+10 {
			return false
		} else {
			this.history[message] = timestamp
			return true
		}
	}
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
// @lc code=end
