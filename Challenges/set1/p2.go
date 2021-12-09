package set1

import (
	"encoding/hex"
)

func FixedXor(s1 string, s2 string) (res string, err error) {
	s1B, err := hex.DecodeString(s1)
	if err != nil {
		return
	}
	s2B, err := hex.DecodeString(s2)
	if err != nil {
		return
	}
	resB := make([]byte, len(s1B))
	for i := 0; i < len(s1B); i++ {
		resB[i] = s1B[i] ^ s2B[i]
	}
	res = hex.EncodeToString(resB)
	return res, nil
}
