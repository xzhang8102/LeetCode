package weekly288

func largestInteger(num int) int {
	digits := [10]int{}
	parity := []int{}
	for num > 0 {
		parity = append(parity, (num%10)&1)
		digits[num%10]++
		num /= 10
	}
	ans := 0
	odd, even := 9, 8
	for i := len(parity) - 1; i >= 0; i-- {
		if parity[i] == 1 {
			for odd >= 1 && digits[odd] == 0 {
				odd -= 2
			}
			ans = ans*10 + odd
			digits[odd]--
		} else {
			for even >= 0 && digits[even] == 0 {
				even -= 2
			}
			ans = ans*10 + even
			digits[even]--
		}
	}
	return ans
}
