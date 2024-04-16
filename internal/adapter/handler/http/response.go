package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Global response
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// Success returns a simple payload for success
func Success(ctx *gin.Context, msg string) {
	res := Response{
		Success: true,
		Message: msg,
	}

	ctx.JSON(http.StatusOK, res)
}

// Failure returns a simple payload for failed
func Failure(ctx *gin.Context, msg string) {
	res := Response{
		Success: false,
		Message: msg,
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
}
