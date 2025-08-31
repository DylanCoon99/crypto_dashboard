package controllers


import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/storage"
	//"github.com/DylanCoon99/crypto_dashboard/crypto-service/api"
	"github.com/DylanCoon99/crypto_dashboard/insight-service/cloud"
)

// stores configuration details; e.g. cloud client stuff
type ApiCfg struct {
	Client *storage.Client
}


func (cfg *ApiCfg) Test(c *gin.Context) {

	log.Println("Test endpoint")

	c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
    return

} 



func (cfg *ApiCfg) TestBucket(c *gin.Context) {

	// pulls data from the cloud bucket for that coin name

	coin_name := c.Param("coin_name")

	data := cloud.ReadFromBucket(cfg.Client, coin_name, c.Request.Context())


	c.JSON(http.StatusOK, gin.H{
      "message": data,
    })

}