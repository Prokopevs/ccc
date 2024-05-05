package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/Prokopevs/ccc/auth/docs"
)

func (h *HTTP) setRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/auth") 
	{
		api.GET("/me", h.me)
		api.GET("/referrals/:id", h.referrals)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}