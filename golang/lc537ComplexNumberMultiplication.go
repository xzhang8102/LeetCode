package golang

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] Complex Number Multiplication
 */

// @lc code=start
func complexNumberMultiply(num1 string, num2 string) string {
	arr1, arr2 := strings.Split(num1[:len(num1)-1], "+"), strings.Split(num2[:len(num2)-1], "+")
	re1, _ := strconv.Atoi(arr1[0])
	img1, _ := strconv.Atoi(arr1[1])
	re2, _ := strconv.Atoi(arr2[0])
	img2, _ := strconv.Atoi(arr2[1])
	re := re1*re2 - img1*img2
	img := re1*img2 + re2*img1
	return fmt.Sprintf("%d+%di", re, img)
}

// @lc code=end
