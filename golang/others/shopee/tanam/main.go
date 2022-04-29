package main

// description: https://leetcode.cn/problems/m8KnhV/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	line := 0
	var T, N, M int
	var data [][][]int
	for i, t := 0, 0; sc.Scan(); i++ {
		if line == 0 {
			T, _ = strconv.Atoi(sc.Text())
			data = make([][][]int, T)
			line++
		} else {
			raw := strings.Split(sc.Text(), " ")
			if i == line {
				N, _ = strconv.Atoi(raw[0])
				M, _ = strconv.Atoi(raw[1])
				data[t] = make([][]int, N)
				t++
				line += N + 1
			} else {
				data[t-1][i-line+N] = make([]int, M)
				for j := range data[t-1][i-line+N] {
					data[t-1][i-line+N][j], _ = strconv.Atoi(raw[j])
				}
			}
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	for _, input := range data {
		fmt.Println(weirdGame(input))
	}
}

func weirdGame(input [][]int) int {
	N, M := len(input), len(input[0])
	dp := make([][2]int, N)
	sum := 0
	for _, fruit := range input[0] {
		sum += fruit
		dp[0][0] = max(dp[0][0], sum)
	}
	dp[0][1] = sum
	for i := 1; i < N; i++ {
		sum = 0
		maxLeft := 0
		for _, fruit := range input[i] {
			sum += fruit
			maxLeft = max(maxLeft, sum)
		}
		dp[i][0] = max(dp[i-1][0]+maxLeft, dp[i-1][1]+sum)
		sum = 0
		maxRight := 0
		for j := M - 1; j >= 0; j-- {
			sum += input[i][j]
			maxRight = max(maxRight, sum)
		}
		dp[i][1] = max(dp[i-1][1]+maxRight, dp[i-1][0]+sum)
	}
	return max(dp[N-1][0], dp[N-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
