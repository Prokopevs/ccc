package server

import (
	"net/http"

	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/gin-gonic/gin"
)

type GameResponse struct {
	OwnerId    int `db:"ownerid" json:"ownerId"`
	Score      int `db:"score" json:"score"`
	GasStorage int `db:"gasstorage" json:"gasStorage"`
	GasMining  int `db:"gasmining" json:"gasMining"`
	Protection int `db:"protection" json:"protection"`
}

func (u *GameResponse) writeJSON(c *gin.Context) {
	writeJSONResponse(c, http.StatusOK, u)
}
func convertCoreGameToResponse(u *model.Game) *GameResponse {
	return &GameResponse{
		OwnerId:    u.OwnerId,
		Score:      u.Score,
		GasStorage: u.GasStorage,
		GasMining:  u.GasMining,
		Protection: u.Protection,
	}
}

type PriceResponse struct {
	Data map[int]int
}

func (u *PriceResponse) writeJSON(c *gin.Context) {
	writeJSONResponse(c, http.StatusOK, u)
}
func convertPriceToResponse(u map[int]int) *PriceResponse {
	return &PriceResponse{
		Data: u,
	}
}
