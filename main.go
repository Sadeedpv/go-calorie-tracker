package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sadeedpv/go-calorie-tracker/routes"
	"github.com/Sadeedpv/go-calorie-tracker/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main(){
	err := godotenv.Load()
	if err != nil{
		fmt.Print("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)

	// Set retry intervals
	maxRetries := 100
	retryInterval := 10 * time.Second
	for i:=0;i<maxRetries;i++{
		if err := utils.InitializeDatabase(); err != nil{
			log.Print("Failed to Connect to DB: ", err)
			time.Sleep(retryInterval)
		}
		break
	}
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.Default()

	fmt.Print("Router is running on: ", port)

	// Configuration
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.ForwardedByClientIP = true
	// router.SetTrustedProxies([]string{os.Getenv("PROXY")})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	routes.SetUpRoutes(router)

	runErr := router.Run("0.0.0.0:" + port)
	if runErr != nil{
		log.Fatal("Error running server: ", runErr)
	}


}