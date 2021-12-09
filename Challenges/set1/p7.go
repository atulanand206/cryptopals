package set1

import (
	"crypto/aes"
	"encoding/base64"
)

func DecryptECBStrings(encoded string, key string) (res string, err error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}
	keyBytes := []byte(key)
	decryptedBytes, err := DecryptECB(encryptedBytes, keyBytes)
	if err != nil {
		return
	}
	return string(decryptedBytes), nil
}

func DecryptECB(encryptedBytes, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(encryptedBytes))
	for i, j := 0, 16; i < len(encryptedBytes); i, j = i+16, j+16 {
		cipher.Decrypt(plainText[i:j], encryptedBytes[i:j])
	}
	return plainText, nil
}
