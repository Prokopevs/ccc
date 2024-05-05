package server

import (
	"net/http"

	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/gin-gonic/gin"
)

type GameResponse struct {
	OwnerId    int    `db:"ownerId"`
	Score      string `db:"score"`
	GasStorage int    `db:"gasStorage"`
	GasMining  string `db:"gasMining"`
	Protection int    `db:"protection"`
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
