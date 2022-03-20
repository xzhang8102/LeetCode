package golang

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] Valid Number
 */

// @lc code=start
// status
const (
	LC65_START = iota
	LC65_INTEGER
	LC65_FRACTION_POINT // no digits followed
	LC65_DECIMAL
	LC65_EXP        // no digits followed
	LC65_EXP_SIGNED // no digits followed
	LC65_EXP_NUMBER
	LC65_SIGNED // no digits followed
	LC65_FAIL
)

// symbol
const (
	LC65_DIGIT = iota
	LC65_DOT
	LC65_SIGN
	LC65_OTHER_CHAR
	LC65_CHARACTER_E
)

var lc65table = [8][5]int{
	// 0-9   .    +/-     a-d,f-z      e
	{LC65_INTEGER, LC65_FRACTION_POINT, LC65_SIGNED, LC65_FAIL, LC65_FAIL}, // START
	{LC65_INTEGER, LC65_DECIMAL, LC65_FAIL, LC65_FAIL, LC65_EXP},           // INTEGER
	{LC65_DECIMAL, LC65_FAIL, LC65_FAIL, LC65_FAIL, LC65_FAIL},             // FRACTION_POINT
	{LC65_DECIMAL, LC65_FAIL, LC65_FAIL, LC65_FAIL, LC65_EXP},              // DECIMAL
	{LC65_EXP_NUMBER, LC65_FAIL, LC65_EXP_SIGNED, LC65_FAIL, LC65_FAIL},    // EXP
	{LC65_EXP_NUMBER, LC65_FAIL, LC65_FAIL, LC65_FAIL, LC65_FAIL},          // EXP_SIGNED
	{LC65_EXP_NUMBER, LC65_FAIL, LC65_FAIL, LC65_FAIL, LC65_FAIL},          // EXP_NUMBER
	{LC65_INTEGER, LC65_FRACTION_POINT, LC65_FAIL, LC65_FAIL, LC65_FAIL},   // SIGNED
}

func isNumber(s string) bool {
	state := LC65_START
	for i := 0; i < len(s); i++ {
		state = lc65table[state][lc65ParseChar(s[i])]
		if state == LC65_FAIL {
			return false
		}
	}
	return state == LC65_INTEGER || state == LC65_DECIMAL || state == LC65_EXP_NUMBER
}

func lc65ParseChar(ch byte) int {
	if ch >= '0' && ch <= '9' {
		return LC65_DIGIT
	} else if ch == 'e' || ch == 'E' {
		return LC65_CHARACTER_E
	} else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return LC65_OTHER_CHAR
	} else if ch == '+' || ch == '-' {
		return LC65_SIGN
	} else if ch == '.' {
		return LC65_DOT
	}
	return -1
}

// @lc code=end
