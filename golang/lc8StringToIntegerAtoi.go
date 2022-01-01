package golang

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
import (
	"math"
	"unicode"
)

const (
	START = iota
	SIGNED
	IN_NUMBER
	END
)

const (
	SPACE = iota
	SIGN
	NUMBER
	OTHER
)

var table = [4][4]int{
	// ' ', +/-, number, other
	{START, SIGNED, IN_NUMBER, END}, // START
	{END, END, IN_NUMBER, END},      // SIGNED
	{END, END, IN_NUMBER, END},      // IN_NUMBER
	{END, END, END, END},            // OTHER
}

func myAtoi(s string) int {
	ans := 0
	negative := false
	state := START
	for i := 0; i < len(s); i++ {
		state = table[state][getCharCategory(s[i])]
		if state == IN_NUMBER {
			num := s[i] - '0'
			if negative {
				if int(num) >= ans*10-math.MinInt32 {
					return math.MinInt32
				}
				ans = ans*10 - int(num)
			} else {
				if int(num) >= math.MaxInt32-ans*10 {
					return math.MaxInt32
				}
				ans = ans*10 + int(num)
			}
		}
		if state == SIGNED && s[i] == '-' {
			negative = true
		}
		if state == END {
			return ans
		}
	}
	return ans
}

func getCharCategory(char byte) int {
	if unicode.IsSpace(rune(char)) {
		return SPACE
	} else if char == '-' || char == '+' {
		return SIGN
	} else if unicode.IsDigit(rune(char)) {
		return NUMBER
	} else {
		return OTHER
	}
}

// @lc code=end
