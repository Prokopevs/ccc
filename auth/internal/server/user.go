package server

import (
	"strconv"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
)

const (
	codeNoHeader = "NO_HEADER"
	codeNoParam  = "NO_PARAM"
)

type response interface {
	writeJSON(*gin.Context)
}

// @Summary  	 Get user data
// @Tags 		 Auth
// @Description  Get user data
// @Accept 	 	 json
// @Produce 	 json
// @Param		 initData	header	string	true   "InitData header"
// @Param        inviterId  query   int     false  "get param inviterId"
// @Success 	 200  {object}  userInfoResponse
// @Failure      401  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/auth/me [get]“
func (h *HTTP) Me(c *gin.Context) {
	resp := h.getMeResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) getMeResponse(r *gin.Context) response {
	initData := r.Request.Header.Get("initData")
	if initData == "" {
		return getUnauthorizedErrorWithMsgResponse("no initData", codeNoHeader)
	}

	id, ok := r.GetQuery("inviterId")
	idInt := 0
	var err error
	if ok && id != "" {
		idInt, err = strconv.Atoi(id)
		if err != nil {
			return getBadRequestWithMsgResponse("wrong inviterId provided", core.CodeBadRequest)
		}
	}

	userInfo, code, err := h.service.GetUserInfo(r.Request.Context(), initData, idInt)
	if err != nil {
		if code == core.CodeInternal {
			h.log.Errorw("Get user info.", "err", err)
			return getInternalServerErrorResponse("internal error", code)
		}

		return getUnauthorizedErrorWithMsgResponse(err.Error(), code)
	}

	return convertCoreUserInfoToResponse(userInfo)
}


// @Summary  	 Get user referrals
// @Tags 		 Auth
// @Description  Get user referrals
// @Accept 	 	 json
// @Produce 	 json
// @Param 		 id path int true "Inviter Id"
// @Success 	 200  {array}   core.UserReferrals
// @Failure      400  {object}  errorResponse
// @Failure      401  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/auth/referrals/{id} [get]“
func (h *HTTP) Referrals(c *gin.Context) {
	resp := h.getReferralResponse(c)

	resp.writeJSON(c)
}
func (h *HTTP) getReferralResponse(r *gin.Context) response {
	id := r.Param("id")
	if id == "" {
		return getBadRequestWithMsgResponse("no param", codeNoParam)
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return getBadRequestWithMsgResponse("wrong id provided", core.CodeBadRequest)
	}

	referrals, code, err := h.service.GetUserReferrals(r.Request.Context(), idInt)
	if err != nil {
		if code == core.CodeInvalidUserID {
			return getBadRequestWithMsgResponse("invalid user id", code)
		}
		h.log.Errorw("Get referral response.", "err", err)
		return getInternalServerErrorResponse("internal error", code)
	}

	return newOKResponse(referrals)
}


// @Summary  	 Get user referrals
// @Tags 		 Auth
// @Description  Get user referrals
// @Accept 	 	 json
// @Produce 	 json
// @Param 		 id path int true "Inviter Id"
// @Success 	 200  {array}   core.UserReferrals
// @Failure      400  {object}  errorResponse
// @Failure      401  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/auth/referrals/{id} [get]“
func (h *HTTP) Users(c *gin.Context) {
	resp := h.getUsersResponse(c)

	resp.writeJSON(c)
}
func (h *HTTP) getUsersResponse(r *gin.Context) response {
	password := r.Request.Header.Get("password")
	if password == "" || password != h.password {
		return getUnauthorizedErrorWithMsgResponse("wrong header", codeNoHeader)
	}

	users, code, err := h.service.GetUsers(r.Request.Context())
	if err != nil {
		h.log.Errorw("Get users response.", "err", err)
		return getInternalServerErrorResponse("internal error", code)
	}

	return newOKUsersResponse(users)
}
