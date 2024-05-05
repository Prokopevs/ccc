package server

import (
	"net/http"

	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status    int       `json:"-"`
	Code      core.Code `json:"code"`
	ErrorInfo string    `json:"errorInfo"`
}

func (e errorResponse) writeJSON(c *gin.Context) {
	writeJSONResponse(c, e.Status, e)
}
func getBadRequestWithMsgResponse(msg string, code core.Code) errorResponse {
	return errorResponse{
		Status:    http.StatusBadRequest,
		Code:      code,
		ErrorInfo: msg,
	}
}

func getInternalServerErrorResponse(msg string, code core.Code) errorResponse {
	return errorResponse{
		Status:    http.StatusInternalServerError,
		Code:      code,
		ErrorInfo: msg,
	}
}