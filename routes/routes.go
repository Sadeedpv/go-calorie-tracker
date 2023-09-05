package routes

import (
	"github.com/Sadeedpv/go-calorie-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine){
	routeGroup := r.Group("/v1")
	{
		routeGroup.GET("/calories", controllers.GetAllCalories) // done
		routeGroup.POST("/calories", controllers.AddCalories) 
		routeGroup.DELETE("/calories", controllers.DeleteCalories) // done
		routeGroup.GET("/calories/:id", controllers.GetCaloriesById) // done
		routeGroup.DELETE("/calories/:id", controllers.DeleteCaloriesById) // done
		routeGroup.PUT("/calories/:id", controllers.UpdateCalories)
		routeGroup.PATCH("/calories/:id", controllers.PatchCalories)
		routeGroup.GET("/totalcalories", controllers.GetTotalCalories)
	}
}