package set1

import (
	"encoding/hex"
)

func LetterRatio(s []byte) (score float64) {
	total := 0.0
	for _, l := range s {
		if (l >= 97 && l <= 122) || l == 32 {
			total++
		}
	}
	return total / float64(len(s))
}

func SingleByteXor(s1 string) (result string, err error) {
	k := make([]int, 256)
	for i := 0; i < 256; i++ {
		k[i] = i
	}
	s1B, err := hex.DecodeString(s1)
	if err != nil {
		return
	}
	score := 0.0
	res := -1
	for j, l := range k {
		resB := SBX(s1B, l)
		ratio := LetterRatio(resB)
		if ratio > 0.8 && ratio > score {
			score = ratio
			res = j
		}
	}
	return string(SBX(s1B, res)), nil
}

func SBX(s1B []byte, l int) []byte {
	resB := make([]byte, len(s1B))
	for i := 0; i < len(s1B); i++ {
		resB[i] = s1B[i] ^ byte(l)
	}
	return resB
}
