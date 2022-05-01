package main

// description: https://leetcode-cn.com/problems/VdG6tT/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	cnt := 0
	var N, K int
	var prefixSum []int
	for sc.Scan() {
		v, _ := strconv.Atoi(sc.Text())
		if cnt == 0 {
			N = v
			prefixSum = make([]int, N+1)
		} else if cnt == 1 {
			K = v
		} else {
			prefixSum[cnt-1] = prefixSum[cnt-2] + v
		}
		cnt++
	}
	fmt.Println(divide(prefixSum, N, K))
}

func divide(prefixSum []int, N, K int) int {
	dp := make([]int, N+1)
	for j := 1; j <= N-(K-1); j++ {
		dp[j] = prefixSum[j] * j
	}
	for i := 2; i <= K-1; i++ {
		tmp := make([]int, N+1)
		for j := i; j <= N-(K-i); j++ {
			// possible optimization: Convex Hull Trick
			tmp[j] = math.MaxInt64
			for k := i - 1; k < j; k++ {
				tmp[j] = min(tmp[j], dp[k]+(prefixSum[j]-prefixSum[k])*(j-k))
			}
		}
		dp = tmp
	}
	ans := math.MaxInt64
	for j := K - 1; j <= N-1; j++ {
		ans = min(ans, dp[j]+(prefixSum[N]-prefixSum[j])*(N-j))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
