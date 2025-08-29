package main


import (
	"log"
	"time"
	//"net/http"	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	"github.com/DylanCoon99/crypto_dashboard/insight-service/controllers"
)




func main() {


	/*
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	*/

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
		api.GET("/test", controllers.Test)

		// add other endpoints

	}


	log.Println("Server starting on port 8080...")

	log.Fatal(r.Run(":8080"))



}