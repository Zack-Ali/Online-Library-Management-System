package routes

import (
	"online_library/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(app *gin.Engine) {
	app.GET("/api/user/", controllers.GetAllUsers)
	app.POST("/api/user/", controllers.CreateUser)
	app.POST("/api/user/login/", controllers.UserLogin)
	// app.GET("/api/user/:id", controllers.GetCalorieTracker)
	// app.PATCH("/api/user/:id", controllers.UpdateCalorieTracker)
	// app.DELETE("/api/user/:id", controllers.DeleteCalorieTracker)
}
