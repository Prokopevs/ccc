package server

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/gin-gonic/gin"
)

const (
	codeNoParam   = "NO_PARAM"
	codeEmptyBody = "NO_BODY"
)

type response interface {
	writeJSON(*gin.Context)
}

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
		h.log.Errorw("Get game info", "err", err)
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	return convertCoreGameToResponse(gameInfo)
}

func (h *HTTP) updateScore(c *gin.Context) {
	resp := h.updateScoreResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) updateScoreResponse(r *gin.Context) response {
	var s model.Score
	if err := r.ShouldBindJSON(&s); err != nil {
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	if reflect.DeepEqual(s, model.Score{}) {
		return getBadRequestWithMsgResponse("no payload", codeEmptyBody)
	}

	code, err := h.service.UpdateScore(r.Request.Context(), &s)
	if err != nil {
		h.log.Errorw("Update score info", "err", err)
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	return newOKResponse(code)
}

func (h *HTTP) updateMultiplicator(c *gin.Context) {
	resp := h.updateMultiplicatorResponse(c)

	resp.writeJSON(c)
}

func (h *HTTP) updateMultiplicatorResponse(r *gin.Context) response {
	var m model.MultipUpdate
	if err := r.ShouldBindJSON(&m); err != nil {
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	if reflect.DeepEqual(m, model.Score{}) {
		return getBadRequestWithMsgResponse("no payload", codeEmptyBody)
	}

	code, err := h.service.UpdateMultiplicator(r.Request.Context(), &m)
	if err != nil {
		if errors.Is(err, core.ErrNoSuchMultiplicator) {
			return getBadRequestWithMsgResponse(err.Error(), code)
		}
		
		h.log.Errorw("Update score info", "err", err)
		return getInternalServerErrorResponse("internal error", core.CodeInternal)
	}

	return newOKResponse(code)
}