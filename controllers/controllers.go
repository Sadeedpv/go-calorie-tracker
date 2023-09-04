package controllers

import (
	"fmt"
	"github.com/Sadeedpv/go-calorie-tracker/models"
	"github.com/Sadeedpv/go-calorie-tracker/utils"
	"github.com/gin-gonic/gin"
)




func GetAllCalories(r *gin.Context) {
	// Fetching data from a table
	DB := utils.Db
	rows, err := DB.Query("SELECT * FROM public.calories;")
	if err != nil {
		utils.RespondWithError(r, err, "Internal Server Error")
	}
	defer rows.Close()
	var calories []models.Calories
	for rows.Next(){
		var calorie models.Calories
		err := rows.Scan(&calorie.ID, &calorie.Food, &calorie.Calorie)
		if err != nil{
			utils.RespondWithError(r, err, "Internal Server Error")
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
	DB := utils.Db
	id := r.Param("id")
	row:= DB.QueryRow("SELECT * FROM public.calories WHERE ID=$1;", id)
	var calorie models.Calories
	err := row.Scan(&calorie.ID, &calorie.Food, &calorie.Calorie)
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
	}
	utils.RespondWithJSON(r, calorie)
}
func UpdateCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}
func PatchCalories(r *gin.Context) {
	fmt.Print("Hello world!")
}
func GetTotalCalories(r *gin.Context) {
	DB := utils.Db
	row := DB.QueryRow("SELECT SUM(calorie) FROM public.calories;")
	var sum int
	err := row.Scan(&sum)
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
	}
	utils.RespondWithJSON(r, gin.H{"TotalCalories": sum})
}

func DeleteCaloriesById(r *gin.Context){
	DB := utils.Db
	id := r.Param("id")
	_,err := DB.Exec("DELETE FROM public.calories WHERE ID=$1", id)
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
	}
	utils.RespondWithJSON(r, gin.H{"message":"Row Deleted Successfully!"})
}