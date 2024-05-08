package claims

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"

	"fmt"
)

func PKCS5UnPadding(src []byte) []byte { 
    length := len(src) 
    unpadding := int(src[length-1]) 
    return src[:(length - unpadding)] 
}

func EncryptSignature(key string, iv string, signature string) ([]byte, error) {

	keyByte := []byte(key)
	ivByte := []byte(iv)

	decodedData, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return nil, err
	}

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		fmt.Println("Error creating cipher block:", err)
		return nil, err
	}

	if len(decodedData) < aes.BlockSize {
		fmt.Println("Decoded data is too short")
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, ivByte)
	mode.CryptBlocks(decodedData, decodedData)

	cutTrailingSpaces := []byte(strings.TrimSpace(string(decodedData)))

	return PKCS5UnPadding(cutTrailingSpaces), nil
}
