package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) PongV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (p *PongController) HelloV1(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}

func (p *PongController) PongV2(c *gin.Context) {
	name := c.DefaultQuery("name", "tranduchuy")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
	})
}
