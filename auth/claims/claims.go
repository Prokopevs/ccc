package claims

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"os"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func EncryptSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			c.Set("encryptedData", nil)
			log.Fatalf("Error loading .env file")
			return
		}

		key := []byte(os.Getenv("KEY"))
		iv := []byte(os.Getenv("IV"))

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
