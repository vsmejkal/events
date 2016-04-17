package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/vsmejkal/events/config"
	"fmt"
	"path"
	"os"
	"log"
	ct "github.com/vsmejkal/events/admin/controllers"
)

func printHelp() {
	fmt.Printf("Usage: %s configFile\n", path.Base(os.Args[0]))
}

func createRenders() multitemplate.Render {
	r := multitemplate.New()
	reg := func(name, layout, content string) {
		r.AddFromFiles(
			name,
			config.Admin.DocumentRoot + "templates/layouts/" + layout + ".tmpl",
			config.Admin.DocumentRoot + "templates/" + content + ".tmpl",
		)
	}

	reg("sources.list", "base", "sources/list")

	return r
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
	router.HTMLRender = createRenders()
	router.Static("/assets", config.Admin.DocumentRoot + "/assets")

	router.GET("/sources", ct.SourceList)
	router.POST("/sources", ct.SourceCreate)
	router.GET("/sources/:id", ct.SourceRead)
	router.POST("/sources/:id", ct.SourceUpdate)
	router.DELETE("/sources/:id", ct.SourceDelete)

	addr := fmt.Sprintf(":%d", config.Admin.Port)
	router.Run(addr)
}