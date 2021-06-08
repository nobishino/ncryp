package ncryp_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nobishino/ncryp"
)

func TestIsPrime(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     uint64
		expect bool
	}{
		{
			in:     2,
			expect: true,
		},
		{
			in:     3,
			expect: true,
		},
		{
			in:     4,
			expect: false,
		},
		{
			in:     53,
			expect: true,
		},
		{
			in:     217,
			expect: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.IsPrime(tt.in)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}
func TestErathosthenes(t *testing.T) {
	testcases := [...]struct {
		title  string
		max    uint64
		expect []uint64
	}{
		{
			max:    2,
			expect: []uint64{2},
		},
		{
			max:    11,
			expect: []uint64{2, 3, 5, 7, 11},
		},
		{
			max: 100,
			expect: []uint64{
				2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
				31, 37, 41, 43, 47, 53, 59, 61, 67,
				71, 73, 79, 83, 89, 97,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.Erathosthenes(tt.max)
			if diff := cmp.Diff(got, tt.expect); diff != "" {
				t.Errorf("expect %v, but got %v. diff:\n%s", tt.expect, got, diff)
			}
		})
	}
}

func TestGenKeyPair(t *testing.T) {
	testcases := [...]struct {
		title   string
		in      string
		expect1 ncryp.RSAKey
		expect2 ncryp.RSAKey
	}{}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			ncryp.GenKeyPair()
			// if got != tt.expect {
			// 	t.Errorf("expect %v, but got %v", tt.expect, got)
			// }
		})
	}
}

func TestPrimeFactor(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     uint64
		expect map[uint64]uint64
	}{
		{
			in: 2,
			expect: map[uint64]uint64{
				2: 1,
			},
		},
		{
			in: 4,
			expect: map[uint64]uint64{
				2: 2,
			},
		},
		{
			in: 2340,
			expect: map[uint64]uint64{
				2:  2,
				3:  2,
				5:  1,
				13: 1,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.PrimeFactor(tt.in)
			if diff := cmp.Diff(got, tt.expect, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("expect %v, but got %v. diff:\n%s", tt.expect, got, diff)
			}
		})
	}
}

func TestCalcUpperSqrt(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     uint64
		expect uint64
	}{
		{
			in:     0b11, // 3
			expect: 0b10, // 2 which is larger than sqrt(3)
		},
		{
			in:     16,
			expect: 8, // which is larger than sqrt(16)
		},
		{
			in:     17,
			expect: 8, // which is larger than sqrt(17)
		},
		{
			in:     1000,
			expect: 32, // which is larger than sqrt(1000)
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.CalcUpperSqrt(tt.in)
			if got != tt.expect {
				t.Errorf("expect %b, but got %b", tt.expect, got)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	testcases := [...]struct {
		title  string
		x      uint64
		y      uint64
		expect uint64
	}{
		{
			x:      18,
			y:      10,
			expect: 2,
		},
		{
			x:      351,
			y:      36,
			expect: 9,
		},
		{
			x:      0,
			y:      36,
			expect: 36,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.Gcd(tt.x, tt.y)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	testcases := [...]struct {
		title  string
		x      uint64
		y      uint64
		expect uint64
	}{
		{
			x:      8,
			y:      12,
			expect: 24,
		},
		{
			x:      80,
			y:      12,
			expect: 240,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.Lcm(tt.x, tt.y)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}

func TestEulerPhi(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     uint64
		expect uint64
	}{
		{
			in:     5,
			expect: 4,
		},
		{
			in:     12,
			expect: 4, // 1,5,7,11
		},
		{
			in:     15,
			expect: 8, // 1,2,4,7,8,11,13,14
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.EulerPhi(tt.in)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}

func TestIntPow(t *testing.T) {
	testcases := [...]struct {
		title  string
		x      uint64
		y      uint64
		expect uint64
	}{
		{
			x:      2,
			y:      2,
			expect: 4,
		},
		{
			x:      3,
			y:      4,
			expect: 81,
		},
		{
			x:      0,
			y:      0,
			expect: 1,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.IntPow(tt.x, tt.y)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}

func TestRSAKeyGen(t *testing.T) {
	k1, k2 := ncryp.GenKeyPair()
	rand.Seed(1)
	plaintext := rand.Uint64() % k1.Mod
	fmt.Println(k1)

	cipher := ncryp.ModPow(plaintext, k1.Power, k1.Mod)

	decipher := ncryp.ModPow(cipher, k2.Power, k2.Mod)

	if decipher != plaintext {
		t.Errorf("expect %d, but got %d", plaintext, decipher)
	}
}

func TestRSAKeyPayload(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     ncryp.RSAKey
		expect ncryp.Payload
	}{
		{
			in: ncryp.RSAKey{
				Mod:   232613,
				Power: 20329,
			},
			expect: []byte{
				0xA5,
				0x8C,
				0x03,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x69,
				0x4F,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
				0x00,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.in.Payload()
			if diff := cmp.Diff(got, tt.expect); diff != "" {
				t.Errorf("expect %v, but got %v. diff:\n%s", tt.expect, got, diff)
			}
		})
	}
}
