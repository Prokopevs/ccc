package server

import (
	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
)

const (
	codeNoHeader = "NO_HEADER"
)

type response interface {
	writeJSON(*gin.Context)
}

func (h *HTTP) me(c *gin.Context) {
	resp := h.getMeResponse(c)

	resp.writeJSON(c)
}

// @Summary  	 Get user data
// @Tags 		 Auth
// @Description  Get user data
// @Accept 	 	 json
// @Produce 	 json
// @Param		 initData	header	string	true	"InitData header"
// @Success 	 200  {object} core.UserInfo
// @Failure      401  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/auth/me [get]``
func (h *HTTP) getMeResponse(r *gin.Context) response {
	initData := r.Request.Header.Get("initData")
	if initData == "" {
		return getUnauthorizedErrorWithMsgResponse("no initData", codeNoHeader)
	}

	userInfo, code, err := h.service.GetUserInfo(r.Request.Context(), initData)
	if err != nil {
		if code == core.CodeInternal {
			h.log.Errorw("Get user info.", "err", err)
			return getInternalServerErrorResponse("internal error", code)
		}

		return getUnauthorizedErrorWithMsgResponse(err.Error(), code)
	}

	return convertCoreUserInfoToResponse(userInfo)
}
// @Param		 message	body    core.UserInfo	true	"Account Info"
