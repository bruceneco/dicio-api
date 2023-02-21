package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct{}

func NewPingController() PingController {
	return PingController{}
}

func (pc *PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
