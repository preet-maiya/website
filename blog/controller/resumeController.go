package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/preet-maiya/website/blog/config"
)

func GetResume(c *gin.Context) {
	config := config.GetConfig()
	// f, err := os.Open(config.ResumeFile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }
	// defer f.Close()

	// //Set header
	// c.Header("Content-type", "application/pdf")

	//Stream to response
	fmt.Println(config.ResumeFile)
	c.File(config.ResumeFile)
	c.Status(http.StatusOK)
}
