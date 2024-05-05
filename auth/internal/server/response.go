package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type okResponse struct {
	jsonData interface{}
}

func newOKResponse(data interface{}) response {
	return &okResponse{
		jsonData: data,
	}
}

func (o *okResponse) writeJSON(c *gin.Context) {
	writeJSONResponse(c, http.StatusOK, o.jsonData)
}

func writeJSONResponse(c *gin.Context, status int, v interface{})  {
	c.JSON(status, v)
}
