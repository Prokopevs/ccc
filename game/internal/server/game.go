package server

import (
	"errors"
	"strconv"

	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/Prokopevs/ccc/game/internal/pg"
	"github.com/gin-gonic/gin"
)

const (
	codeNoParam   = "NO_PARAM"
	codeEmptyBody = "NO_BODY"
)

type response interface {
	writeJSON(*gin.Context)
}

// @Summary  	 Get game data
// @Tags 		 Game
// @Description  Get game data
// @Accept 	 	 json
// @Produce 	 json
// @Param 		 id path int true "user Id"
// @Success 	 200  {object}  model.Game
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/game/getGame/{id} [get]“
func (h *HTTP) getGame(c *gin.Context) {
	resp := h.getGameResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) getGameResponse(r *gin.Context) response {
	id := r.Param("id")
	if id == "" {
		return getBadRequestWithMsgResponse("no id provided", codeNoParam)
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	gameInfo, err := h.service.GetGame(r.Request.Context(), idInt)
	if err != nil {
		if errors.Is(err, core.ErrNoSuchUser) {
			return getBadRequestWithMsgResponse(err.Error(), core.CodeBadRequest)
		}

		h.log.Errorw("Get game info", "err", err)
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	return convertCoreGameToResponse(gameInfo)
}

// @Summary  	 Update score
// @Tags 		 Game
// @Description  Update score
// @Accept 	 	 json
// @Produce 	 json
// @Param		 message	body    model.Score	true	"Body"
// @Success 	 200  {object}  OKStruct
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/game/updateScore [post]“
func (h *HTTP) updateScore(c *gin.Context) {
	resp := h.updateScoreResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) updateScoreResponse(r *gin.Context) response {
	var s model.Score
	if err := r.ShouldBindJSON(&s); err != nil {
		return getBadRequestWithMsgResponse("no payload", codeEmptyBody)
	}

	code, err := h.service.UpdateScore(r.Request.Context(), &s)
	if err != nil {
		if errors.Is(err, core.ErrNoSuchUser) {
			return getBadRequestWithMsgResponse(err.Error(), core.CodeBadRequest)
		}
		
		h.log.Errorw("Update score info", "err", err)
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	return newOKResponse(code)
}

// @Summary  	 Update multiplicator
// @Tags 		 Game
// @Description  Update multiplicator
// @Accept 	 	 json
// @Produce 	 json
// @Param		 message	body    model.MultipUpdate	true	"Body"
// @Success 	 200  {object}  OKStruct
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/v1/game/updateMultiplicator [post]“
func (h *HTTP) updateMultiplicator(c *gin.Context) {
	resp := h.updateMultiplicatorResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) updateMultiplicatorResponse(r *gin.Context) response {
	var m model.MultipUpdate
	if err := r.ShouldBindJSON(&m); err != nil {
		return getBadRequestWithMsgResponse("no payload", codeEmptyBody)
	}

	code, err := h.service.UpdateMultiplicator(r.Request.Context(), &m)
	if err != nil {
		if errors.Is(err, core.ErrNoSuchMultiplicator) {
			return getBadRequestWithMsgResponse(err.Error(), code)
		}
		if errors.Is(err, pg.ErrNoEnoughScore) || errors.Is(err, pg.ErrMaxLevel) {
			return getBadRequestWithMsgResponse(err.Error(), code)
		}
		if errors.Is(err, core.ErrNoSuchUser) {
			return getBadRequestWithMsgResponse(err.Error(), core.CodeBadRequest)
		}
		
		h.log.Errorw("Update multiplicator info", "err", err)
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	return newOKResponse(code)
}