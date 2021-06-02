package ncryp

func IsPrime(x uint64) bool {
	for d := uint64(2); d < x; d++ {
		if x%d == 0 {
			return false
		}
	}
	return true
}
