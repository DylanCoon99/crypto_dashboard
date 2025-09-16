package main

import (
	"log"
	"time"

	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	//"github.com/joho/godotenv"
	"github.com/DylanCoon99/crypto_dashboard/crypto-service/controllers"
	//"github.com/DylanCoon99/crypto_dashboard/crypto-service/crypto-api"
)

func main() {

	/*
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}

	*/

	// Api configuration setup
	var apiCfg controllers.ApiConfig

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	apiCfg.Upgrader = &upgrader

	// gin server setup
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		//test endpoint
		api.GET("/test", controllers.Test)

		// endpoint for retrieving AI Insights
		api.GET("/insight/:coin_name", controllers.Insight)

		// endpoint for retrieving sentiment
		api.GET("/sentiment/:coin_name", controllers.Sentiment)

		// endpoint for retrieving historical price data for past 24hrs
		api.GET("/price/historic/:coin_name", controllers.HistoricalPrice)

		// endpoint for retrieving real time price data for past 24hrs
		api.GET("/price/realtime/:coin_name", controllers.RealTimePrice)

		// endpoint for streaming "real-time" price data
		api.GET("/ws", apiCfg.HandleWebSocket)

	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start a goroutine to get real time prices via api
	go controllers.GetRealTimePrices(ctx)
	go controllers.Broadcast(ctx)

	log.Println("Server starting on port 8080...")

	log.Fatal(r.Run(":8080"))

}
