package routes

import (
	"github.com/Sadeedpv/go-calorie-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine){
	routeGroup := r.Group("/v1")
	{
		routeGroup.GET("/calories", controllers.GetAllCalories)
		routeGroup.POST("/calories", controllers.AddCalories)
		routeGroup.DELETE("/calories", controllers.DeleteCalories)
		routeGroup.GET("/calories/:id", controllers.GetCaloriesById)
		routeGroup.DELETE("/calories/:id", controllers.DeleteCaloriesById)
		routeGroup.PUT("/calories/:id", controllers.UpdateCalories)
		routeGroup.PATCH("/calories/:id", controllers.PatchCalories)
		routeGroup.GET("/totalcalories", controllers.GetTotalCalories)
	}
}