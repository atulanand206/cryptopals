package set1

import (
	"encoding/hex"
)

func RepeatingKeyXOR(s1 string, key string) string {
	s1Bytes := []byte(s1)
	s2 := []byte{}
	for i := 0; i < len(s1Bytes); i++ {
		s2 = append(s2, (s1Bytes[i] ^ key[i%len(key)]))
	}
	x := hex.EncodeToString(s2)
	return x
}
