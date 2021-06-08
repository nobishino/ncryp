package ncryp

import (
	"math/rand"
	"time"
)

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

type RSAKey struct {
	Mod   uint64
	Power uint64
}

func (r RSAKey) Payload() Payload {
	return append(
		PayloadUint64(r.Mod),
		PayloadUint64(r.Power)...,
	)
}

func GenKeyPair() (RSAKey, RSAKey) {
	var primes []uint64 // [2^8, 2^10)の素数リストを作る
	for _, p := range Erathosthenes(1 << 10) {
		if p < 1<<8 {
			continue
		}
		primes = append(primes, p)
	}
	rand.Seed(time.Now().UnixNano())
	i, j := rand.Intn(len(primes)), rand.Intn(len(primes))
	p, q := primes[i], primes[j]
	prod := p * q
	// 二つの素数p, qからL = LCM(p-1,q-1)を作る(カーマイケルの定理)
	l := Lcm(p-1, q-1)
	// Lと互いに素な正の整数eを適当に取る
	e := rand.Uint64()%(l-2) + 2 // e is in [2,l-1]
	for Gcd(l, e) != 1 {
		e = rand.Uint64()%(l-2) + 2
	}
	// EulerPhi(L)を使ってeのLを法とする逆元を求める
	d := ModPow(e, EulerPhi(l)-1, l)
	// (prod, e)と(prod, d)を鍵ペアにする
	return RSAKey{
			Mod:   prod,
			Power: e,
		}, RSAKey{
			Mod:   prod,
			Power: d,
		}
}

func PrimeFactor(n uint64) map[uint64]uint64 {
	result := make(map[uint64]uint64)
	min := uint64(3)
	upper := CalcUpperSqrt(n)
OUTER:
	for n > 1 {
		// fmt.Println(n, result)
		if n%2 == 0 {
			result[2]++
			n /= 2
			continue OUTER
		}
		for d := min; d <= upper; d += 2 {
			if n%d == 0 {
				result[d]++
				n /= d
				continue OUTER
			}
			min = d + 2
		}
		result[n]++
		break
	}
	return result
}

// xのsqrtより大きくxより小さい数を返します
func CalcUpperSqrt(x uint64) uint64 {
	var i uint64
	for ; x > 0; x >>= 1 {
		i++
	}
	// original x is smaller than 1<<(i-1)
	return 1 << (i/2 + i%2)
}

func Gcd(x, y uint64) uint64 {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func Lcm(x, y uint64) uint64 {
	g := Gcd(x, y)
	x /= g
	y /= g
	return x * y * g
}

// xを法とする逆元を探すのに使うのでオイラーφ関数を定義する
func EulerPhi(x uint64) uint64 {
	var result uint64 = 1
	for p, exponent := range PrimeFactor(x) {
		result *= IntPow(p, exponent-1) * (p - 1)
	}
	return result
}

func IntPow(x, y uint64) uint64 {
	result := uint64(1)
	for y > 0 {
		if y%2 == 1 {
			result *= x
		}
		y /= 2
		x *= x
	}
	return result
}
