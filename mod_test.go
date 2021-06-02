package ncryp_test

import (
	"testing"

	"github.com/nobishino/ncryp"
)

func TestModPow(t *testing.T) {
	testcases := [...]struct {
		title  string
		x      uint64
		y      uint64
		mod    uint64
		expect uint64
	}{
		{
			title:  "#1",
			x:      3,
			y:      3,
			mod:    100,
			expect: 27,
		},
		{
			title:  "#2",
			x:      3,
			y:      3,
			mod:    10,
			expect: 7,
		},
		{
			title:  "#3",
			x:      4,
			y:      4,
			mod:    5,
			expect: 1,
		},
		{
			title:  "mod = 221",
			x:      2,
			y:      5 * 29,
			mod:    221,
			expect: 2,
		},
		{
			title:  "mod = 221",
			x:      3,
			y:      5 * 29,
			mod:    221,
			expect: 3,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.ModPow(tt.x, tt.y, tt.mod)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}
