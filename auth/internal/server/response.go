package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

)

func writeJSONResponse(c *gin.Context, status int, v interface{}) error {
	bb, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	c.Writer.Header().Set("content-type", "application/json")

	c.JSON(status, bb)

	return err
}
