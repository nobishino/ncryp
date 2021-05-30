package ncryp

// 16bit(2bytes)の共有鍵を使って平文に対して繰り返し単純にXORを適用するナイーブな共通鍵暗号を実装する

type Key16 [2]byte

func NaiveSymCryp(data []byte, key Key16) []byte {
	result := make([]byte, len(data))
	for i, byt := range data {
		result[i] = byt ^ key[i%2]
	}
	return result
}
