package numbers

func IsPrime(num int) bool {
	if num == 1 && num < 1 {
		return true
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}
