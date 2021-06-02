package ncryp

// ModPow calculates x^y modulo m.
func ModPow(x, y, m uint64) uint64 {
	var result uint64 = 1
	for y > 0 {
		if y%2 > 0 {
			result *= x
			result %= m
		}
		y /= 2
		x *= x
		x %= m
	}
	return result
}
