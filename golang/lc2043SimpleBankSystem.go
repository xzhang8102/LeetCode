package golang

import "math"

/*
 * @lc app=leetcode.cn id=2043 lang=golang
 *
 * [2043] Simple Bank System
 */

// @lc code=start
type Bank struct {
	balance []int64
}

func lc2043Constructor(balance []int64) Bank {
	return Bank{
		balance,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 < 1 || account1 > len(this.balance) {
		return false
	}
	if account2 < 1 || account2 > len(this.balance) {
		return false
	}
	if this.balance[account1-1] < money {
		return false
	}
	if math.MaxInt64-this.balance[account2-1] < money {
		return false
	}
	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account < 1 || account > len(this.balance) {
		return false
	}
	if math.MaxInt64-this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account < 1 || account > len(this.balance) {
		return false
	}
	if this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
// @lc code=end
