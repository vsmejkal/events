package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/vsmejkal/events/model"
)

func SourceList(c *gin.Context) {
	data, err := model.Create("event", "id");
	data, err := model.Get("event", id);
	data, err := model.Update("event", id);
	data, err := model.Delete("event", id);

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Seznam zdrojů",
		"view": "sources.index"
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
