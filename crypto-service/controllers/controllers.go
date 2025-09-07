package controllers

import (
	"log"
	"net/http"

	"github.com/DylanCoon99/crypto_dashboard/crypto-service/api"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ApiConfig struct {
	Upgrader *websocket.Upgrader
}

func Test(c *gin.Context) {

	log.Println("Test endpoint")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	return

}

func Insight(c *gin.Context) {

	//log.Println("Insight endpoint is doing it's thing")

	coin_name := c.Param("coin_name")

	data := api.InsightServiceAPI(coin_name)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

	return
}

func Sentiment(c *gin.Context) {

	coin_name := c.Param("coin_name")

	data := api.SentimentServiceAPI(coin_name)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

	return
}

func HistoricalPrice(c *gin.Context) {

	coin_name := c.Param("coin_name")

	data := api.HistoricPriceAPI(coin_name)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

	return

}

func RealTimePrice(c *gin.Context) {

	coin_name := c.Param("coin_name")

	data := api.RealTimePriceAPI(coin_name)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

	return
}

