package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preet-maiya/website/blog/controller"
)

func main() {
	router := gin.Default()
	router.GET("/", root)
	router.GET("/hello", helloWorld)
	router.GET("/resume", controller.GetResume)

	router.Run("0.0.0.0:8080")
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func root(c *gin.Context) {
	c.Redirect(http.StatusFound, "/resume")
}
