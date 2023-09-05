package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


func RespondWithJSON(r *gin.Context, p interface{}){
	// Convert the 'calories' slice to JSON
	caloriesJSON, err := json.Marshal(p)
	if err != nil {
		fmt.Print("Error marshaling JSON:", err)
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Send the JSON response
	r.Data(http.StatusOK, "application/json", caloriesJSON)
}

func RespondWithError(r *gin.Context, err error, msg string){
	fmt.Print("Error Executing query: ", err)
	r.JSON(http.StatusInternalServerError, gin.H{"error": msg})
}

