package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vsmejkal/events/config"
	"fmt"
	"path"
	"os"
	"log"
	ctl "github.com/vsmejkal/events/admin/controllers"
)

func printHelp() {
	fmt.Printf("Usage: %s configFile\n", path.Base(os.Args[0]))
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
	router.LoadHTMLGlob(config.Admin.DocumentRoot + "/templates/**/*")
	router.Static("/assets", config.Admin.DocumentRoot + "/assets")

	router.GET("/sources", ctl.SourceList)
	router.POST("/sources", ctl.SourceCreate)
	router.GET("/sources/:id", ctl.SourceRead)
	router.PUT("/sources/:id", ctl.SourceUpdate)
	router.DELETE("/sources/:id", ctl.SourceDelete)

	addr := fmt.Sprintf(":%d", config.Admin.Port)
	router.Run(addr)
}