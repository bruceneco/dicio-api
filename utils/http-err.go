package utils

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code    int    `json:"code,omitempty" example:"400"`
	Message string `json:"message,omitempty" example:"status bad request"`
}

func NewError(c *gin.Context, status int, message string) {
	er := HTTPError{
		Code:    status,
		Message: message,
	}
	c.JSON(status, er)
}
