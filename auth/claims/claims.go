package claims

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
)

type EncryptImpl struct {
	key string
	iv  string
}

func NewEncrypt(key string, iv string) *EncryptImpl {
	return &EncryptImpl{
		key: key,
		iv:  iv,
	}
}

func (s *EncryptImpl) EncryptSignature(initData string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := []byte(s.key)
		iv := []byte(s.iv)

		decodedData, err := base64.StdEncoding.DecodeString(initData)
		if err != nil {
			fmt.Println("Error decoding data:", err)
			c.Set("encryptedData", nil)
		}

		block, err := aes.NewCipher(key)
		if err != nil {
			fmt.Println("Error creating cipher block:", err)
			c.Set("encryptedData", nil)
		}

		if len(decodedData) < aes.BlockSize {
			fmt.Println("Decoded data is too short")
			c.Set("encryptedData", nil)
		}

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(decodedData, decodedData)

		cutTrailingSpaces := []byte(strings.TrimSpace(string(decodedData)))

		c.Set("encryptedData", cutTrailingSpaces)

		c.Next()
	}
}
