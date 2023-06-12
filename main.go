package main

import (
	"github.com/ericprd/articles-be/handlers"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", handlers.ShowIndex)

	r.Run(":3000")
}
