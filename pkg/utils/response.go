package utils

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
