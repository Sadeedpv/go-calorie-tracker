package routes

import (
	"github.com/Sadeedpv/go-calorie-tracker/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine){
	routeGroup := r.Group("/v1")
	{
		routeGroup.GET("/calories", controllers.GetAllCalories) // done
		routeGroup.POST("/calories", controllers.AddCalories)  // done
		routeGroup.DELETE("/calories", controllers.DeleteCalories) // done
		routeGroup.GET("/calories/:id", controllers.GetCaloriesById) // done
		routeGroup.PUT("/calories/:id", controllers.UpdateCalories) // done
		routeGroup.GET("/totalcalories", controllers.GetTotalCalories) // done
		routeGroup.DELETE("/calories/:id", controllers.DeleteCaloriesById) // done
	}
}