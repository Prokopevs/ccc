package server

import (
	"net/http"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
)

type okResponse struct {
	Data interface{}
}

func newOKResponse(data []*core.UserReferrals) response {
	if len(data) == 0 {
        data = []*core.UserReferrals{}
    }
	return &okResponse{
		Data: data,
	}
}

func newOKUsersResponse(data []*core.User) response {
	if len(data) == 0 {
        data = []*core.User{}
    }
	return &okResponse{
		Data: data,
	}
}

func (o *okResponse) writeJSON(c *gin.Context) {
	writeJSONResponse(c, http.StatusOK, o)
}

func writeJSONResponse(c *gin.Context, status int, v interface{})  {
	c.JSON(status, v)
}
