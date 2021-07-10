package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("./*.html")
	engine.Static("/js", "./js")
	engine.Static("/css", "./css")
	engine.Static("/image", "./image")

	engine.GET("/select", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal("Run: ", err)
	}
}
