package ncryp_test

import (
	"testing"

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
