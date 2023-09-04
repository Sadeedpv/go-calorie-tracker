package main

import (
	"log"
	"os"

	"github.com/Sadeedpv/go-calorie-tracker/routes"
	"github.com/Sadeedpv/go-calorie-tracker/utils"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)


func main(){
	godotenv.Load()
	if err := utils.InitializeDatabase(); err != nil{
		log.Fatal("Failed to Connect to DB")
	}
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	router := gin.Default()
	routes.SetUpRoutes(router)

	router.Run("localhost:" + port)

}