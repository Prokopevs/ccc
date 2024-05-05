package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeJSONResponse(c *gin.Context, status int, v interface{})  {
	c.JSON(status, v)
}


type OKStruct struct {
	jsonData    interface{}
}
func newOKResponse(data interface{}) *OKStruct {
	return &OKStruct{
		jsonData: data,
	}
}

func (u *OKStruct) writeJSON(c *gin.Context) {
	writeJSONResponse(c, http.StatusOK, u.jsonData)
}

