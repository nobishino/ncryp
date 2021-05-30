package ncryp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nobishino/ncryp"
)

func TestNaiveSymCryp(t *testing.T) {
	testcases := [...]struct {
		title  string
		data   ncryp.Payload
		key    ncryp.Key16
		expect ncryp.Payload
	}{
		{
			data:   []byte{0x0, 0x0, 0x0, 0x0},
			key:    [2]byte{0xF, 0x0},
			expect: []byte{0xF, 0x0, 0xF, 0x0},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.NaiveSymCryp(tt.data, tt.key)
			if diff := cmp.Diff(got, tt.expect); diff != "" {
				t.Errorf("expect %v, but got %v. diff:\n%s", tt.expect, got, diff)
			}
		})
	}
}
