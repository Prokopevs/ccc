package server

import (
	"encoding/json"

	"github.com/Prokopevs/ccc/auth/claims"
	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/gin-gonic/gin"
)

type CombinedScoreData struct {
	Timestamp int64 
	Data      model.Score 
}

type CombinedMultiplicatorData struct {
	Timestamp int64 
	Data      model.MultipUpdate 
}

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

func CheckScoreData(encryptedData []byte, s model.Score) bool {
	var result CombinedScoreData

	err := json.Unmarshal(encryptedData, &result)
	if err != nil {
		return false
	}

	if result.Data != s {
		return false
	}

	return true
}

func CheckMultiplicatorData(encryptedData []byte, m model.MultipUpdate) bool {
	var result CombinedMultiplicatorData

	err := json.Unmarshal(encryptedData, &result)
	if err != nil {
		return false
	}

	if result.Data != m {
		return false
	}

	return true
}
 

