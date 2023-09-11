package main

import (
	"log"
	"os"

	"github.com/Sadeedpv/go-calorie-tracker/routes"
	"github.com/Sadeedpv/go-calorie-tracker/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)


func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
	if err := utils.InitializeDatabase(); err != nil{
		log.Fatal("Failed to Connect to DB")
	}
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.Default()

	// Configuration
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.ForwardedByClientIP = true
	// router.SetTrustedProxies([]string{os.Getenv("PROXY")})
	routes.SetUpRoutes(router)

	router.Run("localhost:" + port)

}