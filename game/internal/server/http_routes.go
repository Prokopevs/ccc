package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/Prokopevs/ccc/game/docs"
)

func (h *HTTP) setRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/game") 
	{
		api.GET("/getGame/:id", h.getGame)
		api.POST("/updateScore", h.updateScore)
		api.POST("/updateMultiplicator", h.updateMultiplicator)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}