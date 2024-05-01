package server

import (
	"fmt"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
)

const (
	codeNoHeader = "NO_HEADER"
)

type response interface {
	writeJSON(*gin.Context) error
}

func (h *HTTP) me(c *gin.Context) {
	resp := h.getMeResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) getMeResponse(r *gin.Context) response {
	initData := r.Request.Header.Get("initData")
	if initData == "" {
		return getUnauthorizedErrorWithMsgResponse("no initData", codeNoHeader)
	}
	fmt.Println(initData)

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
