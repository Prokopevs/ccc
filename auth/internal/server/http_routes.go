package server

import (
	"github.com/gin-gonic/gin"
)

func (h *HTTP) setRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/auth") 
	{
		api.GET("/me", h.me)
	}
}