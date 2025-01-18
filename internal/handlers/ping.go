package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mamcer/cookbook/internal/services"
)

type PingHandler struct {
	Service services.IPingService
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func (ph *PingHandler) Ping(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	c.JSON(200, gin.H{
		"message": ph.Service.GetMessage(),
	})
}

func NewPingHandler(g *gin.Engine, ps *services.PingService) {
	ph := PingHandler{ps}
	g.GET("/ping", ph.Ping)
	g.OPTIONS("/ping", preflight)
}
