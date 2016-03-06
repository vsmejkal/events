package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vsmejkal/events/config"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	addr := fmt.Sprintf(":%d", config.Admin.Port)

	r.Run(addr)
}