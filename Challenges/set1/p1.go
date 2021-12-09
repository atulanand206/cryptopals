package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexS string) string {
	p, err := hex.DecodeString(hexS)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(p)
}
