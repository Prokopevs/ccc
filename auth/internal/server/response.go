package server

import (
	"github.com/gin-gonic/gin"

)
func writeJSONResponse(c *gin.Context, status int, v interface{})  {
	c.JSON(status, v)
}
