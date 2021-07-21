package main

import (
	"fmt"
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

	engine.GET("/select/candidates", controller.GetCandidates)

	engine.GET("/result", func(c *gin.Context) {
		fmt.Println(c.QueryArray("ids[]"))
		c.HTML(http.StatusOK, "result.html", gin.H{})
	})
	engine.POST("/result", controller.GetResult)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Run: ", err)
	}
}
