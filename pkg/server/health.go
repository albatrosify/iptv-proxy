package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthHandler returns a simple 200 OK status
func (c *Config) healthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
