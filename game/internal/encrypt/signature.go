package encrypt

import (
	"encoding/json"
	"math"
	"time"

	"github.com/Prokopevs/ccc/game/internal/model"
)

type CombinedScoreData struct {
	Timestamp int64
	Data      model.Score
}

type CombinedMultiplicatorData struct {
	Timestamp int64
	Data      model.MultipUpdate
}

func CheckScoreData(encryptedData []byte, s model.Score) bool {
	var result CombinedScoreData

	err := json.Unmarshal(encryptedData, &result)
	if err != nil {
		return false
	}

	currentTime := time.Now().Unix() 
	timestamp := currentTime * 1000
	diff := math.Abs(float64(timestamp - result.Timestamp))
	if diff > 60000 {
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

	currentTime := time.Now().Unix() 
	timestamp := currentTime * 1000
	diff := math.Abs(float64(timestamp - result.Timestamp))
	if diff > 55000 {
		return false
	}

	if result.Data != m {
		return false
	}

	return true
}
