package server

import (
	"net/http"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status    int       `json:"-"`
	Code      core.Code `json:"code"`
	ErrorInfo string    `json:"errorInfo"`
}

func (e errorResponse) writeJSON(c *gin.Context) error {
	return writeJSONResponse(c, e.Status, e)
}

func getInternalServerErrorResponse(msg string, code core.Code) errorResponse {
	return errorResponse{
		Status:    http.StatusInternalServerError,
		Code:      code,
		ErrorInfo: msg,
	}
}

func getUnauthorizedErrorWithMsgResponse(msg string, code core.Code) errorResponse {
	return errorResponse{
		Status:    http.StatusUnauthorized,
		Code:      code,
		ErrorInfo: msg,
	}
}