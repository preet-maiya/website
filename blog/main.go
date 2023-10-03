package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/preet-maiya/website/blog/controller"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("/app/dist", true)))
	router.GET("/hello", helloWorld)
	router.GET("/resume", controller.GetResume)

	router.Run("0.0.0.0:8080")
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
