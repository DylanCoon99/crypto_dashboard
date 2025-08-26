package controllers


import (
	"log"
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


func Insight(c *gin.Context) {

	log.Println("Insight endpoint")

	c.JSON(http.StatusOK, gin.H{
      "message": "<Insert Insight Data Here>",
    })

    return
} 


func Sentiment(c *gin.Context) {

	log.Println("Sentiment endpoint")

	c.JSON(http.StatusOK, gin.H{
      "message": "<Insert Sentiment Data Here>",
    })

    return
} 


