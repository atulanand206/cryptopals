package set1

import (
	"fmt"
	"strconv"
)

func FixedXor(s1 string, s2 string) (res string, err error) {
	s1B, err := strconv.ParseInt(s1, 16, 16)
	if err != nil {
		return
	}
	s2B, err := strconv.ParseInt(s2, 16, 16)
	if err != nil {
		return
	}
	fmt.Println(s1B)
	fmt.Println(s2B)
	return "hex.EncodeToString(k)", nil
}
