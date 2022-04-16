package biweekly76

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	ans := 0
	for i := 0; i <= total/cost1; i++ {
		remain := total - i*cost1
		ans += remain/cost2 + 1
	}
	return int64(ans)
}
