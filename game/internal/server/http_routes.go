package server

import (
	_ "github.com/Prokopevs/ccc/game/docs"
	"github.com/Prokopevs/ccc/game/internal/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (h *HTTP) setRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/game") 
	{
		api.GET("/getGame/:id", h.getGame)
		api.POST("/updateScore", encrypt.Encrypt(h.key, h.iv), h.updateScore)
		api.POST("/updateMultiplicator", encrypt.Encrypt(h.key, h.iv), h.updateMultiplicator)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
