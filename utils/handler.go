package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func RespondWithJSON(r *gin.Context, statusCode int, p interface{}){
	// Convert the 'calories' slice to JSON
	caloriesJSON, err := json.Marshal(p)
	if err != nil {
		log.Print("Error marshaling JSON:", err)
		r.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Send the JSON response
	r.Data(statusCode, "application/json", caloriesJSON)
}

func RespondWithError(r *gin.Context, err error, statusCode int, msg string){
	log.Print("Error Executing query: ", err)
	r.JSON(statusCode, gin.H{"error": msg})
}

