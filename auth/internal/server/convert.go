package server

import (
	"net/http"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
)

type userInfoResponse struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Username  string `json:"username"`
}

func (u *userInfoResponse) writeJSON(c *gin.Context) {
	writeJSONResponse(c, http.StatusOK, u)
}
func convertCoreUserInfoToResponse(u *core.UserInfo) *userInfoResponse {
	return &userInfoResponse{
		Id:        u.Id,
		Firstname: u.Firstname,
		Username:  u.Username,
	}
}
