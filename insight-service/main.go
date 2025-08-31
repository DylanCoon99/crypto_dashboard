package main

import (
	"context"
	"log"
	"time"

	//"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//"github.com/joho/godotenv"
	"cloud.google.com/go/storage"
	"github.com/DylanCoon99/crypto_dashboard/insight-service/controllers"
)

func main() {

	// create the cloud client here
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatal("Failed to create cloud client")
	}

	// create the api config here
	apiCfg := controllers.ApiCfg{
		Client: client,
	}

	// gin server setup
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		//test endpoint
		api.GET("/test", apiCfg.Test)

		// add other endpoints
		api.GET("/testbucket/:coin_name", apiCfg.TestBucket)

	}

	log.Println("Server starting on port 8080...")

	log.Fatal(r.Run(":8080"))

}
