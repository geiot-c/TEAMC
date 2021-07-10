package main

import (
	"log"
	"net/http"

	"TEAMC/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("./static/html/*.html")
	engine.Static("/js", "./static/js")
	engine.Static("/css", "./static/css")
	engine.Static("/image", "./static/image")

	engine.GET("/select", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	engine.POST("/result", controller.GetCandidates)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Run: ", err)
	}
}
