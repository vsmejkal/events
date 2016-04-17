package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SourceList(c *gin.Context) {
	c.HTML(http.StatusOK, "sources.list", gin.H{
		"test": "ABCD",
	})
}

func SourceCreate(c *gin.Context) {

}

func SourceRead(c *gin.Context) {

}

func SourceUpdate(c *gin.Context) {

}

func SourceDelete(c *gin.Context) {

}
