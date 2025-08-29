package controllers


import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/DylanCoon99/crypto_dashboard/crypto-service/api"
)


func Test(c *gin.Context) {

	log.Println("Test endpoint")

	c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
    return

} 



