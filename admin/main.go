package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vsmejkal/events/config"
	"fmt"
	"net/http"
	"path"
	"os"
	"log"
)

func printHelp() {
	fmt.Printf("Usage: %s config.json\n", path.Base(os.Args[0]))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		printHelp()
		return
	}

	configFile := os.Args[1]
	if err := config.Load(configFile); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob(config.Admin.DocumentRoot + "/templates/*")
	router.Static("/assets", config.Admin.DocumentRoot + "/assets")

	router.GET("/source/:action/*id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	addr := fmt.Sprintf(":%d", config.Admin.Port)
	router.Run(addr)
}