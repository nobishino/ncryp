package ncryp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
