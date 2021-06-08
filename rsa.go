package ncryp

func IsPrime(x uint64) bool {
	for d := uint64(2); d < x; d++ {
		if x%d == 0 {
			return false
		}
	}
	return true
}

func Erathosthenes(max uint64) []uint64 {
	isNotPrime := make([]bool, max+1)
	var result []uint64
	for i := uint64(2); i <= max; i++ {
		if isNotPrime[i] {
			continue
		}
		result = append(result, i)
		for j := 2 * i; j <= max; j += i {
			isNotPrime[j] = true
		}
	}
	return result
}
