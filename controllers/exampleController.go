package controllers

import (
	"github.com/gin-gonic/gin"
)

type ExampleController struct{}

func (s *ExampleController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
