package controllers

import (
	"sync"

	"github.com/Sadeedpv/go-calorie-tracker/models"
	"github.com/Sadeedpv/go-calorie-tracker/utils"
	"github.com/gin-gonic/gin"
)

/*
	v1/calories
	GET
	Select all from DB
*/

func GetAllCalories(r *gin.Context) {
	// Fetching data from a table
	DB := utils.Db
	rows, err := DB.Query("SELECT * FROM public.calories;")
	if err != nil {
		utils.RespondWithError(r, err, "Internal Server Error")
		return
	}
	defer rows.Close()
	var calories []models.Calories

	// Create a channel to recieve calorie data
	calorieChan := make(chan models.Calories)

	var wg sync.WaitGroup
	for rows.Next(){
		var calorie models.Calories
		err := rows.Scan(&calorie.ID, &calorie.Food, &calorie.Calorie)
		if err != nil{
			utils.RespondWithError(r, err, "Internal Server Error")
		}
		wg.Add(1)
		go func(c models.Calories){
			defer wg.Done()
			// Send data to the channel
			calorieChan <- c
		}(calorie)
	}

	// Close the channel when all goroutines are done
	go func(){
		wg.Wait()
		close(calorieChan)
	}()

	for c := range calorieChan{
		calories = append(calories, c)
	}
	if len(calories) == 0{
		utils.RespondWithJSON(r, []models.Calories{})
	}else{
		utils.RespondWithJSON(r, calories)
	}
}


/*
	v1/calories
	POST
	Insert into DB
*/


func AddCalories(r *gin.Context) {
	var calorie models.Calories
	if err := r.ShouldBindJSON(&calorie); err != nil{
		utils.RespondWithError(r, err, "Something wrong with the sent data format")
		return
	}
	DB := utils.Db
	if calorie.ID == nil{
		row,err := DB.Exec("INSERT INTO public.calories(food, calorie) VALUES ($1, $2)", calorie.Food, calorie.Calorie)
		if err != nil{
			utils.RespondWithError(r, err, "Failed to add the data, check the input values")
			return
		}
		// Check If any rows affected 
		if result,_ := row.RowsAffected(); result < 1{
			utils.RespondWithError(r, err,"Failed to Execute DB Query")
			return
		}
	}else{
		row,err := DB.Exec("INSERT INTO public.calories(id, food, calorie) VALUES ($1, $2, $3)", calorie.ID, calorie.Food, calorie.Calorie)
		if err != nil{
			utils.RespondWithError(r, err, "Failed to add the data, check the input values")
			return
		}
		// Check If any rows affected 
		if result,_ := row.RowsAffected(); result < 1{
			utils.RespondWithError(r,err, "Failed to Execute DB Query")
			return
		}
	}
	utils.RespondWithJSON(r, gin.H{"message":"Data Added Successfully"})
}


/*
	v1/calories
	DELETE
	Erase all rows
*/


func DeleteCalories(r *gin.Context) {
	DB := utils.Db
	row, err := DB.Exec("DELETE FROM public.calories")
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
		return
	}
	// Check if any rows affected
	result,_ := row.RowsAffected(); if result < 1{
		utils.RespondWithError(r,err, "Failed to Execute DB Query")
		return
	}
	utils.RespondWithJSON(r, gin.H{"message":"Database deleted successfully!"})
}



/*
	v1/calories/:id
	GET
	Query one row with ID
*/

func GetCaloriesById(r *gin.Context) {
	DB := utils.Db
	id := r.Param("id")
	row:= DB.QueryRow("SELECT * FROM public.calories WHERE ID=$1;", id)
	var calorie models.Calories
	err := row.Scan(&calorie.ID, &calorie.Food, &calorie.Calorie)
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
		return
	}
	utils.RespondWithJSON(r, calorie)
}

/*
	v1/calories/:id
	PUT
	Update row with ID
*/


func UpdateCalories(r *gin.Context) {
	DB := utils.Db
	id := r.Param("id")
	var calorie models.Calories
	if err := r.ShouldBindJSON(&calorie); err != nil{
		utils.RespondWithError(r, err, "Something wrong with the sent data format")
		return
	}
	row, err := DB.Exec("UPDATE public.calories SET food=$1, calorie=$2 WHERE id=$3", calorie.Food, calorie.Calorie, id)
	if err != nil{
		utils.RespondWithError(r, err, "Failed to add the data, check the input values")
		return
	}
	// Check if any rows affected
	result,_ := row.RowsAffected(); if result < 1{
		utils.RespondWithError(r,err, "Failed to Execute DB Query")
		return
	}

	utils.RespondWithJSON(r, gin.H{"message": "Data Updated successfully!"})
}

/*
	v1/totalcalories
	GET
	Query Sum of calories
*/
func GetTotalCalories(r *gin.Context) {
	DB := utils.Db
	row := DB.QueryRow("SELECT COALESCE(SUM(calorie), 0) FROM public.calories;")
	var sum int
	err := row.Scan(&sum)
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
		return
	}
	utils.RespondWithJSON(r, gin.H{"TotalCalories": sum})
}


/*
	v1/calories/:id
	DELETE
	Delete one row
*/

func DeleteCaloriesById(r *gin.Context){
	DB := utils.Db
	id := r.Param("id")
	row,err := DB.Exec("DELETE FROM public.calories WHERE ID=$1", id)
	if err != nil{
		utils.RespondWithError(r, err, "Internal Server Error")
		return
	}
	// Check if any rows affected
	result,_ := row.RowsAffected(); if result < 1{
		utils.RespondWithError(r,err, "Failed to Execute DB Query")
		return
	}
	utils.RespondWithJSON(r, gin.H{"message":"Row Deleted Successfully!"})
}