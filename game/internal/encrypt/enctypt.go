package encrypt

import (
	"github.com/Prokopevs/ccc/auth/claims"
	"github.com/gin-gonic/gin"
)

func Encrypt(key, iv string) gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.Request.Header.Get("signature")

		encryptedData, err := claims.EncryptSignature(key, iv, signature)
		if err != nil {
			c.Next()
		}

		c.Set("encryptedData", encryptedData)
		c.Next()
	}
}