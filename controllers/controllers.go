package controllers

import (
	"fmt"
	"log"
	"github.com/Sadeedpv/go-calorie-tracker/models"
	"github.com/Sadeedpv/go-calorie-tracker/utils"
	"github.com/gin-gonic/gin"
)




func GetAllCalories(r *gin.Context) {
	// Fetching data from a table
	DB := utils.Db
	rows, err := DB.Query("SELECT * FROM public.calories;")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()
	var calories []models.Calories
	for rows.Next(){
		var calorie models.Calories
		err := rows.Scan(&calorie.ID, &calorie.Food, &calorie.Calorie)
		if err != nil{
			utils.RespondWithError(r, err, "Internal Server Error")
			return
		}
		calories = append(calories, calorie)
	}
	utils.RespondWithJSON(r, calories)
}
func AddCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}
func DeleteCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}
func GetCaloriesById(r *gin.Context) {
	fmt.Print("Hello world!")
}
func UpdateCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}
func PatchCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}
func GetTotalCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}