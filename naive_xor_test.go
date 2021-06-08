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

func TestNewPayload(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     string
		expect ncryp.Payload
		isErr  bool
	}{
		{
			in:     "FE",
			expect: []byte{0xFE},
		},
		{
			in:     "FE12",
			expect: []byte{0xFE, 0x12},
		},
		{
			in:    "FG",
			isErr: true,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got, err := ncryp.NewPayload(tt.in)
			switch {
			case !tt.isErr && err != nil:
				t.Errorf("expect err to be nil but got %v", err)
			case tt.isErr && err == nil:
				t.Error("expect non-nil error but got nil")
			}
			if diff := cmp.Diff(got, tt.expect); diff != "" {
				t.Errorf("expect %v, but got %v. diff:\n%s", tt.expect, got, diff)
			}
		})
	}
}

func TestPayloadString(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     ncryp.Payload
		expect string
	}{
		{
			in:     []byte{0xFE},
			expect: "FE",
		},
		{
			in:     []byte{0xFE, 0x12},
			expect: "FE12",
		},
		{
			in:     []byte{0xFE, 0x12, 0, 0, 0, 0, 0, 0},
			expect: "FE12000000000000",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.in.String()
			if diff := cmp.Diff(got, tt.expect); diff != "" {
				t.Errorf("expect %v, but got %v. diff:\n%s", tt.expect, got, diff)
			}
		})
	}
}

func TestKey16ToUint64(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     ncryp.Key16
		expect uint64
	}{
		{
			in:     [2]byte{0xF, 0x0},
			expect: 15,
		},
		{
			in:     [2]byte{0xFF, 0xFF},
			expect: 65535,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.in.Uint64()
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}

func TestKey16FromUint64(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     uint64
		expect ncryp.Key16
	}{
		{
			in:     15,
			expect: [2]byte{0xF, 0x0},
		},
		{
			in:     65535,
			expect: [2]byte{0xFF, 0xFF},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := ncryp.Key16FromUint64(tt.in)
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}

func TestPayloadUint64(t *testing.T) {
	testcases := [...]struct {
		title  string
		in     ncryp.Payload
		expect uint64
	}{
		{
			in:     []byte{0xFF, 0, 0xFF},
			expect: 255 + 255*(1<<16),
		},
		{
			in:     []byte{0x0F},
			expect: 15,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.in.Uint64()
			if got != tt.expect {
				t.Errorf("expect %v, but got %v", tt.expect, got)
			}
		})
	}
}
