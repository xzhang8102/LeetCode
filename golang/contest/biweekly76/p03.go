package biweekly76

type ATM struct {
	money [5]int
	cnt   []int
}

func Constructor() ATM {
	return ATM{[5]int{20, 50, 100, 200, 500}, make([]int, 5)}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, count := range banknotesCount {
		this.cnt[i] += count
	}
}

func (this *ATM) Withdraw(amount int) []int {
	tmp, ret := make([]int, 5), make([]int, 5)
	copy(tmp, this.cnt)
	for i := 4; i >= 0 && amount > 0; i-- {
		value := this.money[i]
		if amount >= value && tmp[i] > 0 {
			need := amount / value
			if tmp[i] <= need {
				amount -= tmp[i] * value
				ret[i] = tmp[i]
				tmp[i] = 0
			} else {
				amount %= value
				ret[i] = need
				tmp[i] -= need
			}
		}
	}
	if amount > 0 {
		return []int{-1}
	}
	copy(this.cnt, tmp)
	return ret
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
