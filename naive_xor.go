package ncryp

import (
	"strconv"
)

// 16bit(2bytes)の共有鍵を使って平文に対して繰り返し単純にXORを適用するナイーブな共通鍵暗号を実装する

type Key16 [2]byte

type Payload []byte

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
