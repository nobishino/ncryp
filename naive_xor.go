package ncryp

import (
	"fmt"
	"strconv"
	"strings"
)

// 16bit(2bytes)の共有鍵を使って平文に対して繰り返し単純にXORを適用するナイーブな共通鍵暗号を実装する

type Key16 [2]byte

func (k Key16) Uint64() uint64 {
	return uint64(k[1])<<8 + uint64(k[0])
}

func Key16FromUint64(x uint64) Key16 {
	var k Key16
	k[0] = byte(x)
	k[1] = byte(x >> 8)
	return k
}

type Payload []byte

func PayloadUint64(x uint64) Payload {
	return []byte{
		byte(x),
		byte(x >> 8),
		byte(x >> 16),
		byte(x >> 24),
		byte(x >> 32),
		byte(x >> 40),
		byte(x >> 48),
		byte(x >> 56),
	}
}

func (p Payload) Uint64() uint64 {
	var result uint64
	for i := 0; i < 8 && i < len(p); i++ {
		result += uint64(p[i]) << (8 * i)
	}
	return result
}

func NaiveSymCryp(data Payload, key Key16) Payload {
	result := make([]byte, len(data))
	for i, byt := range data {
		result[i] = byt ^ key[i%2]
	}
	return result
}

func NewPayload(s string) (Payload, error) {
	var result Payload
	for i := 0; i < len(s); i += 2 {
		upper := i + 2
		if upper > len(s) {
			upper = len(s)
		}
		in := s[i:upper]
		i, err := strconv.ParseInt(in, 16, 16)
		if err != nil {
			return nil, err
		}
		result = append(result, byte(i))
	}
	return result, nil
}

func (p Payload) String() string {
	var ss []string
	for _, byt := range p {
		ss = append(ss, fmt.Sprintf("%02X", byt))
	}
	return strings.Join(ss, "")
}

func is(r rune) bool {
	switch {
	case '0' <= r && r <= '9':
		return true
	case 'A' <= r && r <= 'F':
		return true
	case 'A' <= r && r <= 'F':
		return true
	default:
		return false
	}
}
