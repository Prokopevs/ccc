package claims

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
)

func EncryptSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		envKey, ok := os.LookupEnv("KEY")
		if !ok {
			panic("provede envKey")
		}
		ivKey, ok := os.LookupEnv("IV")
		if !ok {
			panic("provede ivKey")
		}

		key := []byte(envKey)
		iv := []byte(ivKey)

		signature := c.Request.Header.Get("signature")

		decodedData, err := base64.StdEncoding.DecodeString(signature)
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
