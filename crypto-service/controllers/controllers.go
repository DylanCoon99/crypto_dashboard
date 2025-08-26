package controllers


import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func Test(c *gin.Context) {

	log.Println("Test endpoint")

	c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
    return

} 

